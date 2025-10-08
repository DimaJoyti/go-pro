.PHONY: help build run test clean install demo

APP_NAME=encrypt
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

run: build ## Run the application
	@./bin/$(APP_NAME)

test: ## Run tests
	@go test -v ./...

clean: ## Clean build artifacts
	@rm -rf bin/
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
	@echo "Creating test file..."
	@echo "This is a secret message!" > test.txt
	@echo "Encrypting test.txt..."
	@./bin/$(APP_NAME) encrypt --input test.txt --password demo123
	@echo ""
	@echo "Decrypting test.txt.enc..."
	@./bin/$(APP_NAME) decrypt --input test.txt.enc --password demo123
	@echo ""
	@echo "Comparing files..."
	@diff test.txt test.txt.dec && echo "✓ Files match!" || echo "✗ Files don't match"
	@rm -f test.txt test.txt.enc test.txt.dec

.DEFAULT_GOAL := help

