#!/bin/bash

# GO-PRO Kubernetes Deployment Script
# This script deploys the complete GO-PRO application stack to Kubernetes

set -euo pipefail

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
K8S_DIR="$PROJECT_ROOT/deploy/k8s"

# Default values
NAMESPACE="${NAMESPACE:-go-pro}"
ENVIRONMENT="${ENVIRONMENT:-production}"
KUBECTL_CONTEXT="${KUBECTL_CONTEXT:-}"
DRY_RUN="${DRY_RUN:-false}"
SKIP_BUILD="${SKIP_BUILD:-false}"
ENABLE_ISTIO="${ENABLE_ISTIO:-true}"
ENABLE_MONITORING="${ENABLE_MONITORING:-true}"
WAIT_FOR_ROLLOUT="${WAIT_FOR_ROLLOUT:-true}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Help function
show_help() {
    cat << EOF
GO-PRO Kubernetes Deployment Script

Usage: $0 [OPTIONS]

OPTIONS:
    -n, --namespace NAMESPACE    Kubernetes namespace [default: go-pro]
    -e, --environment ENV        Environment (development|staging|production) [default: production]
    -c, --context CONTEXT        Kubectl context to use
    --dry-run                    Show what would be deployed without applying
    --skip-build                 Skip building Docker images
    --no-istio                   Disable Istio service mesh
    --no-monitoring              Disable monitoring stack
    --no-wait                    Don't wait for rollout completion
    --help                       Show this help message

EXAMPLES:
    # Deploy to production
    $0 -e production

    # Deploy to staging with custom namespace
    $0 -e staging -n go-pro-staging

    # Dry run deployment
    $0 --dry-run

    # Deploy without Istio
    $0 --no-istio

ENVIRONMENT VARIABLES:
    NAMESPACE                    Kubernetes namespace (same as -n)
    ENVIRONMENT                  Environment (same as -e)
    KUBECTL_CONTEXT              Kubectl context (same as -c)
    DRY_RUN                      Dry run mode (same as --dry-run)
    SKIP_BUILD                   Skip build (same as --skip-build)
    ENABLE_ISTIO                 Enable Istio (opposite of --no-istio)
    ENABLE_MONITORING            Enable monitoring (opposite of --no-monitoring)

EOF
}

# Parse command line arguments
parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -n|--namespace)
                NAMESPACE="$2"
                shift 2
                ;;
            -e|--environment)
                ENVIRONMENT="$2"
                shift 2
                ;;
            -c|--context)
                KUBECTL_CONTEXT="$2"
                shift 2
                ;;
            --dry-run)
                DRY_RUN="true"
                shift
                ;;
            --skip-build)
                SKIP_BUILD="true"
                shift
                ;;
            --no-istio)
                ENABLE_ISTIO="false"
                shift
                ;;
            --no-monitoring)
                ENABLE_MONITORING="false"
                shift
                ;;
            --no-wait)
                WAIT_FOR_ROLLOUT="false"
                shift
                ;;
            --help)
                show_help
                exit 0
                ;;
            *)
                log_error "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
    done
}

# Check prerequisites
check_prerequisites() {
    log_info "Checking prerequisites..."

    # Check if kubectl is installed
    if ! command -v kubectl &> /dev/null; then
        log_error "kubectl is not installed or not in PATH"
        exit 1
    fi

    # Check if Docker is installed (if not skipping build)
    if [[ "$SKIP_BUILD" == "false" ]] && ! command -v docker &> /dev/null; then
        log_error "Docker is not installed or not in PATH"
        exit 1
    fi

    # Set kubectl context if provided
    if [[ -n "$KUBECTL_CONTEXT" ]]; then
        kubectl config use-context "$KUBECTL_CONTEXT"
    fi

    # Check if cluster is accessible
    if ! kubectl cluster-info &> /dev/null; then
        log_error "Cannot connect to Kubernetes cluster"
        exit 1
    fi

    # Check if Istio is installed (if enabled)
    if [[ "$ENABLE_ISTIO" == "true" ]]; then
        if ! kubectl get namespace istio-system &> /dev/null; then
            log_warning "Istio system namespace not found. Istio may not be installed."
            read -p "Continue without Istio? (y/N): " -n 1 -r
            echo
            if [[ ! $REPLY =~ ^[Yy]$ ]]; then
                exit 1
            fi
            ENABLE_ISTIO="false"
        fi
    fi

    log_success "Prerequisites check passed"
}

# Build Docker images
build_images() {
    if [[ "$SKIP_BUILD" == "true" ]]; then
        log_info "Skipping Docker image build"
        return
    fi

    log_info "Building Docker images..."

    # Build backend image
    "$PROJECT_ROOT/scripts/build-docker.sh" \
        -e "$ENVIRONMENT" \
        -v "$(git describe --tags --always --dirty 2>/dev/null || echo 'latest')" \
        --platform linux/amd64

    log_success "Docker images built successfully"
}

# Create namespace
create_namespace() {
    log_info "Creating namespace: $NAMESPACE"

    if [[ "$DRY_RUN" == "true" ]]; then
        kubectl create namespace "$NAMESPACE" --dry-run=client -o yaml
    else
        kubectl create namespace "$NAMESPACE" --dry-run=client -o yaml | kubectl apply -f -
        
        # Label namespace for Istio injection if enabled
        if [[ "$ENABLE_ISTIO" == "true" ]]; then
            kubectl label namespace "$NAMESPACE" istio-injection=enabled --overwrite
        fi
    fi
}

# Deploy base resources
deploy_base() {
    log_info "Deploying base resources..."

    local dry_run_flag=""
    if [[ "$DRY_RUN" == "true" ]]; then
        dry_run_flag="--dry-run=client"
    fi

    # Apply in order of dependencies
    kubectl apply $dry_run_flag -f "$K8S_DIR/namespace.yaml"
    kubectl apply $dry_run_flag -f "$K8S_DIR/configmap.yaml"
    kubectl apply $dry_run_flag -f "$K8S_DIR/secrets.yaml"
}

# Deploy storage
deploy_storage() {
    log_info "Deploying storage components..."

    local dry_run_flag=""
    if [[ "$DRY_RUN" == "true" ]]; then
        dry_run_flag="--dry-run=client"
    fi

    # Deploy PostgreSQL
    if [[ -f "$K8S_DIR/postgres.yaml" ]]; then
        kubectl apply $dry_run_flag -f "$K8S_DIR/postgres.yaml"
    fi

    # Deploy Redis
    if [[ -f "$K8S_DIR/redis.yaml" ]]; then
        kubectl apply $dry_run_flag -f "$K8S_DIR/redis.yaml"
    fi

    # Deploy Kafka
    if [[ -f "$K8S_DIR/kafka.yaml" ]]; then
        kubectl apply $dry_run_flag -f "$K8S_DIR/kafka.yaml"
    fi
}

# Deploy application
deploy_application() {
    log_info "Deploying application components..."

    local dry_run_flag=""
    if [[ "$DRY_RUN" == "true" ]]; then
        dry_run_flag="--dry-run=client"
    fi

    # Deploy backend
    kubectl apply $dry_run_flag -f "$K8S_DIR/deployment.yaml"
    kubectl apply $dry_run_flag -f "$K8S_DIR/service.yaml"

    # Deploy frontend (if exists)
    if [[ -f "$K8S_DIR/frontend.yaml" ]]; then
        kubectl apply $dry_run_flag -f "$K8S_DIR/frontend.yaml"
    fi
}

# Deploy networking
deploy_networking() {
    log_info "Deploying networking components..."

    local dry_run_flag=""
    if [[ "$DRY_RUN" == "true" ]]; then
        dry_run_flag="--dry-run=client"
    fi

    # Deploy network policies
    kubectl apply $dry_run_flag -f "$K8S_DIR/network-policy.yaml"

    # Deploy ingress
    kubectl apply $dry_run_flag -f "$K8S_DIR/ingress.yaml"

    # Deploy Istio resources if enabled
    if [[ "$ENABLE_ISTIO" == "true" ]]; then
        kubectl apply $dry_run_flag -f "$K8S_DIR/istio.yaml"
    fi
}

# Deploy autoscaling
deploy_autoscaling() {
    log_info "Deploying autoscaling components..."

    local dry_run_flag=""
    if [[ "$DRY_RUN" == "true" ]]; then
        dry_run_flag="--dry-run=client"
    fi

    kubectl apply $dry_run_flag -f "$K8S_DIR/hpa.yaml"
}

# Deploy monitoring
deploy_monitoring() {
    if [[ "$ENABLE_MONITORING" == "false" ]]; then
        log_info "Skipping monitoring deployment"
        return
    fi

    log_info "Deploying monitoring stack..."

    local dry_run_flag=""
    if [[ "$DRY_RUN" == "true" ]]; then
        dry_run_flag="--dry-run=client"
    fi

    kubectl apply $dry_run_flag -f "$K8S_DIR/monitoring.yaml"
}

# Wait for rollout
wait_for_rollout() {
    if [[ "$WAIT_FOR_ROLLOUT" == "false" ]] || [[ "$DRY_RUN" == "true" ]]; then
        return
    fi

    log_info "Waiting for rollout to complete..."

    # Wait for deployments
    kubectl rollout status deployment/go-pro-backend -n "$NAMESPACE" --timeout=600s
    
    if [[ -f "$K8S_DIR/frontend.yaml" ]]; then
        kubectl rollout status deployment/go-pro-frontend -n "$NAMESPACE" --timeout=600s
    fi

    # Wait for StatefulSets
    if kubectl get statefulset postgres -n "$NAMESPACE" &> /dev/null; then
        kubectl rollout status statefulset/postgres -n "$NAMESPACE" --timeout=600s
    fi

    if kubectl get statefulset redis -n "$NAMESPACE" &> /dev/null; then
        kubectl rollout status statefulset/redis -n "$NAMESPACE" --timeout=600s
    fi

    log_success "Rollout completed successfully"
}

# Verify deployment
verify_deployment() {
    if [[ "$DRY_RUN" == "true" ]]; then
        return
    fi

    log_info "Verifying deployment..."

    # Check pod status
    kubectl get pods -n "$NAMESPACE"

    # Check service endpoints
    kubectl get endpoints -n "$NAMESPACE"

    # Run health checks
    if kubectl get service go-pro-backend-service -n "$NAMESPACE" &> /dev/null; then
        log_info "Running health check..."
        kubectl run health-check --rm -i --restart=Never --image=curlimages/curl -- \
            curl -f "http://go-pro-backend-service.$NAMESPACE.svc.cluster.local:8080/api/v1/health" || \
            log_warning "Health check failed"
    fi

    log_success "Deployment verification completed"
}

# Cleanup function
cleanup() {
    if [[ "$DRY_RUN" == "true" ]]; then
        return
    fi

    log_info "Cleaning up temporary resources..."
    kubectl delete pod health-check -n "$NAMESPACE" --ignore-not-found=true
}

# Main function
main() {
    log_info "Starting GO-PRO Kubernetes deployment..."
    
    parse_args "$@"
    check_prerequisites
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log_info "Running in dry-run mode - no changes will be applied"
    fi
    
    build_images
    create_namespace
    deploy_base
    deploy_storage
    deploy_application
    deploy_networking
    deploy_autoscaling
    deploy_monitoring
    wait_for_rollout
    verify_deployment
    cleanup
    
    log_success "Deployment completed successfully!"
    
    if [[ "$DRY_RUN" == "false" ]]; then
        log_info "Application endpoints:"
        log_info "- Backend API: http://$(kubectl get service go-pro-backend-service -n "$NAMESPACE" -o jsonpath='{.status.loadBalancer.ingress[0].ip}'):8080"
        log_info "- Frontend: http://$(kubectl get service go-pro-frontend-service -n "$NAMESPACE" -o jsonpath='{.status.loadBalancer.ingress[0].ip}'):3000"
        
        if [[ "$ENABLE_MONITORING" == "true" ]]; then
            log_info "- Grafana: http://$(kubectl get service grafana-service -n "$NAMESPACE" -o jsonpath='{.status.loadBalancer.ingress[0].ip}'):3000"
            log_info "- Prometheus: http://$(kubectl get service prometheus-service -n "$NAMESPACE" -o jsonpath='{.status.loadBalancer.ingress[0].ip}'):9090"
            log_info "- Jaeger: http://$(kubectl get service jaeger-service -n "$NAMESPACE" -o jsonpath='{.status.loadBalancer.ingress[0].ip}'):16686"
        fi
    fi
}

# Trap cleanup on exit
trap cleanup EXIT

# Run main function with all arguments
main "$@"
