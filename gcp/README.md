# GO-PRO GCP Cloud Integration

This directory contains Google Cloud Platform deployment configurations and scripts for the GO-PRO learning platform.

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                    Cloud DNS (DNS)                               │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                Cloud CDN + Cloud Armor                           │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│              Cloud Load Balancer (HTTPS)                         │
└────────────────────────┬────────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┬──────────────┐
         │               │               │              │
         ▼               ▼               ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│     GKE     │  │     GKE     │  │     GKE     │  │     GKE     │
│   Cluster   │  │   Cluster   │  │   Cluster   │  │   Cluster   │
│  (Zone-A)   │  │  (Zone-B)   │  │  (Zone-C)   │  │  (Zone-D)   │
└──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘
       │                │                │                │
       └────────────────┴────────────────┴────────────────┘
                         │
         ┌───────────────┼───────────────┬──────────────┐
         │               │               │              │
         ▼               ▼               ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│  Cloud SQL  │  │ Memorystore │  │   Pub/Sub   │  │Cloud Storage│
│ PostgreSQL  │  │    Redis    │  │   Topics    │  │   Buckets   │
│  (HA Mode)  │  │  Cluster    │  │  & Queues   │  │             │
└─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘
```

## GCP Services Used

### Compute
- **GKE (Google Kubernetes Engine)**: Managed Kubernetes
- **Compute Engine**: VM instances for node pools
- **Cloud Run**: Serverless containers (optional)

### Database
- **Cloud SQL PostgreSQL**: Managed PostgreSQL (HA)
- **Memorystore Redis**: Managed Redis
- **Firestore**: NoSQL database (optional)

### Messaging
- **Pub/Sub**: Message queuing and streaming
- **Cloud Tasks**: Task queue (optional)

### Storage
- **Cloud Storage**: Object storage
- **Persistent Disk**: Block storage
- **Filestore**: Managed NFS (optional)

### Networking
- **VPC**: Virtual Private Cloud
- **Cloud Load Balancing**: Global load balancer
- **Cloud DNS**: DNS management
- **Cloud CDN**: Content delivery network
- **Cloud Armor**: DDoS protection and WAF

### Security
- **IAM**: Identity and Access Management
- **Secret Manager**: Secrets storage
- **Cloud KMS**: Encryption key management
- **Binary Authorization**: Container image signing
- **Security Command Center**: Security findings

### Monitoring
- **Cloud Monitoring**: Metrics and dashboards
- **Cloud Logging**: Log aggregation
- **Cloud Trace**: Distributed tracing
- **Cloud Profiler**: Performance profiling
- **Error Reporting**: Error tracking

### CI/CD
- **Artifact Registry**: Container registry
- **Cloud Build**: Build service
- **Cloud Deploy**: Deployment automation

## Prerequisites

### Required Tools
- gcloud CLI
- kubectl 1.28+
- Helm 3.12+
- Terraform 1.6+

### GCP Project Setup
1. GCP project with billing enabled
2. Service account with appropriate permissions
3. gcloud CLI configured
4. APIs enabled (GKE, Cloud SQL, Memorystore, etc.)

## Quick Start

### 1. Configure gcloud CLI
```bash
# Install gcloud CLI
curl https://sdk.cloud.google.com | bash
exec -l $SHELL

# Initialize gcloud
gcloud init

# Set project
gcloud config set project gopro-project

# Authenticate
gcloud auth login
gcloud auth application-default login

# Verify configuration
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
  artifactregistry.googleapis.com
```

### 3. Deploy Infrastructure with Terraform
```bash
# Navigate to terraform directory
cd ../terraform

# Initialize Terraform
terraform init

# Plan deployment
terraform plan -var-file=environments/gcp-dev.tfvars

# Apply deployment
terraform apply -var-file=environments/gcp-dev.tfvars
```

### 4. Configure kubectl
```bash
# Get cluster credentials
gcloud container clusters get-credentials gopro-dev-cluster \
  --region us-central1 \
  --project gopro-project

# Verify connection
kubectl get nodes
```

### 5. Deploy Applications
```bash
# Deploy using kubectl
kubectl apply -k ../k8s/overlays/development

# Or use Helm
helm install gopro ./helm/gopro -f values-gcp-dev.yaml
```

## Deployment Environments

### Development
- **Cluster**: gopro-dev-cluster
- **Region**: us-central1
- **Node Count**: 2-5
- **Machine Type**: e2-medium
- **Database**: db-f1-micro
- **Redis**: M1 (1GB)
- **Cost**: ~$150-250/month

### Staging
- **Cluster**: gopro-staging-cluster
- **Region**: us-central1
- **Node Count**: 3-8
- **Machine Type**: e2-standard-4
- **Database**: db-n1-standard-2
- **Redis**: M3 (5GB)
- **Cost**: ~$400-700/month

### Production
- **Cluster**: gopro-prod-cluster
- **Region**: us-central1 (multi-region optional)
- **Node Count**: 5-20
- **Machine Type**: n2-standard-4
- **Database**: db-n1-standard-8 (HA)
- **Redis**: M5 (20GB, HA)
- **Cost**: ~$1,200-2,000/month

## Directory Structure

```
gcp/
├── README.md                    # This file
├── gke/
│   ├── cluster-config.yaml      # GKE cluster configuration
│   ├── nodepool-config.yaml     # Node pool configuration
│   └── workload-identity.yaml   # Workload Identity configuration
├── scripts/
│   ├── setup-cluster.sh         # Cluster setup script
│   ├── deploy-app.sh            # Application deployment
│   ├── backup-db.sh             # Database backup
│   ├── restore-db.sh            # Database restore
│   └── monitoring-setup.sh      # Monitoring setup
└── terraform/                   # GCP-specific Terraform (optional)
```

## GKE Cluster Setup

### Using gcloud
```bash
# Create GKE cluster
gcloud container clusters create gopro-dev-cluster \
  --region us-central1 \
  --num-nodes 2 \
  --machine-type e2-medium \
  --enable-autoscaling \
  --min-nodes 1 \
  --max-nodes 5 \
  --enable-autorepair \
  --enable-autoupgrade \
  --enable-ip-alias \
  --network gopro-vpc \
  --subnetwork gopro-subnet \
  --enable-stackdriver-kubernetes \
  --enable-cloud-logging \
  --enable-cloud-monitoring \
  --workload-pool=gopro-project.svc.id.goog

# Create node pool
gcloud container node-pools create gopro-pool \
  --cluster gopro-dev-cluster \
  --region us-central1 \
  --machine-type e2-medium \
  --num-nodes 2 \
  --enable-autoscaling \
  --min-nodes 1 \
  --max-nodes 5
```

### Using Terraform
```bash
# Already configured in terraform/modules/gke/
cd ../terraform
terraform apply -var-file=environments/gcp-dev.tfvars
```

## Database Setup

### Cloud SQL PostgreSQL
```bash
# Create Cloud SQL instance
gcloud sql instances create gopro-dev-db \
  --database-version=POSTGRES_15 \
  --tier=db-f1-micro \
  --region=us-central1 \
  --network=gopro-vpc \
  --no-assign-ip

# Create database
gcloud sql databases create gopro \
  --instance=gopro-dev-db

# Create user
gcloud sql users create gopro \
  --instance=gopro-dev-db \
  --password=<password>

# Get connection name
gcloud sql instances describe gopro-dev-db \
  --format='value(connectionName)'
```

### Memorystore Redis
```bash
# Create Redis instance
gcloud redis instances create gopro-dev-redis \
  --size=1 \
  --region=us-central1 \
  --network=gopro-vpc \
  --redis-version=redis_7_0

# Get Redis host
gcloud redis instances describe gopro-dev-redis \
  --region=us-central1 \
  --format='value(host)'
```

## Application Deployment

### Using kubectl
```bash
# Deploy to development
kubectl apply -k k8s/overlays/development

# Deploy to production
kubectl apply -k k8s/overlays/production

# Check deployment status
kubectl rollout status deployment/backend -n gopro
kubectl rollout status deployment/frontend -n gopro
```

### Using Cloud Build
```bash
# Submit build
gcloud builds submit \
  --config=cloudbuild.yaml \
  --substitutions=_ENVIRONMENT=development

# View build logs
gcloud builds log <BUILD_ID>
```

## Monitoring & Logging

### Cloud Monitoring
```bash
# Create dashboard
gcloud monitoring dashboards create --config-from-file=dashboard.json

# Create alert policy
gcloud alpha monitoring policies create \
  --notification-channels=<CHANNEL_ID> \
  --display-name="High CPU Usage" \
  --condition-display-name="CPU > 80%" \
  --condition-threshold-value=0.8
```

### Cloud Logging
```bash
# View logs
gcloud logging read "resource.type=k8s_cluster" --limit 50

# Create log sink
gcloud logging sinks create gopro-logs \
  storage.googleapis.com/gopro-logs-bucket \
  --log-filter='resource.type="k8s_cluster"'
```

## Security Best Practices

### Workload Identity
- Use Workload Identity for pod authentication
- Bind Kubernetes service accounts to GCP service accounts
- Avoid using service account keys

### Network Security
- Private GKE cluster
- VPC-native cluster
- Authorized networks for API access
- Cloud Armor for DDoS protection

### Secrets Management
- Use Secret Manager for secrets
- Rotate secrets regularly
- Never commit secrets to Git
- Use encryption at rest

### Compliance
- Enable Binary Authorization
- Enable Security Command Center
- Regular security scans
- Audit logging enabled

## Backup & Disaster Recovery

### Database Backups
```bash
# Automated backups (configured in Terraform)
# Retention: 7 days (dev), 30 days (prod)

# Manual backup
gcloud sql backups create \
  --instance=gopro-prod-db \
  --description="Manual backup $(date +%Y%m%d)"

# Restore from backup
gcloud sql backups restore <BACKUP_ID> \
  --backup-instance=gopro-prod-db \
  --backup-id=<BACKUP_ID>
```

### Application Backups
```bash
# Backup to Cloud Storage
kubectl exec deployment/backend -- ./backup.sh | \
  gsutil cp - gs://gopro-backups/app-$(date +%Y%m%d).tar.gz

# Restore from Cloud Storage
gsutil cp gs://gopro-backups/app-20240107.tar.gz - | \
  kubectl exec -i deployment/backend -- ./restore.sh
```

## Cost Optimization

### Recommendations
1. Use Preemptible VMs for non-critical workloads
2. Enable cluster autoscaling
3. Use Cloud Storage lifecycle policies
4. Enable Cloud SQL auto-scaling
5. Use Committed Use Discounts
6. Monitor with Cloud Billing reports
7. Set up budget alerts

### Cost Monitoring
```bash
# Export billing data
gcloud beta billing accounts list

# Create budget
gcloud beta billing budgets create \
  --billing-account=<BILLING_ACCOUNT_ID> \
  --display-name="GO-PRO Monthly Budget" \
  --budget-amount=2000
```

## Troubleshooting

### GKE Issues
```bash
# Check cluster status
gcloud container clusters describe gopro-dev-cluster \
  --region us-central1

# Check node pool
gcloud container node-pools describe gopro-pool \
  --cluster gopro-dev-cluster \
  --region us-central1

# View cluster events
kubectl get events --all-namespaces --sort-by='.lastTimestamp'
```

### Database Issues
```bash
# Check Cloud SQL status
gcloud sql instances describe gopro-dev-db

# View logs
gcloud sql operations list --instance=gopro-dev-db

# Connect to database
gcloud sql connect gopro-dev-db --user=gopro
```

### Networking Issues
```bash
# Check VPC
gcloud compute networks describe gopro-vpc

# Check firewall rules
gcloud compute firewall-rules list --filter="network:gopro-vpc"

# Check routes
gcloud compute routes list --filter="network:gopro-vpc"
```

## Maintenance

### Regular Tasks
- Weekly: Review Cloud Monitoring dashboards
- Weekly: Check security findings
- Monthly: Review IAM permissions
- Monthly: Update GKE cluster
- Quarterly: Review costs
- Quarterly: Security audit

### Updates
```bash
# Update GKE cluster
gcloud container clusters upgrade gopro-dev-cluster \
  --region us-central1 \
  --master

# Update node pool
gcloud container clusters upgrade gopro-dev-cluster \
  --region us-central1 \
  --node-pool gopro-pool
```

## Additional Resources

- [GKE Documentation](https://cloud.google.com/kubernetes-engine/docs)
- [Cloud SQL Documentation](https://cloud.google.com/sql/docs)
- [Memorystore Documentation](https://cloud.google.com/memorystore/docs)
- [Cloud Monitoring Documentation](https://cloud.google.com/monitoring/docs)
- [GCP Best Practices](https://cloud.google.com/architecture/framework)

