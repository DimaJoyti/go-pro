#!/bin/bash

# GO-PRO Application Deployment Script
# This script deploys the GO-PRO application to EKS

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
ENVIRONMENT="${ENVIRONMENT:-development}"
NAMESPACE="gopro"
CLUSTER_NAME="${CLUSTER_NAME:-gopro-dev-cluster}"
REGION="${AWS_REGION:-us-east-1}"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}GO-PRO Application Deployment${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Environment: $ENVIRONMENT"
echo "Namespace: $NAMESPACE"
echo "Cluster: $CLUSTER_NAME"
echo "Region: $REGION"
echo ""

# Check prerequisites
echo -e "${YELLOW}Checking prerequisites...${NC}"
command -v kubectl >/dev/null 2>&1 || { echo -e "${RED}kubectl is required but not installed.${NC}" >&2; exit 1; }
command -v aws >/dev/null 2>&1 || { echo -e "${RED}AWS CLI is required but not installed.${NC}" >&2; exit 1; }
echo -e "${GREEN}✓ All prerequisites met${NC}"

# Update kubeconfig
echo -e "${YELLOW}Updating kubeconfig...${NC}"
aws eks update-kubeconfig --name "$CLUSTER_NAME" --region "$REGION"
echo -e "${GREEN}✓ kubeconfig updated${NC}"

# Verify cluster access
echo -e "${YELLOW}Verifying cluster access...${NC}"
kubectl cluster-info > /dev/null 2>&1 || { echo -e "${RED}Cannot access cluster${NC}" >&2; exit 1; }
echo -e "${GREEN}✓ Cluster access verified${NC}"

# Create namespace if it doesn't exist
echo -e "${YELLOW}Creating namespace...${NC}"
kubectl create namespace "$NAMESPACE" --dry-run=client -o yaml | kubectl apply -f -
echo -e "${GREEN}✓ Namespace ready${NC}"

# Check for required secrets
echo -e "${YELLOW}Checking for required secrets...${NC}"
REQUIRED_SECRETS=("database-credentials" "redis-credentials" "jwt-secret")
MISSING_SECRETS=()

for secret in "${REQUIRED_SECRETS[@]}"; do
    if ! kubectl get secret "$secret" -n "$NAMESPACE" > /dev/null 2>&1; then
        MISSING_SECRETS+=("$secret")
    fi
done

if [ ${#MISSING_SECRETS[@]} -ne 0 ]; then
    echo -e "${RED}Missing required secrets:${NC}"
    for secret in "${MISSING_SECRETS[@]}"; do
        echo "  - $secret"
    done
    echo ""
    echo "Please create the missing secrets before deploying."
    echo "Example:"
    echo "  kubectl create secret generic database-credentials \\"
    echo "    --from-literal=host=<db-host> \\"
    echo "    --from-literal=port=5432 \\"
    echo "    --from-literal=username=<db-user> \\"
    echo "    --from-literal=password=<db-password> \\"
    echo "    --from-literal=database=gopro \\"
    echo "    -n $NAMESPACE"
    exit 1
fi
echo -e "${GREEN}✓ All required secrets present${NC}"

# Build and push Docker images (if needed)
if [ "$BUILD_IMAGES" = "true" ]; then
    echo -e "${YELLOW}Building and pushing Docker images...${NC}"
    
    # Backend
    echo -e "${BLUE}Building backend image...${NC}"
    docker build -t ghcr.io/dimajoyti/go-pro/backend:$ENVIRONMENT ../../backend
    docker push ghcr.io/dimajoyti/go-pro/backend:$ENVIRONMENT
    
    # Frontend
    echo -e "${BLUE}Building frontend image...${NC}"
    docker build -t ghcr.io/dimajoyti/go-pro/frontend:$ENVIRONMENT ../../frontend
    docker push ghcr.io/dimajoyti/go-pro/frontend:$ENVIRONMENT
    
    echo -e "${GREEN}✓ Images built and pushed${NC}"
fi

# Deploy application using Kustomize
echo -e "${YELLOW}Deploying application...${NC}"
kubectl apply -k "../../k8s/overlays/$ENVIRONMENT"
echo -e "${GREEN}✓ Application deployed${NC}"

# Wait for deployments to be ready
echo -e "${YELLOW}Waiting for deployments to be ready...${NC}"
kubectl wait --for=condition=available --timeout=300s \
    deployment/backend deployment/frontend -n "$NAMESPACE" || {
    echo -e "${RED}Deployment failed to become ready${NC}"
    echo "Checking pod status:"
    kubectl get pods -n "$NAMESPACE"
    echo ""
    echo "Recent events:"
    kubectl get events -n "$NAMESPACE" --sort-by='.lastTimestamp' | tail -20
    exit 1
}
echo -e "${GREEN}✓ Deployments ready${NC}"

# Get deployment status
echo -e "${YELLOW}Deployment status:${NC}"
kubectl get deployments -n "$NAMESPACE"
echo ""
kubectl get pods -n "$NAMESPACE"
echo ""
kubectl get svc -n "$NAMESPACE"
echo ""

# Get ingress information
echo -e "${YELLOW}Ingress information:${NC}"
kubectl get ingress -n "$NAMESPACE"
echo ""

# Get load balancer URL
echo -e "${YELLOW}Getting load balancer URL...${NC}"
LB_URL=$(kubectl get ingress gopro-ingress -n "$NAMESPACE" -o jsonpath='{.status.loadBalancer.ingress[0].hostname}' 2>/dev/null || echo "Not available yet")
echo "Load Balancer URL: $LB_URL"
echo ""

# Run smoke tests
if [ "$RUN_SMOKE_TESTS" = "true" ]; then
    echo -e "${YELLOW}Running smoke tests...${NC}"
    
    # Wait for load balancer to be ready
    echo "Waiting for load balancer to be ready..."
    sleep 30
    
    # Test backend health endpoint
    echo "Testing backend health endpoint..."
    if curl -f -s "http://$LB_URL/health" > /dev/null; then
        echo -e "${GREEN}✓ Backend health check passed${NC}"
    else
        echo -e "${RED}✗ Backend health check failed${NC}"
    fi
    
    # Test frontend
    echo "Testing frontend..."
    if curl -f -s "http://$LB_URL/" > /dev/null; then
        echo -e "${GREEN}✓ Frontend check passed${NC}"
    else
        echo -e "${RED}✗ Frontend check failed${NC}"
    fi
fi

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Deployment complete!${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Application URLs:"
echo "  Frontend: http://$LB_URL"
echo "  Backend API: http://$LB_URL/api"
echo ""
echo "Useful commands:"
echo "  kubectl get pods -n $NAMESPACE"
echo "  kubectl logs -f deployment/backend -n $NAMESPACE"
echo "  kubectl logs -f deployment/frontend -n $NAMESPACE"
echo "  kubectl describe pod <pod-name> -n $NAMESPACE"
echo "  kubectl exec -it deployment/backend -n $NAMESPACE -- /bin/sh"
echo ""
echo "Monitoring:"
echo "  kubectl port-forward svc/prometheus-kube-prometheus-prometheus 9090:9090 -n monitoring"
echo "  kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring"
echo ""

