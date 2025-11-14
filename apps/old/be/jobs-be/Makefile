.PHONY: help build run test clean docker-build k8s-deploy

help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  run           - Run locally"
	@echo "  test          - Run tests"
	@echo "  clean         - Clean build artifacts"
	@echo "  docker-build  - Build Docker image"
	@echo "  k8s-deploy    - Deploy to Kubernetes"

build:
	go build -o bin/jobs-be cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t skillsier/jobs-be:latest .

k8s-deploy:
	kubectl apply -f deployments/k8s/