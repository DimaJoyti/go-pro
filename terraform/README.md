# GO-PRO Learning Platform - Terraform Infrastructure

This directory contains Terraform configurations for deploying the GO-PRO learning platform infrastructure across AWS and GCP.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Architecture](#architecture)
- [Directory Structure](#directory-structure)
- [Getting Started](#getting-started)
- [Deployment](#deployment)
- [Modules](#modules)
- [Environments](#environments)
- [State Management](#state-management)
- [Security](#security)
- [Cost Optimization](#cost-optimization)
- [Troubleshooting](#troubleshooting)

## Overview

The infrastructure is designed to be:
- **Multi-Cloud**: Supports both AWS and GCP
- **Highly Available**: Multi-AZ deployment with automatic failover
- **Scalable**: Auto-scaling for compute resources
- **Secure**: Encryption at rest and in transit, WAF, GuardDuty
- **Cost-Optimized**: Environment-specific configurations
- **Observable**: Comprehensive monitoring and logging

## Prerequisites

### Required Tools

- [Terraform](https://www.terraform.io/downloads.html) >= 1.5.0
- [AWS CLI](https://aws.amazon.com/cli/) >= 2.0
- [kubectl](https://kubernetes.io/docs/tasks/tools/) >= 1.28
- [helm](https://helm.sh/docs/intro/install/) >= 3.0

### AWS Credentials

```bash
# Configure AWS credentials
aws configure

# Or use environment variables
export AWS_ACCESS_KEY_ID="your-access-key"
export AWS_SECRET_ACCESS_KEY="your-secret-key"
export AWS_DEFAULT_REGION="us-east-1"
```

### GCP Credentials

```bash
# Authenticate with GCP
gcloud auth application-default login

# Set project
gcloud config set project your-project-id
```

## Architecture

### AWS Infrastructure

```
┌─────────────────────────────────────────────────────────────┐
│                         AWS Cloud                            │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌──────────────────────────────────────────────────────┐   │
│  │                    VPC (10.0.0.0/16)                 │   │
│  │                                                       │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │   │
│  │  │   Public    │  │   Public    │  │   Public    │ │   │
│  │  │  Subnet 1   │  │  Subnet 2   │  │  Subnet 3   │ │   │
│  │  │   (ALB)     │  │   (ALB)     │  │   (ALB)     │ │   │
│  │  └─────────────┘  └─────────────┘  └─────────────┘ │   │
│  │                                                       │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │   │
│  │  │  Private    │  │  Private    │  │  Private    │ │   │
│  │  │  Subnet 1   │  │  Subnet 2   │  │  Subnet 3   │ │   │
│  │  │   (EKS)     │  │   (EKS)     │  │   (EKS)     │ │   │
│  │  └─────────────┘  └─────────────┘  └─────────────┘ │   │
│  │                                                       │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │   │
│  │  │  Database   │  │  Database   │  │  Database   │ │   │
│  │  │  Subnet 1   │  │  Subnet 2   │  │  Subnet 3   │ │   │
│  │  │ (RDS/Redis) │  │ (RDS/Redis) │  │ (RDS/Redis) │ │   │
│  │  └─────────────┘  └─────────────┘  └─────────────┘ │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │     RDS      │  │ ElastiCache  │  │     MSK      │      │
│  │  PostgreSQL  │  │    Redis     │  │    Kafka     │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│                                                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │     EKS      │  │     ALB      │  │      S3      │      │
│  │  Kubernetes  │  │Load Balancer │  │   Storage    │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
└─────────────────────────────────────────────────────────────┘
```

### Components

- **VPC**: Isolated network with public, private, and database subnets
- **EKS**: Managed Kubernetes cluster for application workloads
- **RDS**: PostgreSQL database with multi-AZ deployment
- **ElastiCache**: Redis cluster for caching and sessions
- **MSK**: Managed Kafka for event streaming
- **ALB**: Application Load Balancer with SSL termination
- **S3**: Object storage for static assets and backups
- **CloudWatch**: Monitoring, logging, and alerting
- **Secrets Manager**: Secure storage for sensitive data
- **WAF**: Web Application Firewall for security

## Directory Structure

```
terraform/
├── main.tf                 # Main configuration
├── variables.tf            # Variable definitions
├── outputs.tf              # Output definitions
├── versions.tf             # Provider versions
├── README.md               # This file
├── environments/           # Environment-specific configs
│   ├── dev.tfvars
│   ├── staging.tfvars
│   └── production.tfvars
└── modules/                # Reusable modules
    ├── vpc/
    ├── eks/
    ├── rds/
    ├── redis/
    ├── msk/
    ├── alb/
    ├── s3/
    ├── cloudwatch/
    ├── secrets/
    ├── iam/
    ├── security-groups/
    ├── route53/
    ├── acm/
    ├── waf/
    └── backup/
```

## Getting Started

### 1. Initialize Terraform

```bash
cd terraform
terraform init
```

### 2. Create State Backend

First, create the S3 bucket and DynamoDB table for state management:

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

# Create DynamoDB table for locking
aws dynamodb create-table \
  --table-name gopro-terraform-locks \
  --attribute-definitions AttributeName=LockID,AttributeType=S \
  --key-schema AttributeName=LockID,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --region us-east-1
```

### 3. Plan Deployment

```bash
# Development environment
terraform plan -var-file=environments/dev.tfvars

# Production environment
terraform plan -var-file=environments/production.tfvars
```

### 4. Apply Configuration

```bash
# Development environment
terraform apply -var-file=environments/dev.tfvars

# Production environment (with approval)
terraform apply -var-file=environments/production.tfvars
```

## Deployment

### Development Environment

```bash
# Initialize
terraform init

# Plan
terraform plan -var-file=environments/dev.tfvars -out=dev.tfplan

# Apply
terraform apply dev.tfplan

# Get outputs
terraform output
```

### Production Environment

```bash
# Initialize
terraform init

# Validate
terraform validate

# Plan
terraform plan -var-file=environments/production.tfvars -out=prod.tfplan

# Review plan carefully
terraform show prod.tfplan

# Apply with approval
terraform apply prod.tfplan

# Get outputs
terraform output
```

### Configure kubectl

```bash
# Get the kubectl config command from outputs
terraform output kubectl_config_command

# Run the command
aws eks update-kubeconfig --region us-east-1 --name gopro-production-eks

# Verify connection
kubectl get nodes
```

## Modules

### VPC Module
Creates VPC with public, private, and database subnets across multiple AZs.

### EKS Module
Provisions managed Kubernetes cluster with node groups and IRSA.

### RDS Module
Creates PostgreSQL database with automated backups and encryption.

### Redis Module
Sets up ElastiCache Redis cluster for caching.

### MSK Module
Provisions managed Kafka cluster for event streaming.

### ALB Module
Creates Application Load Balancer with SSL termination.

### Security Groups Module
Manages security groups for all components.

### IAM Module
Creates IAM roles and policies for services.

### Secrets Module
Stores sensitive data in AWS Secrets Manager.

### CloudWatch Module
Sets up monitoring, logging, and dashboards.

## Environments

### Development
- Minimal resources for cost optimization
- Single AZ deployment
- Smaller instance types
- Shorter backup retention

### Staging
- Production-like configuration
- Multi-AZ for testing
- Medium instance types
- Moderate backup retention

### Production
- High availability configuration
- Multi-AZ deployment
- Large instance types
- Extended backup retention
- Enhanced monitoring
- WAF and GuardDuty enabled

## State Management

Terraform state is stored in S3 with:
- **Versioning**: Enabled for state history
- **Encryption**: AES-256 encryption at rest
- **Locking**: DynamoDB table prevents concurrent modifications
- **Backup**: Automated backups with lifecycle policies

## Security

### Best Practices

1. **Encryption**: All data encrypted at rest and in transit
2. **Secrets**: Sensitive data stored in Secrets Manager
3. **IAM**: Least privilege access with IRSA
4. **Network**: Private subnets for databases and applications
5. **WAF**: Web Application Firewall for production
6. **GuardDuty**: Threat detection enabled
7. **Security Hub**: Centralized security findings

### Compliance

- SOC 2 compliant infrastructure
- GDPR-ready data handling
- HIPAA-eligible services (RDS, S3)
- PCI DSS Level 1 capable

## Cost Optimization

### Development
- ~$200-300/month
- Single NAT Gateway
- Smaller instance types
- Spot instances for non-critical workloads

### Production
- ~$1,500-2,500/month
- Reserved instances for predictable workloads
- Auto-scaling for variable loads
- S3 lifecycle policies

### Cost Reduction Tips

1. Use Spot instances for development
2. Enable auto-scaling
3. Right-size instances based on metrics
4. Use S3 lifecycle policies
5. Delete unused resources
6. Enable cost allocation tags

## Troubleshooting

### Common Issues

**Issue**: Terraform init fails
```bash
# Solution: Clear cache and reinitialize
rm -rf .terraform
terraform init
```

**Issue**: State lock error
```bash
# Solution: Force unlock (use with caution)
terraform force-unlock <lock-id>
```

**Issue**: Resource already exists
```bash
# Solution: Import existing resource
terraform import module.vpc.aws_vpc.main vpc-xxxxx
```

**Issue**: EKS cluster unreachable
```bash
# Solution: Update kubeconfig
aws eks update-kubeconfig --region us-east-1 --name cluster-name
```

### Getting Help

- Check [Terraform Documentation](https://www.terraform.io/docs)
- Review [AWS Provider Docs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
- Contact DevOps team

## Maintenance

### Regular Tasks

- Review and update Terraform versions
- Rotate secrets and credentials
- Review security group rules
- Optimize costs based on usage
- Update instance types as needed
- Review and update backup policies

### Disaster Recovery

- State backups in S3 with versioning
- Database automated backups
- Cross-region replication for critical data
- Documented recovery procedures

## Contributing

1. Create feature branch
2. Make changes
3. Run `terraform fmt`
4. Run `terraform validate`
5. Create pull request
6. Get approval
7. Merge to main

## License

Proprietary - GO-PRO Learning Platform
