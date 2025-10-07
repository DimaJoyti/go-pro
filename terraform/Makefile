.PHONY: help init plan apply destroy validate fmt lint clean dev staging prod

# Variables
ENV ?= dev
PLAN_FILE = $(ENV).tfplan
VAR_FILE = environments/$(ENV).tfvars

# Colors for output
RED = \033[0;31m
GREEN = \033[0;32m
YELLOW = \033[1;33m
NC = \033[0m # No Color

help: ## Display this help screen
	@echo "$(GREEN)GO-PRO Terraform Management$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(YELLOW)%-30s$(NC) %s\n", $$1, $$2}'

init: ## Initialize Terraform
	@echo "$(GREEN)Initializing Terraform...$(NC)"
	terraform init -upgrade

validate: ## Validate Terraform configuration
	@echo "$(GREEN)Validating Terraform configuration...$(NC)"
	terraform validate

fmt: ## Format Terraform files
	@echo "$(GREEN)Formatting Terraform files...$(NC)"
	terraform fmt -recursive

lint: ## Lint Terraform files with tflint
	@echo "$(GREEN)Linting Terraform files...$(NC)"
	@which tflint > /dev/null || (echo "$(RED)tflint not installed. Installing...$(NC)" && brew install tflint)
	tflint --init
	tflint

plan: validate ## Plan infrastructure changes
	@echo "$(GREEN)Planning infrastructure for $(ENV) environment...$(NC)"
	@if [ ! -f $(VAR_FILE) ]; then \
		echo "$(RED)Error: $(VAR_FILE) not found$(NC)"; \
		exit 1; \
	fi
	terraform plan -var-file=$(VAR_FILE) -out=$(PLAN_FILE)

apply: ## Apply infrastructure changes
	@echo "$(YELLOW)Applying infrastructure for $(ENV) environment...$(NC)"
	@if [ ! -f $(PLAN_FILE) ]; then \
		echo "$(RED)Error: Plan file not found. Run 'make plan ENV=$(ENV)' first$(NC)"; \
		exit 1; \
	fi
	terraform apply $(PLAN_FILE)
	@rm -f $(PLAN_FILE)

destroy: ## Destroy infrastructure (requires confirmation)
	@echo "$(RED)WARNING: This will destroy all infrastructure for $(ENV) environment!$(NC)"
	@read -p "Are you sure? Type 'yes' to confirm: " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		terraform destroy -var-file=$(VAR_FILE); \
	else \
		echo "$(GREEN)Destroy cancelled$(NC)"; \
	fi

output: ## Show Terraform outputs
	@echo "$(GREEN)Terraform outputs for $(ENV) environment:$(NC)"
	terraform output

output-json: ## Show Terraform outputs in JSON format
	@terraform output -json

refresh: ## Refresh Terraform state
	@echo "$(GREEN)Refreshing Terraform state...$(NC)"
	terraform refresh -var-file=$(VAR_FILE)

show: ## Show current state
	@terraform show

graph: ## Generate dependency graph
	@echo "$(GREEN)Generating dependency graph...$(NC)"
	terraform graph | dot -Tpng > graph.png
	@echo "$(GREEN)Graph saved to graph.png$(NC)"

clean: ## Clean Terraform files
	@echo "$(GREEN)Cleaning Terraform files...$(NC)"
	rm -rf .terraform
	rm -f .terraform.lock.hcl
	rm -f *.tfplan
	rm -f *.tfstate*
	rm -f graph.png

# Environment-specific targets
dev: ## Plan and apply for development environment
	@$(MAKE) plan ENV=dev
	@$(MAKE) apply ENV=dev

staging: ## Plan and apply for staging environment
	@$(MAKE) plan ENV=staging
	@$(MAKE) apply ENV=staging

prod: ## Plan and apply for production environment
	@echo "$(YELLOW)Production deployment requires manual approval$(NC)"
	@$(MAKE) plan ENV=production
	@read -p "Review the plan above. Type 'yes' to apply: " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		$(MAKE) apply ENV=production; \
	else \
		echo "$(GREEN)Production deployment cancelled$(NC)"; \
	fi

# State management
state-list: ## List resources in state
	@terraform state list

state-show: ## Show a specific resource (usage: make state-show RESOURCE=module.vpc.aws_vpc.main)
	@terraform state show $(RESOURCE)

state-pull: ## Pull current state
	@terraform state pull

state-backup: ## Backup current state
	@echo "$(GREEN)Backing up Terraform state...$(NC)"
	@mkdir -p backups
	@terraform state pull > backups/terraform.tfstate.$(shell date +%Y%m%d_%H%M%S)
	@echo "$(GREEN)State backed up to backups/$(NC)"

# Import existing resources
import: ## Import existing resource (usage: make import RESOURCE=module.vpc.aws_vpc.main ID=vpc-xxxxx)
	@terraform import -var-file=$(VAR_FILE) $(RESOURCE) $(ID)

# Workspace management
workspace-list: ## List workspaces
	@terraform workspace list

workspace-new: ## Create new workspace (usage: make workspace-new NAME=feature-branch)
	@terraform workspace new $(NAME)

workspace-select: ## Select workspace (usage: make workspace-select NAME=dev)
	@terraform workspace select $(NAME)

# Security and compliance
security-scan: ## Run security scan with tfsec
	@echo "$(GREEN)Running security scan...$(NC)"
	@which tfsec > /dev/null || (echo "$(RED)tfsec not installed. Installing...$(NC)" && brew install tfsec)
	tfsec .

cost-estimate: ## Estimate infrastructure costs with infracost
	@echo "$(GREEN)Estimating infrastructure costs...$(NC)"
	@which infracost > /dev/null || (echo "$(RED)infracost not installed. Visit https://www.infracost.io/docs/$(NC)" && exit 1)
	infracost breakdown --path . --terraform-var-file=$(VAR_FILE)

# Documentation
docs: ## Generate documentation with terraform-docs
	@echo "$(GREEN)Generating documentation...$(NC)"
	@which terraform-docs > /dev/null || (echo "$(RED)terraform-docs not installed. Installing...$(NC)" && brew install terraform-docs)
	terraform-docs markdown table . > TERRAFORM.md
	@echo "$(GREEN)Documentation generated in TERRAFORM.md$(NC)"

# Setup helpers
setup-backend: ## Setup S3 backend and DynamoDB table
	@echo "$(GREEN)Setting up Terraform backend...$(NC)"
	@./scripts/setup-backend.sh

setup-tools: ## Install required tools
	@echo "$(GREEN)Installing required tools...$(NC)"
	@which terraform > /dev/null || (echo "$(RED)Installing Terraform...$(NC)" && brew install terraform)
	@which tflint > /dev/null || (echo "$(RED)Installing tflint...$(NC)" && brew install tflint)
	@which tfsec > /dev/null || (echo "$(RED)Installing tfsec...$(NC)" && brew install tfsec)
	@which terraform-docs > /dev/null || (echo "$(RED)Installing terraform-docs...$(NC)" && brew install terraform-docs)
	@echo "$(GREEN)All tools installed$(NC)"

# Kubernetes helpers
kubectl-config: ## Configure kubectl for EKS cluster
	@echo "$(GREEN)Configuring kubectl...$(NC)"
	@terraform output -raw kubectl_config_command | sh
	@kubectl get nodes

k8s-dashboard: ## Get Kubernetes dashboard URL
	@echo "$(GREEN)Kubernetes Dashboard URL:$(NC)"
	@terraform output -raw cloudwatch_dashboard_url

# Database helpers
db-endpoint: ## Get database endpoint
	@terraform output -raw rds_endpoint

db-password: ## Get database password from Secrets Manager
	@echo "$(GREEN)Retrieving database password...$(NC)"
	@aws secretsmanager get-secret-value --secret-id $(shell terraform output -raw secrets_arn | grep database) --query SecretString --output text

# Monitoring helpers
logs: ## Open CloudWatch logs
	@echo "$(GREEN)Opening CloudWatch logs...$(NC)"
	@open "https://console.aws.amazon.com/cloudwatch/home?region=$(shell terraform output -json environment_info | jq -r '.region')#logsV2:log-groups"

metrics: ## Open CloudWatch metrics
	@echo "$(GREEN)Opening CloudWatch metrics...$(NC)"
	@open "https://console.aws.amazon.com/cloudwatch/home?region=$(shell terraform output -json environment_info | jq -r '.region')#metricsV2:"

dashboard: ## Open CloudWatch dashboard
	@terraform output -raw cloudwatch_dashboard_url | xargs open

# Troubleshooting
unlock: ## Force unlock state (usage: make unlock LOCK_ID=xxxxx)
	@echo "$(YELLOW)Force unlocking state...$(NC)"
	@terraform force-unlock $(LOCK_ID)

debug: ## Run Terraform with debug logging
	@echo "$(GREEN)Running Terraform with debug logging...$(NC)"
	@TF_LOG=DEBUG terraform plan -var-file=$(VAR_FILE)

# CI/CD helpers
ci-plan: ## Plan for CI/CD (no color output)
	@terraform plan -var-file=$(VAR_FILE) -out=$(PLAN_FILE) -no-color

ci-apply: ## Apply for CI/CD (auto-approve)
	@terraform apply -auto-approve $(PLAN_FILE)

ci-validate: ## Validate for CI/CD
	@terraform fmt -check -recursive
	@terraform validate
	@tflint
	@tfsec .

# Default target
.DEFAULT_GOAL := help
