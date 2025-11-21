#!/bin/bash
set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${GREEN}[INFO]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }
log_step() { echo -e "${BLUE}==>${NC} $1"; }

# Configuration
SERVICE_NAME="users-be"
POSTGRES_LOCAL_PORT=5432
KAFKA_LOCAL_PORT=9092
KEYCLOAK_LOCAL_PORT=8080

# K8s configuration
POSTGRES_SERVICE="users-be-postgres.skillsier.svc.cluster.local"
POSTGRES_PORT=5432
KAFKA_SERVICE="kafka-cluster-kafka-bootstrap.kafka.svc.cluster.local"
KAFKA_PORT=9092
KEYCLOAK_SERVICE="keycloak.keycloak.svc.cluster.local"
KEYCLOAK_PORT=8080

# Namespaces
NS_SKILLSIER="skillsier"
NS_KAFKA="kafka"
NS_KEYCLOAK="keycloak"

# PID file for cleanup
PID_FILE="/tmp/${SERVICE_NAME}-port-forwards.pid"

cleanup() {
  log_step "Cleaning up port-forwards..."
  
  if [[ -f "$PID_FILE" ]]; then
    while read -r pid; do
      if ps -p "$pid" > /dev/null 2>&1; then
        log_info "Killing port-forward PID: $pid"
        kill "$pid" 2>/dev/null || true
      fi
    done < "$PID_FILE"
    rm -f "$PID_FILE"
  fi
  
  # Kill any remaining kubectl port-forward processes for this service
  pkill -f "kubectl port-forward.*${SERVICE_NAME}" 2>/dev/null || true
  pkill -f "kubectl port-forward.*kafka.*${KAFKA_LOCAL_PORT}" 2>/dev/null || true
  pkill -f "kubectl port-forward.*postgres.*${POSTGRES_LOCAL_PORT}" 2>/dev/null || true
  pkill -f "kubectl port-forward.*keycloak.*${KEYCLOAK_LOCAL_PORT}" 2>/dev/null || true
  
  log_info "✓ Cleanup complete"
}

# Trap cleanup on exit
trap cleanup EXIT INT TERM

check_k8s() {
  log_step "Checking Kubernetes connectivity..."
  
  if ! kubectl cluster-info > /dev/null 2>&1; then
    log_error "Cannot connect to Kubernetes cluster"
    log_error "Make sure kubectl is configured and cluster is accessible"
    exit 1
  fi
  
  log_info "✓ Kubernetes cluster is accessible"
}

check_secrets() {
  log_step "Checking if secrets are loaded..."
  
  local missing_secrets=()
  
  if [[ -z "${USERS_BE_DATABASE_PASSWORD}" ]]; then
    missing_secrets+=("DATABASE_PASSWORD")
  fi
  
  if [[ -z "${USERS_BE_KAFKA_SASL_PASSWORD}" ]]; then
    missing_secrets+=("KAFKA_SASL_PASSWORD")
  fi
  
  if [[ -z "${USERS_BE_KEYCLOAK_CLIENT_SECRET}" ]]; then
    missing_secrets+=("KEYCLOAK_CLIENT_SECRET")
  fi
  
  if [[ ${#missing_secrets[@]} -gt 0 ]]; then
    log_warn "Missing secrets: ${missing_secrets[*]}"
    log_warn "Run: source scripts/get-secrets.sh"
    log_warn "Or manually export the required environment variables"
    echo ""
    read -p "Continue anyway? (yes/no): " -r
    if [[ ! $REPLY =~ ^yes$ ]]; then
      exit 1
    fi
  else
    log_info "✓ All secrets are loaded"
  fi
}

setup_port_forwards() {
  log_step "Setting up port-forwards to K8s services..."
  
  # Clear PID file
  > "$PID_FILE"
  
  # PostgreSQL port-forward
  log_info "Port-forwarding PostgreSQL (localhost:${POSTGRES_LOCAL_PORT} -> ${POSTGRES_SERVICE}:${POSTGRES_PORT})..."
  kubectl port-forward -n "$NS_SKILLSIER" "svc/users-be-postgres" "${POSTGRES_LOCAL_PORT}:${POSTGRES_PORT}" > /dev/null 2>&1 &
  echo $! >> "$PID_FILE"
  
  # Kafka port-forward
  log_info "Port-forwarding Kafka (localhost:${KAFKA_LOCAL_PORT} -> ${KAFKA_SERVICE}:${KAFKA_PORT})..."
  kubectl port-forward -n "$NS_KAFKA" svc/kafka-cluster-kafka-bootstrap "${KAFKA_LOCAL_PORT}:${KAFKA_PORT}" > /dev/null 2>&1 &
  echo $! >> "$PID_FILE"
  
  # Keycloak port-forward (optional, comment if not needed)
  log_info "Port-forwarding Keycloak (localhost:${KEYCLOAK_LOCAL_PORT} -> ${KEYCLOAK_SERVICE}:${KEYCLOAK_PORT})..."
  kubectl port-forward -n "$NS_KEYCLOAK" svc/keycloak "${KEYCLOAK_LOCAL_PORT}:${KEYCLOAK_PORT}" > /dev/null 2>&1 &
  echo $! >> "$PID_FILE"
  
  # Wait for port-forwards to be ready
  log_info "Waiting for port-forwards to be ready..."
  sleep 3
  
  # Verify port-forwards
  local all_ready=true
  
  if ! nc -z localhost "$POSTGRES_LOCAL_PORT" 2>/dev/null; then
    log_warn "PostgreSQL port-forward may not be ready"
    all_ready=false
  fi
  
  if ! nc -z localhost "$KAFKA_LOCAL_PORT" 2>/dev/null; then
    log_warn "Kafka port-forward may not be ready"
    all_ready=false
  fi
  
  if ! nc -z localhost "$KEYCLOAK_LOCAL_PORT" 2>/dev/null; then
    log_warn "Keycloak port-forward may not be ready"
    all_ready=false
  fi
  
  if [[ "$all_ready" == true ]]; then
    log_info "✓ All port-forwards are ready"
  else
    log_warn "Some port-forwards may not be ready, but continuing..."
  fi
}

run_application() {
  log_step "Building and running application..."
  
  # Build the application
  log_info "Building $SERVICE_NAME..."
  go build -o "./bin/$SERVICE_NAME" ./cmd/api
  
  # Set environment variables for local K8s connectivity
  export USERS_BE_DATABASE_HOST=localhost
  export USERS_BE_DATABASE_PORT=$POSTGRES_LOCAL_PORT
  export USERS_BE_KAFKA_BROKERS=localhost:$KAFKA_LOCAL_PORT
  export USERS_BE_KEYCLOAK_URL=http://localhost:$KEYCLOAK_LOCAL_PORT
  export USERS_BE_APP_ENVIRONMENT=development
  export USERS_BE_DATABASE_AUTO_MIGRATE=true
  
  echo ""
  log_info "Starting $SERVICE_NAME with local K8s connectivity..."
  log_info "PostgreSQL: localhost:$POSTGRES_LOCAL_PORT"
  log_info "Kafka:      localhost:$KAFKA_LOCAL_PORT"
  log_info "Keycloak:   localhost:$KEYCLOAK_LOCAL_PORT"
  echo ""
  log_info "Press Ctrl+C to stop..."
  echo ""
  
  # Run the application
  "./bin/$SERVICE_NAME"
}

main() {
  log_step "Running $SERVICE_NAME locally against K8s infrastructure"
  echo ""
  
  check_k8s
  check_secrets
  setup_port_forwards
  run_application
}

# Show usage
if [[ "$1" == "--help" ]] || [[ "$1" == "-h" ]]; then
  cat <<EOF
Usage: $0

This script runs the $SERVICE_NAME microservice locally while connecting
to the Kubernetes infrastructure (PostgreSQL, Kafka, Keycloak).

Prerequisites:
  1. kubectl configured and connected to K8s cluster
  2. Secrets loaded (run: source scripts/get-secrets.sh)
  3. Go installed (1.21+)
  4. nc (netcat) installed for port checking

What it does:
  1. Checks K8s connectivity
  2. Verifies secrets are loaded
  3. Sets up port-forwards to K8s services:
     - PostgreSQL:  localhost:${POSTGRES_LOCAL_PORT}
     - Kafka:       localhost:${KAFKA_LOCAL_PORT}
     - Keycloak:    localhost:${KEYCLOAK_LOCAL_PORT}
  4. Builds and runs the application locally
  5. Cleans up port-forwards on exit

Environment Variables:
  You can override ports by setting:
    POSTGRES_LOCAL_PORT (default: 5432)
    KAFKA_LOCAL_PORT (default: 9092)
    KEYCLOAK_LOCAL_PORT (default: 8080)

Example:
  $0
  POSTGRES_LOCAL_PORT=5433 $0

EOF
  exit 0
fi

main