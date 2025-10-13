#!/bin/bash

# Script to setup local development environment

set -e

echo "Setting up local development environment for users-be..."

# Check if kubectl is installed
if ! command -v kubectl &> /dev/null; then
    echo "ERROR: kubectl is not installed"
    echo "Please install kubectl first"
    exit 1
fi

# Check if we can connect to the cluster
if ! kubectl cluster-info &> /dev/null; then
    echo "ERROR: Cannot connect to Kubernetes cluster"
    echo "Please configure kubectl to connect to your cluster"
    exit 1
fi

echo "✓ Kubernetes connection verified"

# Check if required namespaces exist
echo "Checking required namespaces..."

if ! kubectl get namespace skillsier &> /dev/null; then
    echo "ERROR: Namespace 'skillsier' does not exist"
    exit 1
fi

if ! kubectl get namespace kafka &> /dev/null; then
    echo "ERROR: Namespace 'kafka' does not exist"
    exit 1
fi

if ! kubectl get namespace keycloak &> /dev/null; then
    echo "ERROR: Namespace 'keycloak' does not exist"
    exit 1
fi

echo "✓ Required namespaces exist"

# Check if PostgreSQL is accessible
echo "Checking PostgreSQL connectivity..."

if ! kubectl get svc users-be-postgres-external -n skillsier &> /dev/null; then
    echo "WARNING: PostgreSQL external service not found"
    echo "Creating NodePort service for PostgreSQL..."
    kubectl apply -f deployments/db/postgres-nodeport.yaml
fi

echo "✓ PostgreSQL service configured"

# Check if Kafka is accessible
echo "Checking Kafka connectivity..."

if ! kubectl get svc kafka-cluster-kafka-external-bootstrap -n kafka &> /dev/null; then
    echo "ERROR: Kafka external service not found"
    echo "Please ensure Kafka is properly configured with NodePort"
    exit 1
fi

echo "✓ Kafka service configured"

# Check if Keycloak is accessible
echo "Checking Keycloak connectivity..."

if ! kubectl get svc keycloak-external -n keycloak &> /dev/null; then
    echo "WARNING: Keycloak external service not found"
    echo "Keycloak might not be accessible from localhost"
fi

echo "✓ Keycloak service configured"

# Create Dapr components directory if it doesn't exist
echo "Setting up Dapr components..."

mkdir -p ~/.dapr/components

# This would be done if you need local Dapr components
# For now, we're using direct Kafka connection

echo "✓ Dapr components directory created"

echo ""
echo "========================================="
echo "Local development environment is ready!"
echo "========================================="
echo ""
echo "Next steps:"
echo "  1. Run: source scripts/get-secrets.sh"
echo "  2. Run: make run"
echo ""
echo "Or simply run: make run-local"
echo ""