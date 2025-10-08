.PHONY: help build run test clean install

APP_NAME=weather
INSTALL_PATH=/usr/local/bin

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
	@go run ./cmd/main.go

test: ## Run tests
	@go test -v ./...

clean: ## Clean build artifacts
	@rm -rf bin/
	@rm -rf ~/.weather-cli/cache
	@echo "✓ Clean complete"

install: build ## Install the application
	@echo "Installing $(APP_NAME) to $(INSTALL_PATH)..."
	@sudo cp bin/$(APP_NAME) $(INSTALL_PATH)/
	@echo "✓ Installed to $(INSTALL_PATH)/$(APP_NAME)"

uninstall: ## Uninstall the application
	@echo "Uninstalling $(APP_NAME)..."
	@sudo rm -f $(INSTALL_PATH)/$(APP_NAME)
	@echo "✓ Uninstalled"

demo: build ## Run demo
	@echo "Running weather demo..."
	@./bin/$(APP_NAME) current --city "London"

.DEFAULT_GOAL := help

