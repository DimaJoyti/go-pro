# GO-PRO GCP Deployment Guide

This guide provides step-by-step instructions for deploying the GO-PRO learning platform to Google Cloud Platform.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [GCP Project Setup](#gcp-project-setup)
3. [Infrastructure Deployment](#infrastructure-deployment)
4. [Application Deployment](#application-deployment)
5. [Post-Deployment Configuration](#post-deployment-configuration)
6. [Monitoring & Logging](#monitoring--logging)
7. [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Tools

Install the following tools on your local machine:

```bash
# gcloud CLI
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
gcloud init

# kubectl
gcloud components install kubectl

# Helm
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Terraform
wget https://releases.hashicorp.com/terraform/1.6.0/terraform_1.6.0_linux_amd64.zip
unzip terraform_1.6.0_linux_amd64.zip
sudo mv terraform /usr/local/bin/

# Verify installations
gcloud version
kubectl version --client
helm version
terraform version
```

### GCP Project Requirements

- GCP project with billing enabled
- Sufficient quotas for:
  - GKE clusters
  - Compute Engine instances
  - Cloud SQL instances
  - Memorystore instances
  - VPCs and subnets

## GCP Project Setup

### 1. Create GCP Project

```bash
# Create new project
gcloud projects create gopro-project --name="GO-PRO Learning Platform"

# Set as default project
gcloud config set project gopro-project

# Link billing account
gcloud beta billing projects link gopro-project \
  --billing-account=<BILLING_ACCOUNT_ID>

# Verify project
gcloud config list
```

### 2. Enable Required APIs

```bash
# Enable all required APIs
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
  artifactregistry.googleapis.com \
  dns.googleapis.com \
  cloudarmor.googleapis.com

# Verify enabled services
gcloud services list --enabled
```

### 3. Create Service Account

```bash
# Create service account
gcloud iam service-accounts create gopro-terraform \
  --display-name="Terraform Service Account"

# Grant necessary roles
gcloud projects add-iam-policy-binding gopro-project \
  --member="serviceAccount:gopro-terraform@gopro-project.iam.gserviceaccount.com" \
  --role="roles/editor"

# Create and download key
gcloud iam service-accounts keys create ~/gopro-terraform-key.json \
  --iam-account=gopro-terraform@gopro-project.iam.gserviceaccount.com

# Set environment variable
export GOOGLE_APPLICATION_CREDENTIALS=~/gopro-terraform-key.json
```

### 4. Create Cloud Storage Bucket for Terraform State

```bash
# Create bucket
gsutil mb -p gopro-project -l us-central1 gs://gopro-terraform-state

# Enable versioning
gsutil versioning set on gs://gopro-terraform-state

# Set lifecycle policy
cat > lifecycle.json <<EOF
{
  "lifecycle": {
    "rule": [
      {
        "action": {"type": "Delete"},
        "condition": {"numNewerVersions": 3}
      }
    ]
  }
}
EOF

gsutil lifecycle set lifecycle.json gs://gopro-terraform-state
```

## Infrastructure Deployment

### Option 1: Using Terraform (Recommended)

```bash
# Navigate to terraform directory
cd terraform

# Initialize Terraform
terraform init

# Review the plan for development environment
terraform plan -var-file=environments/gcp-dev.tfvars

# Apply the configuration
terraform apply -var-file=environments/gcp-dev.tfvars

# Save outputs
terraform output > ../gcp/terraform-outputs.txt
```

### Option 2: Using gcloud

```bash
# Navigate to gcp directory
cd gcp

# Create VPC
gcloud compute networks create gopro-vpc \
  --subnet-mode=custom \
  --bgp-routing-mode=regional

# Create subnet
gcloud compute networks subnets create gopro-subnet \
  --network=gopro-vpc \
  --region=us-central1 \
  --range=10.0.0.0/24 \
  --secondary-range pods=10.1.0.0/16,services=10.2.0.0/16

# Create GKE cluster
gcloud container clusters create gopro-dev-cluster \
  --region us-central1 \
  --num-nodes 2 \
  --machine-type e2-medium \
  --enable-autoscaling \
  --min-nodes 1 \
  --max-nodes 5 \
  --network gopro-vpc \
  --subnetwork gopro-subnet \
  --enable-ip-alias \
  --cluster-secondary-range-name pods \
  --services-secondary-range-name services \
  --workload-pool gopro-project.svc.id.goog \
  --enable-stackdriver-kubernetes

# Create Cloud SQL instance
gcloud sql instances create gopro-dev-db \
  --database-version=POSTGRES_15 \
  --tier=db-f1-micro \
  --region=us-central1 \
  --network=gopro-vpc \
  --no-assign-ip

# Create Memorystore Redis
gcloud redis instances create gopro-dev-redis \
  --size=1 \
  --region=us-central1 \
  --network=gopro-vpc \
  --redis-version=redis_7_0
```

### 3. Verify Infrastructure

```bash
# Get cluster credentials
gcloud container clusters get-credentials gopro-dev-cluster \
  --region us-central1

# Verify cluster access
kubectl get nodes

# Check cluster info
kubectl cluster-info

# List Cloud SQL instances
gcloud sql instances list

# List Redis instances
gcloud redis instances list --region=us-central1
```

## Application Deployment

### 1. Setup Cluster Components

Run the cluster setup script to install required components:

```bash
cd gcp/scripts
./setup-cluster.sh
```

This script installs:
- NGINX Ingress Controller
- cert-manager
- Metrics Server
- Prometheus & Grafana
- Jaeger
- Workload Identity configuration

### 2. Create Secrets

Create required secrets for the application:

```bash
# Get Cloud SQL connection name
SQL_CONNECTION=$(gcloud sql instances describe gopro-dev-db \
  --format='value(connectionName)')

# Get Redis host
REDIS_HOST=$(gcloud redis instances describe gopro-dev-redis \
  --region=us-central1 \
  --format='value(host)')

# Database credentials
kubectl create secret generic database-credentials \
  --from-literal=host=127.0.0.1 \
  --from-literal=port=5432 \
  --from-literal=username=gopro \
  --from-literal=password=<DB_PASSWORD> \
  --from-literal=database=gopro \
  -n gopro

# Redis credentials
kubectl create secret generic redis-credentials \
  --from-literal=url=redis://$REDIS_HOST:6379 \
  -n gopro

# JWT secret
kubectl create secret generic jwt-secret \
  --from-literal=secret=$(openssl rand -base64 32) \
  -n gopro

# Cloud SQL instance credentials
kubectl create secret generic cloudsql-instance-credentials \
  --from-literal=instance-connection-name=$SQL_CONNECTION \
  -n gopro
```

### 3. Deploy Application

```bash
# Deploy to development
cd gcp/scripts
ENVIRONMENT=development ./deploy-app.sh

# Or deploy manually
kubectl apply -k ../../k8s/overlays/development

# Check deployment status
kubectl get pods -n gopro
kubectl get svc -n gopro
kubectl get ingress -n gopro
```

### 4. Verify Deployment

```bash
# Check pod status
kubectl get pods -n gopro

# Check logs
kubectl logs -f deployment/backend -n gopro
kubectl logs -f deployment/frontend -n gopro

# Check services
kubectl get svc -n gopro

# Get load balancer IP
kubectl get svc ingress-nginx-controller -n ingress-nginx
```

## Post-Deployment Configuration

### 1. Configure DNS

```bash
# Get load balancer IP
LB_IP=$(kubectl get svc ingress-nginx-controller -n ingress-nginx \
  -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

# Create Cloud DNS zone
gcloud dns managed-zones create gopro-zone \
  --dns-name=gopro.com \
  --description="GO-PRO DNS Zone"

# Add A record
gcloud dns record-sets transaction start --zone=gopro-zone

gcloud dns record-sets transaction add $LB_IP \
  --name=dev.gopro.com \
  --ttl=300 \
  --type=A \
  --zone=gopro-zone

gcloud dns record-sets transaction execute --zone=gopro-zone
```

### 2. Configure SSL/TLS

```bash
# Create ClusterIssuer for Let's Encrypt
cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: admin@gopro.com
    privateKeySecretRef:
      name: letsencrypt-prod
    solvers:
    - http01:
        ingress:
          class: nginx
EOF

# Create Certificate
cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: gopro-tls
  namespace: gopro
spec:
  secretName: gopro-tls
  issuerRef:
    name: letsencrypt-prod
    kind: ClusterIssuer
  dnsNames:
  - dev.gopro.com
  - dev-api.gopro.com
EOF
```

### 3. Configure Monitoring

```bash
# Access Grafana
kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring

# Default credentials:
# Username: admin
# Password: admin

# Create custom dashboard
gcloud monitoring dashboards create --config-from-file=dashboard.json

# Create alert policy
gcloud alpha monitoring policies create \
  --notification-channels=<CHANNEL_ID> \
  --display-name="High CPU Usage" \
  --condition-display-name="CPU > 80%" \
  --condition-threshold-value=0.8
```

## Monitoring & Logging

### Cloud Logging

```bash
# View cluster logs
gcloud logging read "resource.type=k8s_cluster" --limit 50

# View application logs
gcloud logging read "resource.type=k8s_pod AND resource.labels.namespace_name=gopro" --limit 50

# Create log sink
gcloud logging sinks create gopro-logs \
  storage.googleapis.com/gopro-logs-bucket \
  --log-filter='resource.type="k8s_cluster"'

# Tail logs
gcloud logging tail "resource.type=k8s_pod AND resource.labels.namespace_name=gopro"
```

### Cloud Monitoring

```bash
# List metrics
gcloud monitoring metrics-descriptors list \
  --filter="metric.type:kubernetes.io"

# Get metric data
gcloud monitoring time-series list \
  --filter='metric.type="kubernetes.io/container/cpu/core_usage_time"' \
  --interval-start-time="2024-01-01T00:00:00Z" \
  --interval-end-time="2024-01-01T01:00:00Z"
```

### Cloud Trace

```bash
# List traces
gcloud trace list-traces --limit=10

# Get trace details
gcloud trace get-trace <TRACE_ID>
```

## Troubleshooting

### Common Issues

#### 1. Pods Not Starting

```bash
# Check pod status
kubectl get pods -n gopro

# Describe pod
kubectl describe pod <pod-name> -n gopro

# Check logs
kubectl logs <pod-name> -n gopro

# Check events
kubectl get events -n gopro --sort-by='.lastTimestamp'
```

#### 2. Database Connection Issues

```bash
# Test Cloud SQL connectivity
kubectl run -it --rm debug --image=postgres:15 --restart=Never -- \
  psql -h 127.0.0.1 -U gopro -d gopro

# Check Cloud SQL Proxy logs
kubectl logs <pod-name> -c cloudsql-proxy -n gopro

# Verify Workload Identity
kubectl describe sa backend-sa -n gopro
```

#### 3. Load Balancer Not Created

```bash
# Check ingress controller logs
kubectl logs -n ingress-nginx deployment/ingress-nginx-controller

# Check ingress events
kubectl describe ingress gopro-ingress -n gopro

# Verify service
kubectl get svc ingress-nginx-controller -n ingress-nginx
```

### Useful Commands

```bash
# Get all resources in namespace
kubectl get all -n gopro

# Restart deployment
kubectl rollout restart deployment/backend -n gopro

# Scale deployment
kubectl scale deployment/backend --replicas=5 -n gopro

# Update image
kubectl set image deployment/backend \
  backend=us-central1-docker.pkg.dev/gopro-project/gopro/backend:v1.1.0 \
  -n gopro

# Rollback deployment
kubectl rollout undo deployment/backend -n gopro

# Get deployment history
kubectl rollout history deployment/backend -n gopro

# Execute command in pod
kubectl exec -it deployment/backend -n gopro -- /bin/sh

# Port forward to service
kubectl port-forward svc/backend 8080:80 -n gopro
```

## Cleanup

### Delete Application

```bash
# Delete application resources
kubectl delete -k k8s/overlays/development

# Delete namespace
kubectl delete namespace gopro
```

### Delete Infrastructure

```bash
# Using Terraform
cd terraform
terraform destroy -var-file=environments/gcp-dev.tfvars

# Using gcloud
gcloud container clusters delete gopro-dev-cluster --region us-central1
gcloud sql instances delete gopro-dev-db
gcloud redis instances delete gopro-dev-redis --region us-central1
gcloud compute networks subnets delete gopro-subnet --region us-central1
gcloud compute networks delete gopro-vpc
```

## Additional Resources

- [GKE Documentation](https://cloud.google.com/kubernetes-engine/docs)
- [Cloud SQL Documentation](https://cloud.google.com/sql/docs)
- [Memorystore Documentation](https://cloud.google.com/memorystore/docs)
- [Cloud Monitoring Documentation](https://cloud.google.com/monitoring/docs)
- [GCP Best Practices](https://cloud.google.com/architecture/framework)

