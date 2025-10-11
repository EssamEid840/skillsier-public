#!/bin/bash
set -e

echo "==========================================
Kafka Strimzi Cleanup Script
==========================================
"

NAMESPACE="kafka"
RELEASE_NAME="strimzi-kafka"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info() { echo -e "${GREEN}[INFO]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

wait_gone() {
  local kind="$1" sel="$2" to="${3:-180s}"
  log_info "Waiting for ${kind} (${sel}) to be deleted…"
  kubectl -n "$NAMESPACE" wait --for=delete "$kind" -l "$sel" --timeout="$to" 2>/dev/null || true
}

# Cleanup function
cleanup() {
    log_warn "This will delete all Kafka resources. Are you sure? (yes/no)"
    read -r response
    if [[ "$response" != "yes" ]]; then
        log_info "Cleanup cancelled"
        exit 0
    fi
    
    log_info "Starting cleanup..."

    # --- kcat producer/consumer (no files; delete deployed resources directly) ---
    log_info "Deleting kcat producer/consumer (Deployments/Jobs/Pods)…"
    kubectl -n "$NAMESPACE" delete deploy kcat-producer kcat-consumer --ignore-not-found=true || true
    kubectl -n "$NAMESPACE" delete job kcat-producer kcat-consumer --ignore-not-found=true || true
    kubectl -n "$NAMESPACE" delete pod -l app=kcat --ignore-not-found=true || true
    wait_gone pod "app=kcat" 120s

    
    # Delete Kafka resources
    log_info "Deleting Kafka users and topics..."
    kubectl delete -f kafka-users.yaml -n $NAMESPACE --ignore-not-found=true
    
    log_info "Deleting Kafka cluster..."
    kubectl delete -f kafka-cluster.yaml -n $NAMESPACE --ignore-not-found=true
    
    # Wait for Kafka pods to terminate
    log_info "Waiting for Kafka pods to terminate..."
    kubectl wait --for=delete pod -l strimzi.io/cluster=kafka-cluster -n $NAMESPACE --timeout=300s || true
    
    # # Delete Strimzi operator
    # log_info "Deleting Strimzi operator..."
    # helm uninstall $RELEASE_NAME -n $NAMESPACE || true
    
    # Delete PVCs
    log_info "Deleting Persistent Volume Claims..."
    kubectl delete pvc -l strimzi.io/cluster=kafka-cluster -n $NAMESPACE --ignore-not-found=true
    
    # # Delete namespace
    # log_warn "Delete namespace $NAMESPACE? (yes/no)"
    # read -r response
    # if [[ "$response" == "yes" ]]; then
    #     log_info "Deleting namespace..."
    #     kubectl delete namespace $NAMESPACE --ignore-not-found=true
    # fi
    
    # Clean up credential files
    log_info "Cleaning up credential files..."
    rm -f keycloak-user-password.txt admin-user-password.txt
    
    log_info "==========================================
Cleanup Complete!
=========================================="
}

cleanup