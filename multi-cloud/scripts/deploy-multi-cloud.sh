#!/bin/bash

# GO-PRO Multi-Cloud Deployment Script
# This script deploys the application to both AWS and GCP

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
ENVIRONMENT="${ENVIRONMENT:-production}"
AWS_REGION="${AWS_REGION:-us-east-1}"
GCP_REGION="${GCP_REGION:-us-central1}"
GCP_PROJECT="${GCP_PROJECT:-gopro-project}"
DEPLOY_AWS="${DEPLOY_AWS:-true}"
DEPLOY_GCP="${DEPLOY_GCP:-true}"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}GO-PRO Multi-Cloud Deployment${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Environment: $ENVIRONMENT"
echo "AWS Region: $AWS_REGION"
echo "GCP Region: $GCP_REGION"
echo "Deploy to AWS: $DEPLOY_AWS"
echo "Deploy to GCP: $DEPLOY_GCP"
echo ""

# Check prerequisites
echo -e "${YELLOW}Checking prerequisites...${NC}"
command -v aws >/dev/null 2>&1 || { echo -e "${RED}AWS CLI is required${NC}" >&2; exit 1; }
command -v gcloud >/dev/null 2>&1 || { echo -e "${RED}gcloud CLI is required${NC}" >&2; exit 1; }
command -v kubectl >/dev/null 2>&1 || { echo -e "${RED}kubectl is required${NC}" >&2; exit 1; }
command -v terraform >/dev/null 2>&1 || { echo -e "${RED}Terraform is required${NC}" >&2; exit 1; }
echo -e "${GREEN}✓ All prerequisites met${NC}"

# Deploy to AWS
if [ "$DEPLOY_AWS" = "true" ]; then
    echo ""
    echo -e "${BLUE}========================================${NC}"
    echo -e "${BLUE}Deploying to AWS${NC}"
    echo -e "${BLUE}========================================${NC}"
    
    # Configure AWS
    echo -e "${YELLOW}Configuring AWS...${NC}"
    aws configure set region "$AWS_REGION"
    
    # Update kubeconfig for EKS
    echo -e "${YELLOW}Updating kubeconfig for EKS...${NC}"
    aws eks update-kubeconfig --name gopro-prod-cluster --region "$AWS_REGION"
    
    # Deploy application
    echo -e "${YELLOW}Deploying application to AWS...${NC}"
    cd ../../aws/scripts
    ENVIRONMENT="$ENVIRONMENT" ./deploy-app.sh
    
    # Verify deployment
    echo -e "${YELLOW}Verifying AWS deployment...${NC}"
    kubectl get pods -n gopro
    kubectl get svc -n gopro
    
    echo -e "${GREEN}✓ AWS deployment complete${NC}"
fi

# Deploy to GCP
if [ "$DEPLOY_GCP" = "true" ]; then
    echo ""
    echo -e "${BLUE}========================================${NC}"
    echo -e "${BLUE}Deploying to GCP${NC}"
    echo -e "${BLUE}========================================${NC}"
    
    # Configure GCP
    echo -e "${YELLOW}Configuring GCP...${NC}"
    gcloud config set project "$GCP_PROJECT"
    
    # Update kubeconfig for GKE
    echo -e "${YELLOW}Updating kubeconfig for GKE...${NC}"
    gcloud container clusters get-credentials gopro-prod-cluster --region "$GCP_REGION"
    
    # Deploy application
    echo -e "${YELLOW}Deploying application to GCP...${NC}"
    cd ../../gcp/scripts
    ENVIRONMENT="$ENVIRONMENT" ./deploy-app.sh
    
    # Verify deployment
    echo -e "${YELLOW}Verifying GCP deployment...${NC}"
    kubectl get pods -n gopro
    kubectl get svc -n gopro
    
    echo -e "${GREEN}✓ GCP deployment complete${NC}"
fi

# Setup cross-cloud replication
echo ""
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Setting up cross-cloud replication${NC}"
echo -e "${BLUE}========================================${NC}"

# Database replication
echo -e "${YELLOW}Configuring database replication...${NC}"
echo "Note: Database replication should be configured manually"
echo "AWS RDS → GCP Cloud SQL read replica"
echo ""

# Object storage sync
echo -e "${YELLOW}Setting up object storage sync...${NC}"
cat > /tmp/sync-storage.sh <<'EOF'
#!/bin/bash
# Sync S3 to Cloud Storage
gsutil -m rsync -r -d s3://gopro-assets gs://gopro-assets
# Sync Cloud Storage to S3
aws s3 sync gs://gopro-assets s3://gopro-assets --delete
EOF
chmod +x /tmp/sync-storage.sh
echo "Storage sync script created: /tmp/sync-storage.sh"
echo "Add to cron: */5 * * * * /tmp/sync-storage.sh"
echo ""

# Setup DNS routing
echo -e "${YELLOW}Configuring DNS routing...${NC}"
echo "Note: DNS routing should be configured in Route 53 and Cloud DNS"
echo "Configure geolocation routing, latency-based routing, or weighted routing"
echo ""

# Setup monitoring
echo -e "${YELLOW}Setting up unified monitoring...${NC}"
echo "Note: Configure Prometheus federation to aggregate metrics from both clouds"
echo ""

# Get deployment information
echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Multi-Cloud Deployment Complete!${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""

if [ "$DEPLOY_AWS" = "true" ]; then
    echo "AWS Deployment:"
    aws eks update-kubeconfig --name gopro-prod-cluster --region "$AWS_REGION" > /dev/null 2>&1
    AWS_LB=$(kubectl get ingress gopro-ingress -n gopro -o jsonpath='{.status.loadBalancer.ingress[0].hostname}' 2>/dev/null || echo "Not available")
    echo "  Load Balancer: $AWS_LB"
    echo "  Region: $AWS_REGION"
    echo ""
fi

if [ "$DEPLOY_GCP" = "true" ]; then
    echo "GCP Deployment:"
    gcloud container clusters get-credentials gopro-prod-cluster --region "$GCP_REGION" > /dev/null 2>&1
    GCP_LB=$(kubectl get svc ingress-nginx-controller -n ingress-nginx -o jsonpath='{.status.loadBalancer.ingress[0].ip}' 2>/dev/null || echo "Not available")
    echo "  Load Balancer: $GCP_LB"
    echo "  Region: $GCP_REGION"
    echo ""
fi

echo "Next Steps:"
echo "1. Configure DNS routing (Route 53 + Cloud DNS)"
echo "2. Setup database replication (RDS → Cloud SQL)"
echo "3. Configure storage sync (S3 ↔ Cloud Storage)"
echo "4. Setup unified monitoring (Prometheus federation)"
echo "5. Test failover scenarios"
echo "6. Configure alerting"
echo ""

echo "Useful Commands:"
echo "  # Switch to AWS cluster"
echo "  aws eks update-kubeconfig --name gopro-prod-cluster --region $AWS_REGION"
echo ""
echo "  # Switch to GCP cluster"
echo "  gcloud container clusters get-credentials gopro-prod-cluster --region $GCP_REGION"
echo ""
echo "  # Test failover"
echo "  ./test-failover.sh"
echo ""

