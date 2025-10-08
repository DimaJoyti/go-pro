.PHONY: help build run test clean db-setup db-migrate db-reset

APP_NAME=blog-engine
DB_NAME=blogdb
DB_URL=postgres://localhost/$(DB_NAME)?sslmode=disable

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) ./cmd/server/main.go
	@echo "✓ Build complete: bin/$(APP_NAME)"

run: build ## Run the application
	@echo "Starting $(APP_NAME)..."
	@./bin/$(APP_NAME)

test: ## Run tests
	@go test -v ./...

clean: ## Clean build artifacts
	@rm -rf bin/
	@echo "✓ Clean complete"

db-setup: ## Create database
	@echo "Creating database $(DB_NAME)..."
	@createdb $(DB_NAME) || echo "Database already exists"
	@echo "✓ Database ready"

db-migrate: ## Run database migrations
	@echo "Running migrations..."
	@psql $(DB_URL) < migrations/001_init.sql
	@echo "✓ Migrations complete"

db-reset: ## Reset database
	@echo "Resetting database..."
	@dropdb $(DB_NAME) || true
	@make db-setup
	@make db-migrate
	@echo "✓ Database reset complete"

dev: db-setup db-migrate run ## Setup and run in development mode

.DEFAULT_GOAL := help

