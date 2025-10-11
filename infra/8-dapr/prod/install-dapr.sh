#!/bin/bash

# Dapr Production Installation and Validation Script
# This script installs Dapr with Redis for production use

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/../../.." && pwd)"

# ── Switch to production context ────────────────────────────
echo "🔄 Switching to production context…"
kubectx skillsier-prod >/dev/null 2>&1 || kubectl config use-context skillsier-prod

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# --- Configuration ---
YAML_FILE="dapr-prod.yaml"
DAPR_NAMESPACE="dapr-system"
APP_NAMESPACE="skillsier"
DAPR_VERSION="1.14.4"

# --- Helper Functions ---
print_status() { echo -e "\n${BLUE}INFO:${NC} $1"; }
print_success() { echo -e "${GREEN}SUCCESS:${NC} $1"; }
print_warning() { echo -e "${YELLOW}WARNING:${NC} $1"; }
print_error() { echo -e "${RED}ERROR:${NC} $1"; exit 1; }

# --- Main Functions ---

create_dapr_namespace() {
    print_status "Applying Dapr Namespace..."
    kubectl apply -f "./dapr-namespace.yaml"
    print_success "Dapr Namespace applied."
}

check_prerequisites() {
    print_status "Checking prerequisites (kubectl, helm, dapr cli)..."
    command -v kubectl >/dev/null || print_error "kubectl is not installed."
    command -v helm >/dev/null || print_error "Helm is not installed."
    command -v dapr >/dev/null || print_error "Dapr CLI is not installed."
    if ! kubectl cluster-info &> /dev/null; then
        print_error "Cannot connect to Kubernetes cluster."
    fi
    print_success "Prerequisites are satisfied."
}

install_dapr() {
    print_status "Installing Dapr control plane (production mode with mTLS)..."

    helm repo add dapr https://dapr.github.io/helm-charts/ &>/dev/null || true
    helm repo update

    # Production Helm configuration with mTLS enabled
    helm upgrade --install dapr dapr/dapr \
    --version "$DAPR_VERSION" \
    --namespace "$DAPR_NAMESPACE" \
    --wait \
    --set global.ha.enabled=false \
    --set global.logAsJson=true \
    --set global.mtls.enabled=true \
    --set dapr_placement.replicationFactor=1 \
    --set dapr_placement.resources.limits.memory=128Mi \
    --set dapr_placement.resources.limits.cpu=250m \
    --set dapr_placement.resources.requests.memory=64Mi \
    --set dapr_placement.resources.requests.cpu=100m \
    --set dapr_operator.resources.limits.memory=256Mi \
    --set dapr_operator.resources.limits.cpu=500m \
    --set dapr_operator.resources.requests.memory=128Mi \
    --set dapr_operator.resources.requests.cpu=250m \
    --set dapr_sentry.resources.limits.memory=128Mi \
    --set dapr_sentry.resources.limits.cpu=250m \
    --set dapr_sentry.resources.requests.memory=64Mi \
    --set dapr_sentry.resources.requests.cpu=100m

    print_success "Dapr control plane installation complete."

    # Install Dapr Dashboard
    print_status "Installing Dapr Dashboard..."
    helm upgrade --install dapr-dashboard dapr/dapr-dashboard \
    --namespace "$DAPR_NAMESPACE" \
    --set resources.limits.memory=128Mi \
    --set resources.limits.cpu=250m \
    --set resources.requests.memory=64Mi \
    --set resources.requests.cpu=100m
    print_success "Dapr Dashboard installation complete."
}

applying_dapr_global_configuration() {
    print_status "Applying Dapr global configuration (production with mTLS)..."
    kubectl apply -f "./dapr-config.yaml"
    print_success "Global Dapr configuration applied."
}

validate_dapr() {
    print_status "Waiting for Dapr control plane to become healthy..."
    local timeout=300; local interval=10; local elapsed=0
    
    while ! dapr status -k &>/dev/null; do
        if [ $elapsed -ge $timeout ]; then
            print_error "Timeout: Dapr did not become healthy."
            exit 1
        fi
        echo "Dapr is not ready. Retrying in ${interval}s..."
        sleep $interval
        elapsed=$((elapsed + interval))
    done
    
    print_success "Dapr control plane is healthy."
    dapr status -k
}

deploy_redis_and_components() {
    print_status "Deploying Redis and Dapr components..."
    
    # Ensure namespace exists
    kubectl get namespace "$APP_NAMESPACE" &>/dev/null || kubectl create namespace "$APP_NAMESPACE"
    
    # Apply the Redis and component configuration
    kubectl apply -f "$YAML_FILE"
    print_success "Redis and Dapr components deployed."
    
    # Wait for Redis to be ready
    print_status "Waiting for Redis to be ready..."
    kubectl wait --for=condition=available deployment/redis-master -n "$APP_NAMESPACE" --timeout=300s
    print_success "Redis is ready."
    
    # Verify Redis connectivity
    print_status "Testing Redis connectivity..."
    kubectl run redis-test --image=redis:7-alpine --rm -i --restart=Never -n "$APP_NAMESPACE" -- \
        redis-cli -h redis-master -a Redis@12345 ping && echo "Redis connectivity verified!" || {
            print_warning "Redis connectivity test failed, but continuing..."
        }
}

verify_components() {
    print_status "Verifying Dapr components..."
    
    # Wait a bit for components to initialize
    sleep 30
    
    echo ""
    echo "=== Dapr Components Status ==="
    kubectl get components -n "$APP_NAMESPACE"
    
    echo ""
    echo "=== Redis Status ==="
    kubectl get pods -n "$APP_NAMESPACE" -l app=redis
    
    echo ""
    echo "=== Testing Dapr Component Health ==="
    # Check if we can get the metadata
    if kubectl get ingress dapr-dashboard-ingress -n dapr-system &>/dev/null; then
        print_success "Dashboard ingress is ready"
    fi
}

apply_certificates_and_ingress() {
    print_status "Setting up certificates and ingress..."
    
    # Wait for certificate to be ready
    print_status "Waiting for certificate to be issued..."
    kubectl wait --for=condition=ready certificate/tls-dapr-dashboard -n dapr-system --timeout=300s || {
        print_warning "Certificate took longer than expected to be ready"
        kubectl describe certificate tls-dapr-dashboard -n dapr-system
    }
    
    # Check final status
    print_status "Checking final status..."
    kubectl get ingress --all-namespaces | grep dapr
    kubectl get certificate -n dapr-system | grep dashboard || echo "No certificates found"
    
    # Test the endpoint
    print_status "Testing the Dapr Dashboard endpoint..."
    if curl -I https://dapr.skillsier.com --connect-timeout 10 &>/dev/null; then
        print_success "Dapr Dashboard is accessible!"
    else
        print_warning "Dapr Dashboard might need a few more minutes to be ready"
    fi
}

final_status_check() {
    echo ""
    echo "=== Final Status Check ==="
    
    print_status "Dapr Control Plane Status:"
    dapr status -k
    
    print_status "Redis Status:"
    kubectl get pods -n "$APP_NAMESPACE" -l app=redis -o wide
    
    print_status "Dapr Components:"
    kubectl get components -n "$APP_NAMESPACE"
    
    print_status "Certificates:"
    kubectl get certificate -n dapr-system
    
    print_status "Ingresses:"
    kubectl get ingress --all-namespaces | grep dapr
    
    echo ""
    print_success "✅ Dapr production installation completed!"
    echo ""
    echo "📝 Production endpoints:"
    echo "🌐 Dapr Dashboard: https://dapr.skillsier.com"
    echo "🔗 Dapr API: https://dapr-api.skillsier.com (run apply-dapr-api.sh next)"
    echo ""
    echo "📋 Test commands:"
    echo "   curl -I https://dapr.skillsier.com"
    echo ""
    echo "🔧 Next steps:"
    echo "   1. Run: ./apply-dapr-api.sh"
    echo "   2. Test component connectivity with metadata endpoint"
    echo ""
    echo "🔐 Production features enabled:"
    echo "   ✅ mTLS security"
    echo "   ✅ Redis state store"
    echo "   ✅ Redis pub/sub"
    echo "   ✅ Let's Encrypt certificates"
}

# --- Main Execution ---
main() {
    echo -e "${GREEN}🚀 Starting Dapr Production Installation 🚀${NC}"
    create_dapr_namespace
    check_prerequisites
    install_dapr
    applying_dapr_global_configuration
    validate_dapr
    deploy_redis_and_components
    verify_components
    apply_certificates_and_ingress
    final_status_check
    echo -e "\n${GREEN}✅ Dapr production installation completed successfully!${NC}"
}

main