.PHONY: help build run test clean docker-build docker-push k8s-deploy k8s-delete deps migrate

# Variables
APP_NAME=users-be
VERSION?=latest
DOCKER_REGISTRY?=your-registry
DOCKER_IMAGE=$(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
GO_VERSION=1.23
KNS ?= skillsier

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Install dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@go mod verify
	@echo "✓ Dependencies installed"

db.apply:
	kubectl apply -f deployments/db/postgres-secret.yaml
	kubectl apply -f deployments/db/postgres-statefulset.yaml
	kubectl apply -f deployments/db/postgres-service.yaml
	kubectl apply -f deployments/db/postgres-nodeport.yaml
	@kubectl -n $(KNS) rollout status sts/users-be-postgres

build: ## Build the application binary
	@echo "Building $(APP_NAME)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME) -ldflags="-w -s" cmd/api/main.go
	@echo "✓ Build complete: bin/$(APP_NAME)"

build-local: ## Build for local OS
	@echo "Building $(APP_NAME) for local..."
	@go build -o bin/$(APP_NAME) cmd/api/main.go
	@echo "✓ Build complete: bin/$(APP_NAME)"

run: ## Run the application locally
	@echo "Running $(APP_NAME)..."
	@go run cmd/api/main.go

run-local: setup-local build-local ## Setup local environment and run
	@echo "Starting $(APP_NAME) in local mode..."
	@./scripts/get-secrets.sh
	@./bin/$(APP_NAME)

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@echo "✓ Tests complete"

test-integration: ## Run integration tests
	@echo "Running integration tests..."
	@go test -v -tags=integration ./tests/integration/...
	@echo "✓ Integration tests complete"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out
	@echo "✓ Clean complete"

docker-build: ## Build Docker image
	@echo "Building Docker image: $(DOCKER_IMAGE)"
	@docker build -t $(DOCKER_IMAGE) -f Dockerfile .
	@echo "✓ Docker image built: $(DOCKER_IMAGE)"

docker-push: docker-build ## Push Docker image to registry
	@echo "Pushing Docker image: $(DOCKER_IMAGE)"
	@docker push $(DOCKER_IMAGE)
	@echo "✓ Docker image pushed"

docker-run: ## Run application in Docker locally
	@echo "Running Docker container..."
	@docker run --rm -p 8080:8080 \
		-e DB_PASSWORD=${DB_PASSWORD} \
		-e KAFKA_PASSWORD=${KAFKA_PASSWORD} \
		$(DOCKER_IMAGE)

# Kubernetes commands
k8s-deploy: ## Deploy to Kubernetes
	@echo "Deploying to Kubernetes..."
	@kubectl apply -f deployments/k8s/namespace.yaml
	@kubectl apply -f deployments/k8s/secret.yaml
	@kubectl apply -f deployments/k8s/configmap.yaml
	@kubectl apply -f deployments/k8s/deployment.yaml
	@kubectl apply -f deployments/k8s/service.yaml
	@kubectl apply -f deployments/k8s/dapr-component-kafka.yaml
	@kubectl apply -f deployments/k8s/dapr-component-state.yaml
	@echo "✓ Deployed to Kubernetes"

k8s-delete: ## Delete from Kubernetes
	@echo "Deleting from Kubernetes..."
	@kubectl delete -f deployments/k8s/ --ignore-not-found=true
	@echo "✓ Deleted from Kubernetes"

k8s-logs: ## View Kubernetes logs
	@kubectl logs -f -n skillsier -l app=$(APP_NAME)

k8s-status: ## Check Kubernetes deployment status
	@kubectl get all -n skillsier -l app=$(APP_NAME)

k8s-restart: ## Restart Kubernetes deployment
	@kubectl rollout restart deployment/$(APP_NAME) -n skillsier

# Database commands
db-reset: ## Reset database (drop and recreate)
	@echo "Resetting database..."
	@PGPASSWORD=${DB_PASSWORD} psql -h 173.212.218.251 -p 30432 -U usersuser -d usersdb -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
	@echo "✓ Database reset complete"

# Local development setup
setup-local: ## Setup local development environment
	@echo "Setting up local development environment..."
	@chmod +x scripts/setup-local.sh
	@chmod +x scripts/get-secrets.sh
	@./scripts/setup-local.sh
	@echo "✓ Local environment setup complete"

# Code quality
lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./...
	@echo "✓ Linting complete"

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Code formatted"

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...
	@echo "✓ Vet complete"

# Development helpers
dev: ## Run in development mode with hot reload (requires air)
	@air

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✓ Tools installed"

# Default target
.DEFAULT_GOAL := help