#!/bin/bash

# GO-PRO GKE Cluster Setup Script
# This script sets up a GKE cluster with all required components

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
PROJECT_ID="${GCP_PROJECT_ID:-gopro-project}"
CLUSTER_NAME="${CLUSTER_NAME:-gopro-dev-cluster}"
REGION="${GCP_REGION:-us-central1}"
ENVIRONMENT="${ENVIRONMENT:-development}"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}GO-PRO GKE Cluster Setup${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Project ID: $PROJECT_ID"
echo "Cluster Name: $CLUSTER_NAME"
echo "Region: $REGION"
echo "Environment: $ENVIRONMENT"
echo ""

# Check prerequisites
echo -e "${YELLOW}Checking prerequisites...${NC}"
command -v gcloud >/dev/null 2>&1 || { echo -e "${RED}gcloud CLI is required but not installed.${NC}" >&2; exit 1; }
command -v kubectl >/dev/null 2>&1 || { echo -e "${RED}kubectl is required but not installed.${NC}" >&2; exit 1; }
command -v helm >/dev/null 2>&1 || { echo -e "${RED}Helm is required but not installed.${NC}" >&2; exit 1; }
echo -e "${GREEN}✓ All prerequisites met${NC}"

# Set project
echo -e "${YELLOW}Setting GCP project...${NC}"
gcloud config set project "$PROJECT_ID"
echo -e "${GREEN}✓ Project set${NC}"

# Enable required APIs
echo -e "${YELLOW}Enabling required APIs...${NC}"
gcloud services enable \
  container.googleapis.com \
  compute.googleapis.com \
  sqladmin.googleapis.com \
  redis.googleapis.com \
  pubsub.googleapis.com \
  storage-api.googleapis.com \
  cloudkms.googleapis.com \
  secretmanager.googleapis.com \
  monitoring.googleapis.com \
  logging.googleapis.com \
  cloudtrace.googleapis.com \
  cloudbuild.googleapis.com \
  artifactregistry.googleapis.com
echo -e "${GREEN}✓ APIs enabled${NC}"

# Create GKE cluster
echo -e "${YELLOW}Creating GKE cluster...${NC}"
if gcloud container clusters describe "$CLUSTER_NAME" --region "$REGION" > /dev/null 2>&1; then
    echo -e "${YELLOW}Cluster $CLUSTER_NAME already exists${NC}"
else
    gcloud container clusters create "$CLUSTER_NAME" \
      --region "$REGION" \
      --release-channel regular \
      --num-nodes 2 \
      --machine-type e2-medium \
      --enable-autoscaling \
      --min-nodes 1 \
      --max-nodes 5 \
      --enable-autorepair \
      --enable-autoupgrade \
      --enable-ip-alias \
      --enable-stackdriver-kubernetes \
      --enable-cloud-logging \
      --enable-cloud-monitoring \
      --workload-pool="$PROJECT_ID.svc.id.goog" \
      --addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver \
      --labels environment="$ENVIRONMENT",project=gopro
    echo -e "${GREEN}✓ GKE cluster created${NC}"
fi

# Get cluster credentials
echo -e "${YELLOW}Getting cluster credentials...${NC}"
gcloud container clusters get-credentials "$CLUSTER_NAME" --region "$REGION"
echo -e "${GREEN}✓ Credentials configured${NC}"

# Verify cluster access
echo -e "${YELLOW}Verifying cluster access...${NC}"
kubectl get nodes
echo -e "${GREEN}✓ Cluster access verified${NC}"

# Create namespaces
echo -e "${YELLOW}Creating namespaces...${NC}"
kubectl create namespace gopro --dry-run=client -o yaml | kubectl apply -f -
kubectl create namespace monitoring --dry-run=client -o yaml | kubectl apply -f -
kubectl create namespace pubsub --dry-run=client -o yaml | kubectl apply -f -
echo -e "${GREEN}✓ Namespaces created${NC}"

# Install NGINX Ingress Controller
echo -e "${YELLOW}Installing NGINX Ingress Controller...${NC}"
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-nginx \
  --create-namespace \
  --set controller.service.type=LoadBalancer \
  --set controller.metrics.enabled=true \
  --wait

echo -e "${GREEN}✓ NGINX Ingress Controller installed${NC}"

# Install cert-manager
echo -e "${YELLOW}Installing cert-manager...${NC}"
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.0/cert-manager.yaml
echo -e "${GREEN}✓ cert-manager installed${NC}"

# Install Metrics Server
echo -e "${YELLOW}Installing Metrics Server...${NC}"
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
echo -e "${GREEN}✓ Metrics Server installed${NC}"

# Install Prometheus & Grafana
echo -e "${YELLOW}Installing Prometheus & Grafana...${NC}"
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm upgrade --install prometheus prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --set prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues=false \
  --set grafana.adminPassword=admin \
  --wait

echo -e "${GREEN}✓ Prometheus & Grafana installed${NC}"

# Install Jaeger
echo -e "${YELLOW}Installing Jaeger...${NC}"
kubectl create namespace observability --dry-run=client -o yaml | kubectl apply -f -
kubectl apply -f https://github.com/jaegertracing/jaeger-operator/releases/download/v1.51.0/jaeger-operator.yaml -n observability
echo -e "${GREEN}✓ Jaeger installed${NC}"

# Setup Workload Identity
echo -e "${YELLOW}Setting up Workload Identity...${NC}"

# Create GCP service account
gcloud iam service-accounts create gopro-backend-sa \
  --display-name="GO-PRO Backend Service Account" \
  --project="$PROJECT_ID" || true

# Bind Kubernetes service account to GCP service account
kubectl create serviceaccount backend-sa -n gopro --dry-run=client -o yaml | kubectl apply -f -

gcloud iam service-accounts add-iam-policy-binding \
  "gopro-backend-sa@$PROJECT_ID.iam.gserviceaccount.com" \
  --role roles/iam.workloadIdentityUser \
  --member "serviceAccount:$PROJECT_ID.svc.id.goog[gopro/backend-sa]"

kubectl annotate serviceaccount backend-sa \
  -n gopro \
  iam.gke.io/gcp-service-account="gopro-backend-sa@$PROJECT_ID.iam.gserviceaccount.com" \
  --overwrite

echo -e "${GREEN}✓ Workload Identity configured${NC}"

# Create Cloud SQL Proxy
echo -e "${YELLOW}Setting up Cloud SQL Proxy...${NC}"
kubectl create secret generic cloudsql-instance-credentials \
  --from-literal=instance-connection-name="$PROJECT_ID:$REGION:gopro-dev-db" \
  -n gopro --dry-run=client -o yaml | kubectl apply -f -
echo -e "${GREEN}✓ Cloud SQL Proxy configured${NC}"

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Cluster setup complete!${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Next steps:"
echo "1. Configure secrets: kubectl create secret generic database-credentials ..."
echo "2. Deploy application: kubectl apply -k ../../k8s/overlays/$ENVIRONMENT"
echo "3. Check deployment: kubectl get pods -n gopro"
echo ""
echo "Useful commands:"
echo "  kubectl get nodes"
echo "  kubectl get pods -n gopro"
echo "  kubectl logs -f deployment/backend -n gopro"
echo "  kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring"
echo ""
echo "Get Ingress IP:"
echo "  kubectl get svc ingress-nginx-controller -n ingress-nginx"
echo ""

