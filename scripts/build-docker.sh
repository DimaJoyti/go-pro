#!/bin/bash

# GO-PRO Docker Build Script
# This script builds Docker images for different environments with proper tagging and metadata

set -euo pipefail

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
BACKEND_DIR="$PROJECT_ROOT/backend"

# Default values
REGISTRY="${REGISTRY:-ghcr.io/your-org}"
IMAGE_NAME="${IMAGE_NAME:-go-pro}"
VERSION="${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo 'dev')}"
BUILD_DATE="${BUILD_DATE:-$(date -u +'%Y-%m-%dT%H:%M:%SZ')}"
VCS_REF="${VCS_REF:-$(git rev-parse HEAD 2>/dev/null || echo 'unknown')}"
ENVIRONMENT="${ENVIRONMENT:-development}"
PUSH="${PUSH:-false}"
PLATFORM="${PLATFORM:-linux/amd64,linux/arm64}"

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
GO-PRO Docker Build Script

Usage: $0 [OPTIONS]

OPTIONS:
    -e, --environment ENV    Build environment (development|production|test) [default: development]
    -v, --version VERSION    Image version tag [default: git describe or 'dev']
    -r, --registry REGISTRY  Docker registry [default: ghcr.io/your-org]
    -n, --name NAME          Image name [default: go-pro]
    -p, --push               Push images to registry after build
    --platform PLATFORMS     Target platforms [default: linux/amd64,linux/arm64]
    --no-cache               Build without using cache
    --help                   Show this help message

EXAMPLES:
    # Build development image
    $0 -e development

    # Build and push production image
    $0 -e production -v v1.0.0 -p

    # Build for specific platform
    $0 --platform linux/amd64

    # Build with custom registry
    $0 -r my-registry.com/myorg -n my-app

ENVIRONMENT VARIABLES:
    REGISTRY                 Docker registry (same as -r)
    IMAGE_NAME              Image name (same as -n)
    VERSION                 Version tag (same as -v)
    ENVIRONMENT             Build environment (same as -e)
    PUSH                    Push after build (same as -p)
    PLATFORM                Target platforms
    DOCKER_BUILDKIT         Enable BuildKit (recommended: 1)

EOF
}

# Parse command line arguments
parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -e|--environment)
                ENVIRONMENT="$2"
                shift 2
                ;;
            -v|--version)
                VERSION="$2"
                shift 2
                ;;
            -r|--registry)
                REGISTRY="$2"
                shift 2
                ;;
            -n|--name)
                IMAGE_NAME="$2"
                shift 2
                ;;
            -p|--push)
                PUSH="true"
                shift
                ;;
            --platform)
                PLATFORM="$2"
                shift 2
                ;;
            --no-cache)
                NO_CACHE="--no-cache"
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

# Validate environment
validate_environment() {
    case $ENVIRONMENT in
        development|production|test)
            ;;
        *)
            log_error "Invalid environment: $ENVIRONMENT. Must be one of: development, production, test"
            exit 1
            ;;
    esac
}

# Check prerequisites
check_prerequisites() {
    log_info "Checking prerequisites..."

    # Check if Docker is installed and running
    if ! command -v docker &> /dev/null; then
        log_error "Docker is not installed or not in PATH"
        exit 1
    fi

    if ! docker info &> /dev/null; then
        log_error "Docker daemon is not running"
        exit 1
    fi

    # Check if buildx is available for multi-platform builds
    if [[ "$PLATFORM" == *","* ]]; then
        if ! docker buildx version &> /dev/null; then
            log_error "Docker buildx is required for multi-platform builds"
            exit 1
        fi
    fi

    # Check if we're in the right directory
    if [[ ! -f "$BACKEND_DIR/go.mod" ]]; then
        log_error "Backend go.mod not found. Are you in the project root?"
        exit 1
    fi

    log_success "Prerequisites check passed"
}

# Get Dockerfile based on environment
get_dockerfile() {
    case $ENVIRONMENT in
        development)
            echo "Dockerfile.dev"
            ;;
        production)
            echo "Dockerfile.prod"
            ;;
        test)
            echo "Dockerfile.dev"
            ;;
        *)
            echo "Dockerfile"
            ;;
    esac
}

# Build Docker image
build_image() {
    local dockerfile=$(get_dockerfile)
    local full_image_name="$REGISTRY/$IMAGE_NAME"
    local build_args=(
        "--build-arg" "VERSION=$VERSION"
        "--build-arg" "BUILD_DATE=$BUILD_DATE"
        "--build-arg" "VCS_REF=$VCS_REF"
    )

    log_info "Building Docker image..."
    log_info "Environment: $ENVIRONMENT"
    log_info "Dockerfile: $dockerfile"
    log_info "Image: $full_image_name:$VERSION"
    log_info "Platform: $PLATFORM"

    # Prepare build command
    local build_cmd="docker"
    local build_args_str=""

    # Use buildx for multi-platform builds
    if [[ "$PLATFORM" == *","* ]]; then
        build_cmd="docker buildx"
        build_args_str="--platform $PLATFORM"
        
        # Create builder if it doesn't exist
        if ! docker buildx inspect go-pro-builder &> /dev/null; then
            log_info "Creating buildx builder..."
            docker buildx create --name go-pro-builder --use
        fi
    fi

    # Add cache options
    if [[ -z "${NO_CACHE:-}" ]]; then
        build_args_str="$build_args_str --cache-from type=local,src=/tmp/.buildx-cache"
        build_args_str="$build_args_str --cache-to type=local,dest=/tmp/.buildx-cache-new,mode=max"
    else
        build_args_str="$build_args_str $NO_CACHE"
    fi

    # Build the image
    $build_cmd build \
        $build_args_str \
        "${build_args[@]}" \
        -f "$BACKEND_DIR/$dockerfile" \
        -t "$full_image_name:$VERSION" \
        -t "$full_image_name:latest" \
        "$BACKEND_DIR"

    # Move cache
    if [[ -z "${NO_CACHE:-}" ]] && [[ -d "/tmp/.buildx-cache-new" ]]; then
        rm -rf /tmp/.buildx-cache
        mv /tmp/.buildx-cache-new /tmp/.buildx-cache
    fi

    log_success "Image built successfully: $full_image_name:$VERSION"
}

# Push Docker image
push_image() {
    if [[ "$PUSH" == "true" ]]; then
        local full_image_name="$REGISTRY/$IMAGE_NAME"
        
        log_info "Pushing Docker image to registry..."
        
        docker push "$full_image_name:$VERSION"
        docker push "$full_image_name:latest"
        
        log_success "Image pushed successfully"
    else
        log_info "Skipping push (use -p or --push to enable)"
    fi
}

# Clean up old images
cleanup() {
    log_info "Cleaning up old images..."
    
    # Remove dangling images
    docker image prune -f
    
    log_success "Cleanup completed"
}

# Main function
main() {
    log_info "Starting GO-PRO Docker build process..."
    
    parse_args "$@"
    validate_environment
    check_prerequisites
    
    # Enable BuildKit for better performance
    export DOCKER_BUILDKIT=1
    
    build_image
    push_image
    cleanup
    
    log_success "Build process completed successfully!"
    log_info "Image: $REGISTRY/$IMAGE_NAME:$VERSION"
}

# Run main function with all arguments
main "$@"
