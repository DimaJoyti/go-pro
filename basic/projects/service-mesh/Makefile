.PHONY: help build test clean proto

APP_NAME=service-mesh

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o bin/mesh ./cmd/mesh/main.go
	@echo "✓ Build complete: bin/mesh"

test: ## Run tests
	@go test -v ./...

clean: ## Clean build artifacts
	@rm -rf bin/
	@echo "✓ Clean complete"

proto: ## Generate protobuf code
	@echo "Generating protobuf code..."
	@protoc --go_out=. --go-grpc_out=. proto/*.proto
	@echo "✓ Protobuf generation complete"

.DEFAULT_GOAL := help

