.PHONY: help deps build run test clean docker-build docker-push k8s-deploy

# ==========================================
# Variables
# ==========================================
APP_NAME=<service>-be
VERSION?=latest
DOCKER_REGISTRY?=skillsier
DOCKER_IMAGE=$(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
GO_VERSION=1.23
K8S_NAMESPACE?=skillsier

# ==========================================
# Help
# ==========================================
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# ==========================================
# Dependencies
# ==========================================
deps: ## Install Go dependencies
	@echo "→ Installing dependencies..."
	@go mod download
	@go mod tidy
	@go mod verify
	@echo "✓ Dependencies installed"

# ==========================================
# Build
# ==========================================
build: ## Build the application binary for Linux
	@echo "→ Building $(APP_NAME)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME) -ldflags="-w -s" cmd/api/main.go
	@echo "✓ Build complete: bin/$(APP_NAME)"

build-local: ## Build for local OS
	@echo "→ Building $(APP_NAME) for local..."
	@go build -o bin/$(APP_NAME) cmd/api/main.go
	@echo "✓ Build complete: bin/$(APP_NAME)"

# ==========================================
# Run
# ==========================================
run: ## Run the application locally
	@echo "→ Running $(APP_NAME)..."
	@go run cmd/api/main.go

run-local: deps build-local ## Setup local environment and run (connects to K8s infra)
	@echo "→ Starting $(APP_NAME) in local mode..."
	@echo "ℹ This will connect to K8s-hosted infrastructure (Kafka, Postgres, Keycloak)"
	@echo "ℹ Make sure you have exported required environment variables"
	@echo "ℹ Run 'source scripts/get-secrets.sh' if you haven't already"
	@./bin/$(APP_NAME)

# ==========================================
# Testing
# ==========================================
test: ## Run unit tests
	@echo "→ Running unit tests..."
	@go test -v -race -coverprofile=coverage.out ./tests/unit/...
	@go tool cover -func=coverage.out
	@echo "✓ Tests complete"

test-coverage: test ## Run tests with coverage report
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report generated: coverage.html"

test-integration: ## Run integration tests (requires K8s infrastructure)
	@echo "→ Running integration tests..."
	@echo "⚠ These tests require connection to K8s infrastructure"
	@INTEGRATION_TEST=true go test -v ./tests/integration/...
	@echo "✓ Integration tests complete"

# ==========================================
# Code Quality
# ==========================================
fmt: ## Format Go code
	@echo "→ Formatting code..."
	@go fmt ./...
	@echo "✓ Code formatted"

lint: ## Run linter
	@echo "→ Running linter..."
	@golangci-lint run ./...
	@echo "✓ Linting complete"

vet: ## Run go vet
	@echo "→ Running go vet..."
	@go vet ./...
	@echo "✓ Vet complete"

# ==========================================
# Clean
# ==========================================
clean: ## Clean build artifacts
	@echo "→ Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "✓ Clean complete"

# ==========================================
# Docker
# ==========================================
docker-build: ## Build Docker image
	@echo "→ Building Docker image: $(DOCKER_IMAGE)"
	@docker build -t $(DOCKER_IMAGE) -f Dockerfile .
	@docker tag $(DOCKER_IMAGE) $(DOCKER_REGISTRY)/$(APP_NAME):latest
	@echo "✓ Docker image built: $(DOCKER_IMAGE)"

docker-push: docker-build ## Push Docker image to registry
	@echo "→ Pushing Docker image: $(DOCKER_IMAGE)"
	@docker push $(DOCKER_IMAGE)
	@docker push $(DOCKER_REGISTRY)/$(APP_NAME):latest
	@echo "✓ Docker image pushed"

# ==========================================
# Kubernetes
# ==========================================
k8s-apply: ## Apply all K8s manifests
	@echo "→ Applying K8s manifests..."
	@kubectl apply -k k8s/overlays/local -n $(K8S_NAMESPACE)
	@echo "✓ K8s manifests applied"

k8s-deploy: docker-push k8s-apply ## Build, push, and deploy to K8s
	@echo "→ Deploying $(APP_NAME) to K8s..."
	@kubectl rollout status deployment/$(APP_NAME) -n $(K8S_NAMESPACE)
	@echo "✓ Deployment complete"

k8s-delete: ## Delete K8s deployment
	@echo "→ Deleting $(APP_NAME) from K8s..."
	@kubectl delete -k k8s/overlays/local -n $(K8S_NAMESPACE)
	@echo "✓ Deleted from K8s"

k8s-logs: ## View K8s logs
	@kubectl logs -f -l app=$(APP_NAME) -n $(K8S_NAMESPACE)

k8s-status: ## Check K8s deployment status
	@kubectl get all -l app=$(APP_NAME) -n $(K8S_NAMESPACE)

k8s-restart: ## Restart K8s deployment
	@echo "→ Restarting $(APP_NAME)..."
	@kubectl rollout restart deployment/$(APP_NAME) -n $(K8S_NAMESPACE)
	@kubectl rollout status deployment/$(APP_NAME) -n $(K8S_NAMESPACE)
	@echo "✓ Deployment restarted"

k8s-shell: ## Get a shell in the running pod
	@kubectl exec -it $$(kubectl get pod -l app=$(APP_NAME) -n $(K8S_NAMESPACE) -o jsonpath='{.items[0].metadata.name}') -n $(K8S_NAMESPACE) -- /bin/sh

# ==========================================
# Database
# ==========================================
db-migrate: ## Run database migrations
	@echo "→ Running database migrations..."
	@echo "ℹ Auto-migrations run on startup. Set AUTO_MIGRATE=true to enable."
	@echo "ℹ For manual migrations, see migrations/README.md"

# ==========================================
# Development Helpers
# ==========================================
port-forward-postgres: ## Port-forward PostgreSQL (local:5432 -> k8s:5432)
	@echo "→ Port-forwarding PostgreSQL..."
	@echo "ℹ PostgreSQL will be available at localhost:5432"
	@kubectl port-forward svc/<service>-be-postgres 5432:5432 -n $(K8S_NAMESPACE)

port-forward-kafka: ## Port-forward Kafka (local:9092 -> k8s:9092)
	@echo "→ Port-forwarding Kafka..."
	@echo "ℹ Kafka will be available at localhost:9092"
	@kubectl port-forward svc/kafka-external 9092:9092 -n kafka

generate-proto: ## Generate protobuf code
	@echo "→ Generating protobuf code..."
	@cd contracts/events && buf generate
	@echo "✓ Protobuf code generated"

# ==========================================
# Quick Commands
# ==========================================
dev: deps ## Start development environment
	@echo "→ Starting development environment..."
	@echo "ℹ Make sure K8s infrastructure is running"
	@make run

all: clean deps test build ## Run all checks and build
	@echo "✓ All checks passed and build complete"

# ==========================================
# Notes
# ==========================================
# To run locally against K8s infrastructure:
#   1. Export secrets: source scripts/get-secrets.sh
#   2. Run: make run-local
#
# To run integration tests:
#   INTEGRATION_TEST=true make test-integration
#
# To deploy to K8s:
#   make k8s-deploy