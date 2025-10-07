#!/bin/bash

# GO-PRO Learning Platform Backend - Development Setup Script
# This script sets up the complete development environment

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# Functions
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

log_step() {
    echo -e "\n${PURPLE}===${NC} $1 ${PURPLE}===${NC}"
}

check_command() {
    if command -v "$1" &> /dev/null; then
        log_success "$1 is installed"
        return 0
    else
        log_error "$1 is not installed"
        return 1
    fi
}

install_go_tools() {
    log_step "Installing Go development tools"

    tools=(
        "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
        "github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"
        "github.com/cosmtrek/air@latest"
        "golang.org/x/tools/cmd/goimports@latest"
        "golang.org/x/vuln/cmd/govulncheck@latest"
    )

    for tool in "${tools[@]}"; do
        log_info "Installing $tool"
        go install "$tool" || log_warning "Failed to install $tool"
    done

    log_success "Go tools installation completed"
}

install_system_tools() {
    log_step "Installing system tools"

    # Detect OS
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        if command -v apt-get &> /dev/null; then
            # Debian/Ubuntu
            sudo apt-get update
            sudo apt-get install -y curl wget git make python3 python3-pip
        elif command -v yum &> /dev/null; then
            # CentOS/RHEL
            sudo yum update -y
            sudo yum install -y curl wget git make python3 python3-pip
        elif command -v pacman &> /dev/null; then
            # Arch Linux
            sudo pacman -Sy --noconfirm curl wget git make python python-pip
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        if command -v brew &> /dev/null; then
            brew update
            brew install curl wget git make python3
        else
            log_error "Homebrew not found. Please install Homebrew first."
            return 1
        fi
    fi

    # Install Python packages
    if command -v pip3 &> /dev/null; then
        pip3 install --user pre-commit
    elif command -v pip &> /dev/null; then
        pip install --user pre-commit
    else
        log_warning "pip not found, skipping pre-commit installation"
    fi

    log_success "System tools installation completed"
}

setup_git_hooks() {
    log_step "Setting up Git hooks"

    cd "$PROJECT_ROOT"

    if command -v pre-commit &> /dev/null; then
        pre-commit install
        pre-commit install --hook-type commit-msg
        log_success "Pre-commit hooks installed"
    else
        log_warning "pre-commit not found, skipping hook setup"
    fi
}

setup_environment() {
    log_step "Setting up environment"

    cd "$PROJECT_ROOT"

    # Create necessary directories
    mkdir -p backend/bin
    mkdir -p backend/tmp
    mkdir -p backend/logs
    mkdir -p docs
    mkdir -p deploy/k8s
    mkdir -p config

    # Create .env template if it doesn't exist
    if [[ ! -f backend/.env.example ]]; then
        cat > backend/.env.example << EOF
# GO-PRO Backend Environment Variables
GO_ENV=development
PORT=8080
LOG_LEVEL=debug

# Database
DATABASE_URL=postgres://gopro_user:gopro_password@localhost:5432/gopro_dev?sslmode=disable

# Redis
REDIS_URL=redis://localhost:6379

# Security
JWT_SECRET=your-jwt-secret-here

# Optional
READ_TIMEOUT=15s
WRITE_TIMEOUT=15s
IDLE_TIMEOUT=60s
EOF
        log_success "Created .env.example template"
    fi

    log_success "Environment setup completed"
}

verify_prerequisites() {
    log_step "Verifying prerequisites"

    local missing=0

    # Check required tools
    check_command "go" || missing=$((missing + 1))
    check_command "docker" || missing=$((missing + 1))
    check_command "make" || missing=$((missing + 1))

    # Check Go version
    if command -v go &> /dev/null; then
        go_version=$(go version | grep -o 'go[0-9]\+\.[0-9]\+' | sed 's/go//')
        required_version="1.21"
        if [[ "$(printf '%s\n' "$required_version" "$go_version" | sort -V | head -n1)" != "$required_version" ]]; then
            log_error "Go version $go_version is too old. Required: $required_version or higher"
            missing=$((missing + 1))
        else
            log_success "Go version $go_version is compatible"
        fi
    fi

    # Check Docker Compose
    if command -v docker &> /dev/null; then
        if docker compose version &> /dev/null; then
            log_success "Docker Compose is available"
        elif command -v docker-compose &> /dev/null; then
            log_success "Docker Compose (standalone) is available"
        else
            log_warning "Docker Compose not found, some features may not work"
        fi
    fi

    if [[ $missing -gt 0 ]]; then
        log_error "$missing required prerequisites are missing"
        return 1
    fi

    log_success "All prerequisites verified"
    return 0
}

test_installation() {
    log_step "Testing installation"

    cd "$PROJECT_ROOT/backend"

    # Test Go build
    log_info "Testing Go build"
    if go build -o bin/test-build .; then
        log_success "Go build test passed"
        rm -f bin/test-build
    else
        log_error "Go build test failed"
        return 1
    fi

    # Test linter
    log_info "Testing linter"
    if golangci-lint --version &> /dev/null; then
        log_success "Linter test passed"
    else
        log_warning "Linter test failed"
    fi

    # Test security scanner
    log_info "Testing security scanner"
    if gosec -version &> /dev/null; then
        log_success "Security scanner test passed"
    else
        log_warning "Security scanner test failed"
    fi

    log_success "Installation test completed"
}

show_next_steps() {
    log_step "Setup Complete!"

    echo -e "${GREEN}✅ Development environment setup completed successfully!${NC}"
    echo
    echo -e "${CYAN}Next steps:${NC}"
    echo "1. Start the development environment:"
    echo "   ${YELLOW}make dev${NC}              # Start with hot reload"
    echo "   ${YELLOW}make docker-dev${NC}       # Start full Docker environment"
    echo
    echo "2. Run quality checks:"
    echo "   ${YELLOW}make lint${NC}             # Run linter"
    echo "   ${YELLOW}make test${NC}             # Run tests"
    echo "   ${YELLOW}make security${NC}         # Run security scan"
    echo
    echo "3. Build and deploy:"
    echo "   ${YELLOW}make build${NC}            # Build the application"
    echo "   ${YELLOW}make docker-build${NC}     # Build Docker image"
    echo
    echo "4. Access services (with docker-dev):"
    echo "   - Backend API: ${BLUE}http://localhost:8080${NC}"
    echo "   - Database Admin: ${BLUE}http://localhost:8081${NC}"
    echo "   - Redis Commander: ${BLUE}http://localhost:8082${NC}"
    echo "   - Prometheus: ${BLUE}http://localhost:9090${NC}"
    echo "   - Grafana: ${BLUE}http://localhost:3000${NC}"
    echo
    echo "5. Documentation:"
    echo "   - Development: ${BLUE}docs/DEVELOPMENT.md${NC}"
    echo "   - Deployment: ${BLUE}docs/DEPLOYMENT.md${NC}"
    echo
    echo -e "${GREEN}Happy coding! 🚀${NC}"
}

main() {
    echo -e "${CYAN}"
    echo "╔══════════════════════════════════════════════╗"
    echo "║          GO-PRO Development Setup            ║"
    echo "║     Learning Platform Backend Environment    ║"
    echo "╚══════════════════════════════════════════════╝"
    echo -e "${NC}"

    # Check if we're in the right directory
    if [[ ! -f "$PROJECT_ROOT/go.mod" ]] && [[ ! -f "$PROJECT_ROOT/backend/go.mod" ]]; then
        log_error "This doesn't appear to be the GO-PRO project directory"
        log_info "Please run this script from the project root directory"
        exit 1
    fi

    # Parse command line arguments
    INSTALL_SYSTEM_TOOLS=false
    SKIP_VERIFICATION=false

    while [[ $# -gt 0 ]]; do
        case $1 in
            --install-system-tools)
                INSTALL_SYSTEM_TOOLS=true
                shift
                ;;
            --skip-verification)
                SKIP_VERIFICATION=true
                shift
                ;;
            -h|--help)
                echo "Usage: $0 [OPTIONS]"
                echo
                echo "Options:"
                echo "  --install-system-tools    Install system packages (requires sudo)"
                echo "  --skip-verification       Skip prerequisite verification"
                echo "  -h, --help               Show this help message"
                echo
                exit 0
                ;;
            *)
                log_error "Unknown option: $1"
                exit 1
                ;;
        esac
    done

    # Main setup flow
    if [[ "$SKIP_VERIFICATION" != "true" ]]; then
        verify_prerequisites || {
            log_error "Prerequisites check failed"
            echo "You can skip this check with --skip-verification"
            exit 1
        }
    fi

    if [[ "$INSTALL_SYSTEM_TOOLS" == "true" ]]; then
        install_system_tools
    fi

    setup_environment
    install_go_tools
    setup_git_hooks
    test_installation
    show_next_steps

    log_success "Setup completed successfully!"
}

# Check if script is being sourced or executed
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi