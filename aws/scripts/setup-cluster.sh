#!/bin/bash

# GO-PRO EKS Cluster Setup Script
# This script sets up an EKS cluster with all required components

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
CLUSTER_NAME="${CLUSTER_NAME:-gopro-dev-cluster}"
REGION="${AWS_REGION:-us-east-1}"
ENVIRONMENT="${ENVIRONMENT:-development}"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}GO-PRO EKS Cluster Setup${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Cluster Name: $CLUSTER_NAME"
echo "Region: $REGION"
echo "Environment: $ENVIRONMENT"
echo ""

# Check prerequisites
echo -e "${YELLOW}Checking prerequisites...${NC}"
command -v aws >/dev/null 2>&1 || { echo -e "${RED}AWS CLI is required but not installed.${NC}" >&2; exit 1; }
command -v eksctl >/dev/null 2>&1 || { echo -e "${RED}eksctl is required but not installed.${NC}" >&2; exit 1; }
command -v kubectl >/dev/null 2>&1 || { echo -e "${RED}kubectl is required but not installed.${NC}" >&2; exit 1; }
command -v helm >/dev/null 2>&1 || { echo -e "${RED}Helm is required but not installed.${NC}" >&2; exit 1; }
echo -e "${GREEN}✓ All prerequisites met${NC}"

# Verify AWS credentials
echo -e "${YELLOW}Verifying AWS credentials...${NC}"
aws sts get-caller-identity > /dev/null 2>&1 || { echo -e "${RED}AWS credentials not configured${NC}" >&2; exit 1; }
echo -e "${GREEN}✓ AWS credentials verified${NC}"

# Create EKS cluster
echo -e "${YELLOW}Creating EKS cluster...${NC}"
if eksctl get cluster --name "$CLUSTER_NAME" --region "$REGION" > /dev/null 2>&1; then
    echo -e "${YELLOW}Cluster $CLUSTER_NAME already exists${NC}"
else
    eksctl create cluster -f ../eks/cluster-config.yaml
    echo -e "${GREEN}✓ EKS cluster created${NC}"
fi

# Update kubeconfig
echo -e "${YELLOW}Updating kubeconfig...${NC}"
aws eks update-kubeconfig --name "$CLUSTER_NAME" --region "$REGION"
echo -e "${GREEN}✓ kubeconfig updated${NC}"

# Verify cluster access
echo -e "${YELLOW}Verifying cluster access...${NC}"
kubectl get nodes
echo -e "${GREEN}✓ Cluster access verified${NC}"

# Install AWS Load Balancer Controller
echo -e "${YELLOW}Installing AWS Load Balancer Controller...${NC}"
helm repo add eks https://aws.github.io/eks-charts
helm repo update

kubectl create namespace kube-system --dry-run=client -o yaml | kubectl apply -f -

helm upgrade --install aws-load-balancer-controller eks/aws-load-balancer-controller \
  -n kube-system \
  --set clusterName="$CLUSTER_NAME" \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller \
  --wait

echo -e "${GREEN}✓ AWS Load Balancer Controller installed${NC}"

# Install EBS CSI Driver
echo -e "${YELLOW}Installing EBS CSI Driver...${NC}"
kubectl apply -k "github.com/kubernetes-sigs/aws-ebs-csi-driver/deploy/kubernetes/overlays/stable/?ref=release-1.25"
echo -e "${GREEN}✓ EBS CSI Driver installed${NC}"

# Install Cluster Autoscaler
echo -e "${YELLOW}Installing Cluster Autoscaler...${NC}"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/master/cluster-autoscaler/cloudprovider/aws/examples/cluster-autoscaler-autodiscover.yaml

kubectl -n kube-system annotate deployment.apps/cluster-autoscaler \
  cluster-autoscaler.kubernetes.io/safe-to-evict="false" \
  --overwrite

kubectl -n kube-system set image deployment.apps/cluster-autoscaler \
  cluster-autoscaler=registry.k8s.io/autoscaling/cluster-autoscaler:v1.28.0

echo -e "${GREEN}✓ Cluster Autoscaler installed${NC}"

# Install Metrics Server
echo -e "${YELLOW}Installing Metrics Server...${NC}"
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
echo -e "${GREEN}✓ Metrics Server installed${NC}"

# Create namespaces
echo -e "${YELLOW}Creating namespaces...${NC}"
kubectl create namespace gopro --dry-run=client -o yaml | kubectl apply -f -
kubectl create namespace monitoring --dry-run=client -o yaml | kubectl apply -f -
kubectl create namespace kafka --dry-run=client -o yaml | kubectl apply -f -
echo -e "${GREEN}✓ Namespaces created${NC}"

# Install Prometheus & Grafana
echo -e "${YELLOW}Installing Prometheus & Grafana...${NC}"
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm upgrade --install prometheus prometheus-community/kube-prometheus-stack \
  -n monitoring \
  --create-namespace \
  --set prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues=false \
  --wait

echo -e "${GREEN}✓ Prometheus & Grafana installed${NC}"

# Install Jaeger
echo -e "${YELLOW}Installing Jaeger...${NC}"
kubectl create namespace observability --dry-run=client -o yaml | kubectl apply -f -
kubectl apply -f https://github.com/jaegertracing/jaeger-operator/releases/download/v1.51.0/jaeger-operator.yaml -n observability
echo -e "${GREEN}✓ Jaeger installed${NC}"

# Install Strimzi Kafka Operator
echo -e "${YELLOW}Installing Strimzi Kafka Operator...${NC}"
kubectl create namespace kafka --dry-run=client -o yaml | kubectl apply -f -
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
echo -e "${GREEN}✓ Strimzi Kafka Operator installed${NC}"

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
echo "  kubectl port-forward svc/grafana 3000:80 -n monitoring"
echo ""

