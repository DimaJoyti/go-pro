.PHONY: help build up down logs test clean

# Colors for output
RED = \033[0;31m
GREEN = \033[0;32m
YELLOW = \033[1;33m
NC = \033[0m # No Color

help: ## Display this help screen
	@echo "$(GREEN)GO-PRO Microservices Management$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(YELLOW)%-30s$(NC) %s\n", $$1, $$2}'

# Docker Compose Commands
up: ## Start all services
	@echo "$(GREEN)Starting all microservices...$(NC)"
	docker-compose up -d
	@echo "$(GREEN)Services started successfully!$(NC)"
	@echo "API Gateway: http://localhost:8080"
	@echo "User Service: http://localhost:8081"
	@echo "Course Service: http://localhost:8082"
	@echo "Progress Service: http://localhost:8083"
	@echo "Jaeger UI: http://localhost:16686"
	@echo "Prometheus: http://localhost:9090"
	@echo "Grafana: http://localhost:3000"

down: ## Stop all services
	@echo "$(YELLOW)Stopping all microservices...$(NC)"
	docker-compose down

restart: ## Restart all services
	@echo "$(YELLOW)Restarting all microservices...$(NC)"
	docker-compose restart

logs: ## View logs from all services
	docker-compose logs -f

logs-api-gateway: ## View API Gateway logs
	docker-compose logs -f api-gateway

logs-user: ## View User Service logs
	docker-compose logs -f user-service

logs-course: ## View Course Service logs
	docker-compose logs -f course-service

logs-progress: ## View Progress Service logs
	docker-compose logs -f progress-service

# Build Commands
build: ## Build all services
	@echo "$(GREEN)Building all microservices...$(NC)"
	docker-compose build

build-api-gateway: ## Build API Gateway
	@echo "$(GREEN)Building API Gateway...$(NC)"
	cd api-gateway && docker build -t gopro-api-gateway .

build-user: ## Build User Service
	@echo "$(GREEN)Building User Service...$(NC)"
	cd user-service && docker build -t gopro-user-service .

build-course: ## Build Course Service
	@echo "$(GREEN)Building Course Service...$(NC)"
	cd course-service && docker build -t gopro-course-service .

build-progress: ## Build Progress Service
	@echo "$(GREEN)Building Progress Service...$(NC)"
	cd progress-service && docker build -t gopro-progress-service .

# Development Commands
dev-api-gateway: ## Run API Gateway in development mode
	cd api-gateway && go run cmd/main.go

dev-user: ## Run User Service in development mode
	cd user-service && go run cmd/main.go

dev-course: ## Run Course Service in development mode
	cd course-service && go run cmd/main.go

dev-progress: ## Run Progress Service in development mode
	cd progress-service && go run cmd/main.go

# Testing Commands
test: ## Run tests for all services
	@echo "$(GREEN)Running tests for all services...$(NC)"
	@cd api-gateway && go test ./... -v
	@cd user-service && go test ./... -v
	@cd course-service && go test ./... -v
	@cd progress-service && go test ./... -v

test-api-gateway: ## Run API Gateway tests
	cd api-gateway && go test ./... -v

test-user: ## Run User Service tests
	cd user-service && go test ./... -v

test-course: ## Run Course Service tests
	cd course-service && go test ./... -v

test-progress: ## Run Progress Service tests
	cd progress-service && go test ./... -v

test-integration: ## Run integration tests
	@echo "$(GREEN)Running integration tests...$(NC)"
	go test ./tests/integration/... -v

# Database Commands
db-migrate: ## Run database migrations
	@echo "$(GREEN)Running database migrations...$(NC)"
	docker-compose exec postgres psql -U gopro -d gopro -f /docker-entrypoint-initdb.d/init-db.sql

db-shell: ## Open PostgreSQL shell
	docker-compose exec postgres psql -U gopro -d gopro

db-reset: ## Reset database
	@echo "$(RED)WARNING: This will delete all data!$(NC)"
	@read -p "Are you sure? Type 'yes' to confirm: " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		docker-compose down -v; \
		docker-compose up -d postgres; \
		sleep 5; \
		$(MAKE) db-migrate; \
	else \
		echo "$(GREEN)Database reset cancelled$(NC)"; \
	fi

# Monitoring Commands
jaeger: ## Open Jaeger UI
	@open http://localhost:16686 || xdg-open http://localhost:16686

prometheus: ## Open Prometheus UI
	@open http://localhost:9090 || xdg-open http://localhost:9090

grafana: ## Open Grafana UI
	@open http://localhost:3000 || xdg-open http://localhost:3000

# Health Check Commands
health: ## Check health of all services
	@echo "$(GREEN)Checking service health...$(NC)"
	@echo "API Gateway:"
	@curl -s http://localhost:8080/health | jq . || echo "$(RED)Not responding$(NC)"
	@echo "\nUser Service:"
	@curl -s http://localhost:8081/health | jq . || echo "$(RED)Not responding$(NC)"
	@echo "\nCourse Service:"
	@curl -s http://localhost:8082/health | jq . || echo "$(RED)Not responding$(NC)"
	@echo "\nProgress Service:"
	@curl -s http://localhost:8083/health | jq . || echo "$(RED)Not responding$(NC)"

# Cleanup Commands
clean: ## Clean up containers, volumes, and images
	@echo "$(YELLOW)Cleaning up...$(NC)"
	docker-compose down -v
	docker system prune -f

clean-all: ## Clean up everything including images
	@echo "$(RED)WARNING: This will remove all containers, volumes, and images!$(NC)"
	@read -p "Are you sure? Type 'yes' to confirm: " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		docker-compose down -v --rmi all; \
		docker system prune -af; \
	else \
		echo "$(GREEN)Cleanup cancelled$(NC)"; \
	fi

# Code Quality Commands
lint: ## Run linters for all services
	@echo "$(GREEN)Running linters...$(NC)"
	@cd api-gateway && golangci-lint run
	@cd user-service && golangci-lint run
	@cd course-service && golangci-lint run
	@cd progress-service && golangci-lint run

fmt: ## Format code for all services
	@echo "$(GREEN)Formatting code...$(NC)"
	@cd api-gateway && go fmt ./...
	@cd user-service && go fmt ./...
	@cd course-service && go fmt ./...
	@cd progress-service && go fmt ./...

vet: ## Run go vet for all services
	@echo "$(GREEN)Running go vet...$(NC)"
	@cd api-gateway && go vet ./...
	@cd user-service && go vet ./...
	@cd course-service && go vet ./...
	@cd progress-service && go vet ./...

# Dependency Commands
deps: ## Download dependencies for all services
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	@cd api-gateway && go mod download
	@cd user-service && go mod download
	@cd course-service && go mod download
	@cd progress-service && go mod download

tidy: ## Tidy dependencies for all services
	@echo "$(GREEN)Tidying dependencies...$(NC)"
	@cd api-gateway && go mod tidy
	@cd user-service && go mod tidy
	@cd course-service && go mod tidy
	@cd progress-service && go mod tidy

# Documentation Commands
docs: ## Generate documentation
	@echo "$(GREEN)Generating documentation...$(NC)"
	@godoc -http=:6060 &
	@echo "Documentation server started at http://localhost:6060"

# Performance Commands
bench: ## Run benchmarks for all services
	@echo "$(GREEN)Running benchmarks...$(NC)"
	@cd api-gateway && go test -bench=. -benchmem ./...
	@cd user-service && go test -bench=. -benchmem ./...
	@cd course-service && go test -bench=. -benchmem ./...
	@cd progress-service && go test -bench=. -benchmem ./...

load-test: ## Run load tests
	@echo "$(GREEN)Running load tests...$(NC)"
	@echo "Not implemented yet"

# Security Commands
security-scan: ## Run security scans
	@echo "$(GREEN)Running security scans...$(NC)"
	@cd api-gateway && gosec ./...
	@cd user-service && gosec ./...
	@cd course-service && gosec ./...
	@cd progress-service && gosec ./...

# Status Commands
ps: ## Show running containers
	docker-compose ps

stats: ## Show container stats
	docker stats

# Default target
.DEFAULT_GOAL := help

