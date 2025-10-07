# GO-PRO AWS Deployment Guide

This guide provides step-by-step instructions for deploying the GO-PRO learning platform to AWS.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [AWS Account Setup](#aws-account-setup)
3. [Infrastructure Deployment](#infrastructure-deployment)
4. [Application Deployment](#application-deployment)
5. [Post-Deployment Configuration](#post-deployment-configuration)
6. [Monitoring & Logging](#monitoring--logging)
7. [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Tools

Install the following tools on your local machine:

```bash
# AWS CLI
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

# kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# eksctl
curl --silent --location "https://github.com/weksctl-io/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin

# Helm
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Terraform
wget https://releases.hashicorp.com/terraform/1.6.0/terraform_1.6.0_linux_amd64.zip
unzip terraform_1.6.0_linux_amd64.zip
sudo mv terraform /usr/local/bin/

# Verify installations
aws --version
kubectl version --client
eksctl version
helm version
terraform version
```

### AWS Account Requirements

- AWS account with appropriate permissions
- IAM user with programmatic access
- Sufficient service quotas for:
  - EKS clusters
  - EC2 instances
  - RDS databases
  - ElastiCache clusters
  - VPCs and subnets

## AWS Account Setup

### 1. Configure AWS Credentials

```bash
# Configure AWS CLI
aws configure

# Enter your credentials:
# AWS Access Key ID: YOUR_ACCESS_KEY
# AWS Secret Access Key: YOUR_SECRET_KEY
# Default region name: us-east-1
# Default output format: json

# Verify configuration
aws sts get-caller-identity
```

### 2. Create S3 Bucket for Terraform State

```bash
# Create S3 bucket
aws s3 mb s3://gopro-terraform-state --region us-east-1

# Enable versioning
aws s3api put-bucket-versioning \
  --bucket gopro-terraform-state \
  --versioning-configuration Status=Enabled

# Enable encryption
aws s3api put-bucket-encryption \
  --bucket gopro-terraform-state \
  --server-side-encryption-configuration '{
    "Rules": [{
      "ApplyServerSideEncryptionByDefault": {
        "SSEAlgorithm": "AES256"
      }
    }]
  }'
```

### 3. Create DynamoDB Table for State Locking

```bash
aws dynamodb create-table \
  --table-name gopro-terraform-locks \
  --attribute-definitions AttributeName=LockID,AttributeType=S \
  --key-schema AttributeName=LockID,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --region us-east-1
```

## Infrastructure Deployment

### Option 1: Using Terraform (Recommended)

```bash
# Navigate to terraform directory
cd terraform

# Initialize Terraform
terraform init

# Review the plan for development environment
terraform plan -var-file=environments/dev.tfvars

# Apply the configuration
terraform apply -var-file=environments/dev.tfvars

# Save outputs
terraform output > ../aws/terraform-outputs.txt
```

### Option 2: Using eksctl

```bash
# Navigate to aws directory
cd aws

# Create EKS cluster
eksctl create cluster -f eks/cluster-config.yaml

# This will create:
# - VPC with public and private subnets
# - EKS cluster
# - Managed node group
# - IAM roles and policies
# - EKS addons (VPC CNI, CoreDNS, kube-proxy)
```

### 3. Verify Infrastructure

```bash
# Update kubeconfig
aws eks update-kubeconfig --name gopro-dev-cluster --region us-east-1

# Verify cluster access
kubectl get nodes

# Check cluster info
kubectl cluster-info

# Verify addons
kubectl get pods -n kube-system
```

## Application Deployment

### 1. Setup Cluster Components

Run the cluster setup script to install required components:

```bash
cd aws/scripts
./setup-cluster.sh
```

This script installs:
- AWS Load Balancer Controller
- EBS CSI Driver
- Cluster Autoscaler
- Metrics Server
- Prometheus & Grafana
- Jaeger
- Strimzi Kafka Operator

### 2. Create Secrets

Create required secrets for the application:

```bash
# Database credentials
kubectl create secret generic database-credentials \
  --from-literal=host=<RDS_ENDPOINT> \
  --from-literal=port=5432 \
  --from-literal=username=gopro \
  --from-literal=password=<DB_PASSWORD> \
  --from-literal=database=gopro \
  -n gopro

# Redis credentials
kubectl create secret generic redis-credentials \
  --from-literal=url=redis://<ELASTICACHE_ENDPOINT>:6379 \
  -n gopro

# JWT secret
kubectl create secret generic jwt-secret \
  --from-literal=secret=<RANDOM_SECRET> \
  -n gopro

# GitHub Container Registry credentials (if using private registry)
kubectl create secret docker-registry ghcr-secret \
  --docker-server=ghcr.io \
  --docker-username=<GITHUB_USERNAME> \
  --docker-password=<GITHUB_TOKEN> \
  -n gopro
```

### 3. Deploy Application

```bash
# Deploy to development
cd aws/scripts
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

# Get load balancer URL
kubectl get ingress gopro-ingress -n gopro
```

## Post-Deployment Configuration

### 1. Configure DNS

```bash
# Get load balancer hostname
LB_HOSTNAME=$(kubectl get ingress gopro-ingress -n gopro -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')

# Create Route 53 records
aws route53 change-resource-record-sets \
  --hosted-zone-id <ZONE_ID> \
  --change-batch '{
    "Changes": [{
      "Action": "CREATE",
      "ResourceRecordSet": {
        "Name": "dev.gopro.com",
        "Type": "CNAME",
        "TTL": 300,
        "ResourceRecords": [{"Value": "'$LB_HOSTNAME'"}]
      }
    }]
  }'
```

### 2. Configure SSL/TLS

```bash
# Request ACM certificate
aws acm request-certificate \
  --domain-name gopro.com \
  --subject-alternative-names "*.gopro.com" \
  --validation-method DNS \
  --region us-east-1

# Get certificate ARN
CERT_ARN=$(aws acm list-certificates --query 'CertificateSummaryList[?DomainName==`gopro.com`].CertificateArn' --output text)

# Update ingress with certificate ARN
kubectl annotate ingress gopro-ingress \
  alb.ingress.kubernetes.io/certificate-arn=$CERT_ARN \
  -n gopro
```

### 3. Configure Monitoring

```bash
# Access Grafana
kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring

# Default credentials:
# Username: admin
# Password: prom-operator

# Access Prometheus
kubectl port-forward svc/prometheus-kube-prometheus-prometheus 9090:9090 -n monitoring

# Access Jaeger
kubectl port-forward svc/jaeger-query 16686:16686 -n observability
```

## Monitoring & Logging

### CloudWatch Logs

```bash
# View cluster logs
aws logs tail /aws/eks/gopro-dev-cluster/cluster --follow

# View application logs
aws logs tail /gopro/application --follow

# Create log insights query
aws logs start-query \
  --log-group-name /gopro/application \
  --start-time $(date -u -d '1 hour ago' +%s) \
  --end-time $(date -u +%s) \
  --query-string 'fields @timestamp, @message | filter @message like /ERROR/ | sort @timestamp desc | limit 20'
```

### CloudWatch Metrics

```bash
# View EKS metrics
aws cloudwatch get-metric-statistics \
  --namespace AWS/EKS \
  --metric-name cluster_failed_node_count \
  --dimensions Name=ClusterName,Value=gopro-dev-cluster \
  --start-time $(date -u -d '1 hour ago' +%Y-%m-%dT%H:%M:%S) \
  --end-time $(date -u +%Y-%m-%dT%H:%M:%S) \
  --period 300 \
  --statistics Average
```

### Container Insights

```bash
# Install CloudWatch agent
kubectl apply -f https://raw.githubusercontent.com/aws-samples/amazon-cloudwatch-container-insights/latest/k8s-deployment-manifest-templates/deployment-mode/daemonset/container-insights-monitoring/quickstart/cwagent-fluentd-quickstart.yaml

# Verify installation
kubectl get pods -n amazon-cloudwatch
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
# Test database connectivity
kubectl run -it --rm debug --image=postgres:15 --restart=Never -- \
  psql -h <RDS_ENDPOINT> -U gopro -d gopro

# Check security groups
aws ec2 describe-security-groups \
  --filters "Name=tag:Name,Values=gopro-dev-db-sg"
```

#### 3. Load Balancer Not Created

```bash
# Check AWS Load Balancer Controller logs
kubectl logs -n kube-system deployment/aws-load-balancer-controller

# Check ingress events
kubectl describe ingress gopro-ingress -n gopro

# Verify IAM role
kubectl describe sa aws-load-balancer-controller -n kube-system
```

#### 4. High Costs

```bash
# Check resource usage
kubectl top nodes
kubectl top pods -n gopro

# Review AWS Cost Explorer
aws ce get-cost-and-usage \
  --time-period Start=2024-01-01,End=2024-01-31 \
  --granularity MONTHLY \
  --metrics BlendedCost \
  --group-by Type=SERVICE
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
kubectl set image deployment/backend backend=ghcr.io/dimajoyti/go-pro/backend:v1.1.0 -n gopro

# Rollback deployment
kubectl rollout undo deployment/backend -n gopro

# Get deployment history
kubectl rollout history deployment/backend -n gopro

# Execute command in pod
kubectl exec -it deployment/backend -n gopro -- /bin/sh

# Copy files from pod
kubectl cp gopro/<pod-name>:/path/to/file ./local-file

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
terraform destroy -var-file=environments/dev.tfvars

# Using eksctl
eksctl delete cluster --name gopro-dev-cluster --region us-east-1
```

## Additional Resources

- [AWS EKS Best Practices](https://aws.github.io/aws-eks-best-practices/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/)
- [Terraform AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

