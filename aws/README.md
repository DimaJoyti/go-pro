# GO-PRO AWS Cloud Integration

This directory contains AWS deployment configurations and scripts for the GO-PRO learning platform.

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                         Route 53 (DNS)                           │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                    CloudFront (CDN)                              │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                Application Load Balancer                         │
│                    (ALB + WAF)                                   │
└────────────────────────┬────────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┬──────────────┐
         │               │               │              │
         ▼               ▼               ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│     EKS     │  │     EKS     │  │     EKS     │  │     EKS     │
│   Cluster   │  │   Cluster   │  │   Cluster   │  │   Cluster   │
│   (AZ-1)    │  │   (AZ-2)    │  │   (AZ-3)    │  │   (AZ-4)    │
└──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘
       │                │                │                │
       └────────────────┴────────────────┴────────────────┘
                         │
         ┌───────────────┼───────────────┬──────────────┐
         │               │               │              │
         ▼               ▼               ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│     RDS     │  │ElastiCache  │  │     MSK     │  │     S3      │
│ PostgreSQL  │  │    Redis    │  │    Kafka    │  │   Storage   │
│ Multi-AZ    │  │  Cluster    │  │   Cluster   │  │   Buckets   │
└─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘
```

## AWS Services Used

### Compute
- **EKS (Elastic Kubernetes Service)**: Container orchestration
- **EC2**: Worker nodes for EKS
- **Fargate**: Serverless compute for Kubernetes pods (optional)

### Database
- **RDS PostgreSQL**: Primary database (Multi-AZ)
- **ElastiCache Redis**: Caching and session storage
- **DynamoDB**: Terraform state locking

### Messaging
- **MSK (Managed Streaming for Kafka)**: Event streaming
- **SQS**: Message queuing (optional)
- **SNS**: Notifications

### Storage
- **S3**: Object storage for assets, backups, logs
- **EBS**: Persistent volumes for Kubernetes
- **EFS**: Shared file system (optional)

### Networking
- **VPC**: Virtual Private Cloud
- **ALB**: Application Load Balancer
- **Route 53**: DNS management
- **CloudFront**: CDN for static assets

### Security
- **IAM**: Identity and Access Management
- **Secrets Manager**: Secrets storage
- **KMS**: Encryption key management
- **WAF**: Web Application Firewall
- **GuardDuty**: Threat detection
- **Security Hub**: Security findings aggregation

### Monitoring
- **CloudWatch**: Logs, metrics, alarms
- **X-Ray**: Distributed tracing
- **CloudTrail**: API audit logging

### CI/CD
- **ECR**: Container registry
- **CodePipeline**: CI/CD orchestration (optional)
- **CodeBuild**: Build service (optional)

## Prerequisites

### Required Tools
- AWS CLI 2.x
- kubectl 1.28+
- eksctl 0.160+
- Helm 3.12+
- Terraform 1.6+

### AWS Account Setup
1. AWS account with appropriate permissions
2. IAM user with programmatic access
3. AWS CLI configured with credentials
4. S3 bucket for Terraform state
5. DynamoDB table for state locking

## Quick Start

### 1. Configure AWS Credentials
```bash
# Configure AWS CLI
aws configure

# Verify configuration
aws sts get-caller-identity
```

### 2. Deploy Infrastructure with Terraform
```bash
# Navigate to terraform directory
cd ../terraform

# Initialize Terraform
terraform init

# Plan deployment
terraform plan -var-file=environments/dev.tfvars

# Apply deployment
terraform apply -var-file=environments/dev.tfvars
```

### 3. Configure kubectl
```bash
# Update kubeconfig
aws eks update-kubeconfig --name gopro-dev-cluster --region us-east-1

# Verify connection
kubectl get nodes
```

### 4. Deploy Applications
```bash
# Deploy using kubectl
kubectl apply -k ../k8s/overlays/development

# Or use Helm
helm install gopro ./helm/gopro -f values-dev.yaml
```

## Deployment Environments

### Development
- **Cluster**: gopro-dev-cluster
- **Region**: us-east-1
- **Node Count**: 2-5
- **Instance Type**: t3.medium
- **Database**: db.t3.micro
- **Redis**: cache.t4g.micro
- **Cost**: ~$200-300/month

### Staging
- **Cluster**: gopro-staging-cluster
- **Region**: us-east-1
- **Node Count**: 3-8
- **Instance Type**: t3.large
- **Database**: db.r6g.large
- **Redis**: cache.r6g.large
- **Cost**: ~$500-800/month

### Production
- **Cluster**: gopro-prod-cluster
- **Region**: us-east-1
- **Node Count**: 5-20
- **Instance Type**: c6i.xlarge
- **Database**: db.r6g.xlarge (Multi-AZ)
- **Redis**: cache.r6g.xlarge (Cluster mode)
- **Cost**: ~$1,500-2,500/month

## Directory Structure

```
aws/
├── README.md                    # This file
├── eks/
│   ├── cluster-config.yaml      # EKS cluster configuration
│   ├── nodegroup-config.yaml    # Node group configuration
│   └── addons.yaml              # EKS addons configuration
├── scripts/
│   ├── setup-cluster.sh         # Cluster setup script
│   ├── deploy-app.sh            # Application deployment
│   ├── backup-db.sh             # Database backup
│   ├── restore-db.sh            # Database restore
│   └── monitoring-setup.sh      # Monitoring setup
└── cloudformation/              # CloudFormation templates (optional)
```

## EKS Cluster Setup

### Using eksctl
```bash
# Create cluster
eksctl create cluster -f eks/cluster-config.yaml

# Create node group
eksctl create nodegroup -f eks/nodegroup-config.yaml

# Install addons
eksctl create addon -f eks/addons.yaml
```

### Using Terraform
```bash
# Already configured in terraform/modules/eks/
cd ../terraform
terraform apply -var-file=environments/dev.tfvars
```

## Database Setup

### RDS PostgreSQL
```bash
# Get database endpoint
aws rds describe-db-instances \
  --db-instance-identifier gopro-dev-db \
  --query 'DBInstances[0].Endpoint.Address' \
  --output text

# Connect to database
psql -h <endpoint> -U gopro -d gopro

# Run migrations
kubectl exec -it deployment/backend -- ./migrate up
```

### ElastiCache Redis
```bash
# Get Redis endpoint
aws elasticache describe-cache-clusters \
  --cache-cluster-id gopro-dev-redis \
  --show-cache-node-info \
  --query 'CacheClusters[0].CacheNodes[0].Endpoint.Address' \
  --output text

# Test connection
redis-cli -h <endpoint> ping
```

## Application Deployment

### Using kubectl
```bash
# Deploy to development
kubectl apply -k k8s/overlays/development

# Deploy to production
kubectl apply -k k8s/overlays/production

# Check deployment status
kubectl rollout status deployment/backend
kubectl rollout status deployment/frontend
```

### Using Helm
```bash
# Add Helm repository (if using)
helm repo add gopro https://charts.gopro.com

# Install/upgrade
helm upgrade --install gopro ./helm/gopro \
  -f values-dev.yaml \
  --namespace gopro \
  --create-namespace

# Check status
helm status gopro -n gopro
```

## Monitoring & Logging

### CloudWatch
```bash
# View logs
aws logs tail /aws/eks/gopro-dev-cluster/cluster --follow

# Create log group
aws logs create-log-group --log-group-name /gopro/application

# Put metric data
aws cloudwatch put-metric-data \
  --namespace GOPRO \
  --metric-name RequestCount \
  --value 1
```

### Container Insights
```bash
# Install CloudWatch agent
kubectl apply -f https://raw.githubusercontent.com/aws-samples/amazon-cloudwatch-container-insights/latest/k8s-deployment-manifest-templates/deployment-mode/daemonset/container-insights-monitoring/quickstart/cwagent-fluentd-quickstart.yaml
```

## Security Best Practices

### IAM Roles
- Use IRSA (IAM Roles for Service Accounts)
- Principle of least privilege
- Separate roles for each service
- Regular access reviews

### Network Security
- Private subnets for databases
- Security groups with minimal access
- Network ACLs for additional protection
- VPC Flow Logs enabled

### Secrets Management
- Use AWS Secrets Manager
- Rotate secrets regularly
- Never commit secrets to Git
- Use encryption at rest

### Compliance
- Enable CloudTrail
- Enable GuardDuty
- Enable Security Hub
- Regular security audits

## Backup & Disaster Recovery

### Database Backups
```bash
# Automated backups (configured in Terraform)
# Retention: 7 days (dev), 30 days (prod)

# Manual snapshot
aws rds create-db-snapshot \
  --db-instance-identifier gopro-prod-db \
  --db-snapshot-identifier gopro-manual-$(date +%Y%m%d)

# Restore from snapshot
aws rds restore-db-instance-from-db-snapshot \
  --db-instance-identifier gopro-restored \
  --db-snapshot-identifier gopro-manual-20240107
```

### Application Backups
```bash
# Backup to S3
kubectl exec deployment/backend -- ./backup.sh | \
  aws s3 cp - s3://gopro-backups/app-$(date +%Y%m%d).tar.gz

# Restore from S3
aws s3 cp s3://gopro-backups/app-20240107.tar.gz - | \
  kubectl exec -i deployment/backend -- ./restore.sh
```

## Cost Optimization

### Recommendations
1. Use Spot Instances for non-critical workloads
2. Enable auto-scaling for EKS nodes
3. Use S3 lifecycle policies
4. Enable RDS auto-scaling
5. Use Reserved Instances for production
6. Monitor with AWS Cost Explorer
7. Set up billing alerts

### Cost Monitoring
```bash
# Get cost and usage
aws ce get-cost-and-usage \
  --time-period Start=2024-01-01,End=2024-01-31 \
  --granularity MONTHLY \
  --metrics BlendedCost

# Create budget
aws budgets create-budget \
  --account-id <account-id> \
  --budget file://budget.json
```

## Troubleshooting

### EKS Issues
```bash
# Check cluster status
aws eks describe-cluster --name gopro-dev-cluster

# Check node group
aws eks describe-nodegroup \
  --cluster-name gopro-dev-cluster \
  --nodegroup-name gopro-dev-nodes

# View cluster logs
aws eks list-updates --name gopro-dev-cluster
```

### Database Issues
```bash
# Check DB status
aws rds describe-db-instances \
  --db-instance-identifier gopro-dev-db

# View DB logs
aws rds describe-db-log-files \
  --db-instance-identifier gopro-dev-db

# Download logs
aws rds download-db-log-file-portion \
  --db-instance-identifier gopro-dev-db \
  --log-file-name error/postgresql.log
```

### Networking Issues
```bash
# Check VPC
aws ec2 describe-vpcs --filters "Name=tag:Name,Values=gopro-dev-vpc"

# Check security groups
aws ec2 describe-security-groups \
  --filters "Name=tag:Environment,Values=development"

# Check route tables
aws ec2 describe-route-tables \
  --filters "Name=tag:Environment,Values=development"
```

## Maintenance

### Regular Tasks
- Weekly: Review CloudWatch alarms
- Weekly: Check security findings
- Monthly: Review IAM permissions
- Monthly: Update EKS cluster
- Quarterly: Review costs
- Quarterly: Security audit

### Updates
```bash
# Update EKS cluster
eksctl upgrade cluster --name gopro-dev-cluster --approve

# Update node group
eksctl upgrade nodegroup \
  --cluster gopro-dev-cluster \
  --name gopro-dev-nodes

# Update addons
eksctl update addon \
  --cluster gopro-dev-cluster \
  --name vpc-cni \
  --version latest
```

## Additional Resources

- [AWS EKS Documentation](https://docs.aws.amazon.com/eks/)
- [AWS RDS Documentation](https://docs.aws.amazon.com/rds/)
- [AWS ElastiCache Documentation](https://docs.aws.amazon.com/elasticache/)
- [AWS MSK Documentation](https://docs.aws.amazon.com/msk/)
- [eksctl Documentation](https://eksctl.io/)
- [Kubernetes on AWS](https://aws.amazon.com/kubernetes/)

