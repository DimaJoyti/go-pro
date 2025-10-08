.PHONY: help build run test clean docker-build docker-run docker-stop

# Variables
APP_NAME=url-shortener
DOCKER_IMAGE=url-shortener:latest
PORT=8080

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) ./cmd/main.go
	@echo "✓ Build complete: bin/$(APP_NAME)"

run: ## Run the application
	@echo "Starting $(APP_NAME)..."
	@go run ./cmd/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -cover -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report: coverage.html"

bench: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./...

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@goimports -w .

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "✓ Clean complete"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "✓ Dependencies updated"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -f docker/Dockerfile -t $(DOCKER_IMAGE) .
	@echo "✓ Docker image built: $(DOCKER_IMAGE)"

docker-run: ## Run Docker container
	@echo "Starting Docker container..."
	@docker-compose -f docker/docker-compose.yml up -d
	@echo "✓ Container started on http://localhost:$(PORT)"

docker-stop: ## Stop Docker container
	@echo "Stopping Docker container..."
	@docker-compose -f docker/docker-compose.yml down
	@echo "✓ Container stopped"

docker-logs: ## View Docker logs
	@docker-compose -f docker/docker-compose.yml logs -f

install: ## Install the application
	@echo "Installing $(APP_NAME)..."
	@go install ./cmd/main.go
	@echo "✓ Installed"

dev: ## Run in development mode with auto-reload
	@echo "Starting development server..."
	@air || go run ./cmd/main.go

.DEFAULT_GOAL := help

