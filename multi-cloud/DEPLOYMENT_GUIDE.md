# GO-PRO Multi-Cloud Deployment Guide

This guide provides step-by-step instructions for deploying the GO-PRO learning platform across multiple cloud providers (AWS and GCP).

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Architecture Overview](#architecture-overview)
3. [Deployment Strategy](#deployment-strategy)
4. [Infrastructure Setup](#infrastructure-setup)
5. [Application Deployment](#application-deployment)
6. [DNS Configuration](#dns-configuration)
7. [Data Replication](#data-replication)
8. [Monitoring Setup](#monitoring-setup)
9. [Failover Testing](#failover-testing)
10. [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Tools

```bash
# AWS CLI
aws --version  # 2.x required

# gcloud CLI
gcloud version  # Latest version

# kubectl
kubectl version --client  # 1.28+

# Terraform
terraform version  # 1.6+

# Helm
helm version  # 3.12+
```

### Cloud Accounts

- AWS account with appropriate permissions
- GCP project with billing enabled
- Domain name for DNS configuration
- SSL certificates (or Let's Encrypt)

## Architecture Overview

### Deployment Model: Active-Passive

**Primary**: AWS (us-east-1)
- Serves 100% of production traffic
- Full application stack
- Primary database (RDS PostgreSQL)
- Primary cache (ElastiCache Redis)

**Secondary**: GCP (us-central1)
- Hot standby (ready but idle)
- Full application stack
- Read replica database (Cloud SQL)
- Secondary cache (Memorystore Redis)

**Failover**: Automatic
- DNS-based failover (Route 53 + Cloud DNS)
- RTO: < 5 minutes
- RPO: < 1 minute

## Deployment Strategy

### Phase 1: Deploy to AWS (Primary)

```bash
# 1. Deploy infrastructure
cd terraform
terraform apply -var-file=environments/aws-prod.tfvars

# 2. Setup EKS cluster
cd ../aws/scripts
./setup-cluster.sh

# 3. Deploy application
ENVIRONMENT=production ./deploy-app.sh

# 4. Verify deployment
kubectl get pods -n gopro
kubectl get svc -n gopro
```

### Phase 2: Deploy to GCP (Secondary)

```bash
# 1. Deploy infrastructure
cd terraform
terraform apply -var-file=environments/gcp-prod.tfvars

# 2. Setup GKE cluster
cd ../gcp/scripts
./setup-cluster.sh

# 3. Deploy application
ENVIRONMENT=production ./deploy-app.sh

# 4. Verify deployment
kubectl get pods -n gopro
kubectl get svc -n gopro
```

### Phase 3: Configure Multi-Cloud

```bash
# 1. Deploy to both clouds
cd multi-cloud/scripts
./deploy-multi-cloud.sh

# 2. Configure DNS routing
# See DNS Configuration section

# 3. Setup data replication
# See Data Replication section

# 4. Configure monitoring
# See Monitoring Setup section
```

## Infrastructure Setup

### AWS Infrastructure

```bash
# Create VPC
aws ec2 create-vpc --cidr-block 10.0.0.0/16

# Create EKS cluster
eksctl create cluster -f aws/eks/cluster-config.yaml

# Create RDS instance
aws rds create-db-instance \
  --db-instance-identifier gopro-prod-db \
  --db-instance-class db.r6g.xlarge \
  --engine postgres \
  --master-username gopro \
  --master-user-password <password> \
  --allocated-storage 100 \
  --multi-az

# Create ElastiCache cluster
aws elasticache create-cache-cluster \
  --cache-cluster-id gopro-prod-redis \
  --cache-node-type cache.r6g.xlarge \
  --engine redis \
  --num-cache-nodes 1
```

### GCP Infrastructure

```bash
# Create VPC
gcloud compute networks create gopro-vpc --subnet-mode=custom

# Create GKE cluster
gcloud container clusters create gopro-prod-cluster \
  --region us-central1 \
  --num-nodes 3 \
  --machine-type n2-standard-4

# Create Cloud SQL instance
gcloud sql instances create gopro-prod-db \
  --database-version=POSTGRES_15 \
  --tier=db-n1-standard-8 \
  --region=us-central1 \
  --availability-type=REGIONAL

# Create Memorystore instance
gcloud redis instances create gopro-prod-redis \
  --size=20 \
  --region=us-central1 \
  --tier=standard
```

## Application Deployment

### Deploy to AWS

```bash
# Get EKS credentials
aws eks update-kubeconfig --name gopro-prod-cluster --region us-east-1

# Create secrets
kubectl create secret generic database-credentials \
  --from-literal=host=<rds-endpoint> \
  --from-literal=port=5432 \
  --from-literal=username=gopro \
  --from-literal=password=<password> \
  --from-literal=database=gopro \
  -n gopro

# Deploy application
kubectl apply -k k8s/overlays/production

# Verify
kubectl get pods -n gopro
kubectl rollout status deployment/backend -n gopro
```

### Deploy to GCP

```bash
# Get GKE credentials
gcloud container clusters get-credentials gopro-prod-cluster --region us-central1

# Create secrets
kubectl create secret generic database-credentials \
  --from-literal=host=127.0.0.1 \
  --from-literal=port=5432 \
  --from-literal=username=gopro \
  --from-literal=password=<password> \
  --from-literal=database=gopro \
  -n gopro

# Deploy application
kubectl apply -k k8s/overlays/production

# Verify
kubectl get pods -n gopro
kubectl rollout status deployment/backend -n gopro
```

## DNS Configuration

### Route 53 (AWS)

```bash
# Create hosted zone
aws route53 create-hosted-zone --name gopro.com --caller-reference $(date +%s)

# Create health check
aws route53 create-health-check \
  --type HTTPS \
  --resource-path /health \
  --fully-qualified-domain-name api.gopro.com \
  --port 443 \
  --request-interval 30 \
  --failure-threshold 3

# Create failover record (Primary)
aws route53 change-resource-record-sets \
  --hosted-zone-id <zone-id> \
  --change-batch '{
    "Changes": [{
      "Action": "CREATE",
      "ResourceRecordSet": {
        "Name": "gopro.com",
        "Type": "A",
        "SetIdentifier": "AWS-Primary",
        "Failover": "PRIMARY",
        "AliasTarget": {
          "HostedZoneId": "<alb-zone-id>",
          "DNSName": "<alb-dns-name>",
          "EvaluateTargetHealth": true
        },
        "HealthCheckId": "<health-check-id>"
      }
    }]
  }'

# Create failover record (Secondary)
aws route53 change-resource-record-sets \
  --hosted-zone-id <zone-id> \
  --change-batch '{
    "Changes": [{
      "Action": "CREATE",
      "ResourceRecordSet": {
        "Name": "gopro.com",
        "Type": "A",
        "SetIdentifier": "GCP-Secondary",
        "Failover": "SECONDARY",
        "TTL": 60,
        "ResourceRecords": [{"Value": "<gcp-lb-ip>"}]
      }
    }]
  }'
```

### Cloud DNS (GCP)

```bash
# Create managed zone
gcloud dns managed-zones create gopro-zone \
  --dns-name=gopro.com \
  --description="GO-PRO DNS Zone"

# Add A record
gcloud dns record-sets transaction start --zone=gopro-zone

gcloud dns record-sets transaction add <gcp-lb-ip> \
  --name=gopro.com \
  --ttl=60 \
  --type=A \
  --zone=gopro-zone

gcloud dns record-sets transaction execute --zone=gopro-zone
```

## Data Replication

### Database Replication (RDS → Cloud SQL)

```bash
# 1. Create read replica in GCP
# Note: This requires VPN or private interconnect between AWS and GCP

# 2. Configure Cloud SQL as read replica
gcloud sql instances create gopro-prod-db-replica \
  --master-instance-name=<aws-rds-instance> \
  --tier=db-n1-standard-8 \
  --region=us-central1

# 3. Monitor replication lag
gcloud sql operations list --instance=gopro-prod-db-replica
```

### Object Storage Sync (S3 ↔ Cloud Storage)

```bash
# Install gsutil
pip install gsutil

# Configure credentials
gcloud auth application-default login
aws configure

# Create sync script
cat > /usr/local/bin/sync-storage.sh <<'EOF'
#!/bin/bash
# Sync S3 to Cloud Storage
gsutil -m rsync -r -d s3://gopro-assets gs://gopro-assets

# Sync Cloud Storage to S3
aws s3 sync gs://gopro-assets s3://gopro-assets --delete
EOF

chmod +x /usr/local/bin/sync-storage.sh

# Add to cron (every 5 minutes)
echo "*/5 * * * * /usr/local/bin/sync-storage.sh" | crontab -
```

### Redis Replication

```bash
# Option 1: Application-level replication
# Write to both Redis instances from application

# Option 2: Redis replication
# Configure ElastiCache as master, Memorystore as replica
# Note: Requires VPN or private interconnect
```

## Monitoring Setup

### Prometheus Federation

```yaml
# prometheus-federation.yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'federate-aws'
    scrape_interval: 30s
    honor_labels: true
    metrics_path: '/federate'
    params:
      'match[]':
        - '{job=~".+"}'
    static_configs:
      - targets:
        - 'prometheus-aws.gopro.com:9090'
        labels:
          cloud: 'aws'
          region: 'us-east-1'

  - job_name: 'federate-gcp'
    scrape_interval: 30s
    honor_labels: true
    metrics_path: '/federate'
    params:
      'match[]':
        - '{job=~".+"}'
    static_configs:
      - targets:
        - 'prometheus-gcp.gopro.com:9090'
        labels:
          cloud: 'gcp'
          region: 'us-central1'
```

### Unified Dashboards

```bash
# Deploy Grafana
helm install grafana grafana/grafana \
  --set persistence.enabled=true \
  --set adminPassword=admin

# Import dashboards
# - Multi-Cloud Overview
# - AWS Metrics
# - GCP Metrics
# - Failover Status
# - Cost Analysis
```

## Failover Testing

### Manual Failover Test

```bash
# Run failover test
cd multi-cloud/scripts
./test-failover.sh

# Test steps:
# 1. Verify normal operation
# 2. Simulate primary failure
# 3. Verify secondary takeover
# 4. Restore primary
# 5. Verify failback
```

### Automated Failover

```yaml
# failover-policy.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: failover-policy
data:
  triggers: |
    - type: health_check_failure
      threshold: 3
      duration: 90s
      action: failover_to_secondary
    
    - type: high_error_rate
      threshold: 5
      duration: 300s
      action: gradual_shift
  
  rollback: |
    automatic: true
    conditions:
      - primary_healthy_for: 600s
      - error_rate_below: 1
```

## Troubleshooting

### Common Issues

#### 1. DNS Failover Not Working

```bash
# Check health checks
aws route53 get-health-check-status --health-check-id <id>

# Verify DNS propagation
dig gopro.com
nslookup gopro.com

# Check TTL
dig gopro.com +noall +answer
```

#### 2. Database Replication Lag

```bash
# Check replication status (AWS)
aws rds describe-db-instances \
  --db-instance-identifier gopro-prod-db \
  --query 'DBInstances[0].ReadReplicaDBInstanceIdentifiers'

# Check replication lag (GCP)
gcloud sql operations list --instance=gopro-prod-db-replica
```

#### 3. Storage Sync Issues

```bash
# Check sync status
gsutil ls -L s3://gopro-assets
aws s3 ls s3://gopro-assets --recursive

# Manual sync
gsutil -m rsync -r -d s3://gopro-assets gs://gopro-assets
```

### Useful Commands

```bash
# Switch between clusters
aws eks update-kubeconfig --name gopro-prod-cluster --region us-east-1
gcloud container clusters get-credentials gopro-prod-cluster --region us-central1

# Check deployment status
kubectl get pods -n gopro --context aws
kubectl get pods -n gopro --context gcp

# View logs
kubectl logs -f deployment/backend -n gopro --context aws
kubectl logs -f deployment/backend -n gopro --context gcp

# Test endpoints
curl -I https://gopro.com/health
curl -I https://api.gopro.com/health
```

## Best Practices

1. **Regular Failover Drills**: Test failover monthly
2. **Monitor Replication Lag**: Keep lag < 5 seconds
3. **Cost Optimization**: Review costs weekly
4. **Security Audits**: Quarterly security reviews
5. **Documentation**: Keep runbooks updated
6. **Automation**: Automate everything possible
7. **Monitoring**: 24/7 monitoring and alerting
8. **Backup**: Regular backups to both clouds

## Additional Resources

- [AWS Multi-Region Architecture](https://aws.amazon.com/solutions/implementations/multi-region-application-architecture/)
- [GCP Multi-Region Deployment](https://cloud.google.com/architecture/deploying-multi-region-web-applications)
- [Terraform Multi-Cloud](https://www.terraform.io/docs/cloud/guides/recommended-practices/part3.html)
- [Kubernetes Federation](https://kubernetes.io/docs/concepts/cluster-administration/federation/)

