#!/usr/bin/env bash
set -Eeuo pipefail

echo "==========================================
Keycloak Installation Script (with Kafka Integration)
==========================================

This script will install:
- PostgreSQL database for Keycloak
- Kafka Bridge (HTTP to Kafka)
- Keycloak with Webhook Plugin
- Network policies
==========================================
"

# -----------------------------
# Configuration
# -----------------------------
NAMESPACE="${NAMESPACE:-keycloak}"
KAFKA_NAMESPACE="${KAFKA_NAMESPACE:-kafka}"
TIMEOUT="${TIMEOUT:-600s}"
DOMAIN="${DOMAIN:-keycloak.skillsier.com}"

# -----------------------------
# Pretty logging
# -----------------------------
RED=$'\033[0;31m'
GREEN=$'\033[0;32m'
YELLOW=$'\033[1;33m'
BLUE=$'\033[0;34m'
NC=$'\033[0m'

log_info()  { echo -e "${GREEN}[INFO]${NC} $*"; }
log_warn()  { echo -e "${YELLOW}[WARN]${NC} $*"; }
log_error() { echo -e "${RED}[ERROR]${NC} $*"; }
log_step()  { echo -e "${BLUE}[STEP]${NC} $*"; }

trap 'log_error "An error occurred on line $LINENO. Exiting."' ERR

# -----------------------------
# Helper Functions
# -----------------------------
ensure_file() {
  local f="$1"
  if [[ ! -s "$f" ]]; then
    log_error "Required file not found or empty: $f"
    exit 1
  fi
}

wait_for_deployment() {
  local ns="$1" name="$2"
  log_info "Waiting for deployment/$name in namespace $ns to be ready..."
  kubectl -n "$ns" rollout status "deployment/$name" --timeout="$TIMEOUT"
}

wait_for_statefulset() {
  local ns="$1" name="$2"
  log_info "Waiting for statefulset/$name in namespace $ns to be ready..."
  kubectl -n "$ns" rollout status "statefulset/$name" --timeout="$TIMEOUT"
}

wait_for_secret() {
  local ns="$1" secret_name="$2"
  log_info "Waiting for secret/$secret_name in namespace $ns..."
  local end=$(( $(date +%s) + ${TIMEOUT%s} ))
  while true; do
    if kubectl -n "$ns" get secret "$secret_name" >/dev/null 2>&1; then
      log_info "Secret $secret_name is ready"
      return 0
    fi
    if (( $(date +%s) >= end )); then
      log_error "Timed out waiting for secret/$secret_name"
      exit 1
    fi
    echo -n "."
    sleep 3
  done
}

# -----------------------------
# Checks
# -----------------------------
check_prerequisites() {
  log_step "Checking prerequisites..."
  
  command -v kubectl >/dev/null || { log_error "kubectl is not installed"; exit 1; }
  kubectl cluster-info >/dev/null || { log_error "Cannot connect to Kubernetes cluster"; exit 1; }
  
  # Check if Kafka namespace exists
  if ! kubectl get namespace "$KAFKA_NAMESPACE" >/dev/null 2>&1; then
    log_error "Kafka namespace '$KAFKA_NAMESPACE' not found. Please install Kafka first."
    exit 1
  fi
  
  # Check if Kafka cluster exists
  if ! kubectl -n "$KAFKA_NAMESPACE" get kafka kafka-cluster >/dev/null 2>&1; then
    log_error "Kafka cluster 'kafka-cluster' not found in namespace '$KAFKA_NAMESPACE'"
    log_error "Please run the Kafka installation script first"
    exit 1
  fi
  
  # Check if keycloak-user exists in Kafka
  if ! kubectl -n "$KAFKA_NAMESPACE" get kafkauser keycloak-user >/dev/null 2>&1; then
    log_error "KafkaUser 'keycloak-user' not found in namespace '$KAFKA_NAMESPACE'"
    log_error "Please ensure Kafka users are created"
    exit 1
  fi
  
  log_info "Prerequisites check passed ✓"
}

check_files() {
  log_step "Checking required files..."
  ensure_file "keycloak-deployment.yaml"
  ensure_file "postgres-plain.yaml"
  
  # Check if ingress.yaml exists (optional)
  if [[ -f "ingress.yaml" ]]; then
    log_info "Found ingress.yaml file"
  else
    log_warn "ingress.yaml not found - ingress will not be configured"
  fi
  
  log_info "All required files present ✓"
}

# -----------------------------
# Installation Steps
# -----------------------------
create_namespace() {
  log_step "Creating namespace '$NAMESPACE'..."
  kubectl create namespace "$NAMESPACE" --dry-run=client -o yaml | kubectl apply -f -
}

install_postgresql() {
  log_step "Installing PostgreSQL for Keycloak..."
  
  # Apply PostgreSQL manifests
  kubectl apply -f postgres-plain.yaml
  
  # Wait for PostgreSQL to be ready
  wait_for_statefulset "$NAMESPACE" "keycloak-postgresql"
  
  log_info "PostgreSQL installation complete ✓"
}

copy_kafka_secret() {
  log_step "Copying Kafka user credentials to Keycloak namespace..."

  wait_for_secret "$KAFKA_NAMESPACE" "keycloak-user"

  local pass_b64 pass jaas_b64
  pass_b64=$(kubectl -n "$KAFKA_NAMESPACE" get secret keycloak-user -o jsonpath='{.data.password}' 2>/dev/null || true)
  if [[ -n "$pass_b64" ]]; then
    pass="$(echo "$pass_b64" | base64 -d)"
  else
    jaas_b64=$(kubectl -n "$KAFKA_NAMESPACE" get secret keycloak-user -o jsonpath='{.data.sasl\.jaas\.config}' 2>/dev/null || true)
    if [[ -n "$jaas_b64" ]]; then
      pass="$(echo "$jaas_b64" | base64 -d | sed -n 's/.*password="\([^"]*\)".*/\1/p')"
    fi
  fi

  if [[ -z "$pass" ]]; then
    log_error "Could not extract password from kafka/keycloak-user secret"
    exit 1
  fi

  kubectl -n "$NAMESPACE" create secret generic keycloak-user \
    --from-literal=password="$pass" \
    --dry-run=client -o yaml | kubectl apply -f -

  kubectl -n "$NAMESPACE" get secret keycloak-user -o jsonpath='{.data.password}' >/dev/null || {
    log_error "Failed to create secret in $NAMESPACE namespace"
    exit 1
  }
  log_info "Kafka credentials copied successfully ✓"
}


update_ingress_domain() {
  log_step "Checking ingress configuration..."
  
  # Check if ingress.yaml file exists
  if [[ ! -f "ingress.yaml" ]]; then
    log_warn "ingress.yaml file not found - skipping ingress configuration"
    return 0
  fi
  
  # Check if ingress already exists
  if kubectl -n "$NAMESPACE" get ingress keycloak-ingress >/dev/null 2>&1; then
    log_info "Ingress 'keycloak-ingress' already exists"
    
    # Verify it points to the correct service
    local current_service=$(kubectl -n "$NAMESPACE" get ingress keycloak-ingress -o jsonpath='{.spec.rules[0].http.paths[0].backend.service.name}' 2>/dev/null || echo "")
    local current_host=$(kubectl -n "$NAMESPACE" get ingress keycloak-ingress -o jsonpath='{.spec.rules[0].host}' 2>/dev/null || echo "")
    
    if [[ "$current_service" == "keycloak" ]] && [[ "$current_host" == *"$DOMAIN"* || "$DOMAIN" == "keycloak.yourdomain.com" ]]; then
      log_info "Ingress is correctly configured ✓"
      log_info "  Service: $current_service"
      log_info "  Host: $current_host"
    else
      log_warn "Ingress exists but points to different service/host:"
      log_warn "  Current service: $current_service (expected: keycloak)"
      log_warn "  Current host: $current_host"
      
      if [[ "$current_service" != "keycloak" ]]; then
        log_info "Updating ingress to point to 'keycloak' service..."
        kubectl -n "$NAMESPACE" patch ingress keycloak-ingress --type merge -p '{"spec":{"rules":[{"host":"'$current_host'","http":{"paths":[{"path":"/","pathType":"Prefix","backend":{"service":{"name":"keycloak","port":{"number":8080}}}}]}}]}}' || log_warn "Failed to patch ingress"
      fi
    fi
  else
    log_info "No existing ingress found"
    log_info "To create ingress, apply ingress.yaml manually with your domain"
    log_info "  Update the domain in ingress.yaml, then run:"
    log_info "  kubectl apply -f ingress.yaml"
  fi
}

deploy_keycloak() {
  log_step "Deploying Keycloak with webhook plugin..."
  
  # Verify the secret exists before deploying
  log_info "Verifying Kafka secret is available..."
  if ! kubectl -n "$NAMESPACE" get secret keycloak-user >/dev/null 2>&1; then
    log_error "Secret keycloak-user not found in $NAMESPACE namespace"
    log_error "Cannot proceed without Kafka credentials"
    exit 1
  fi
  
  # Apply Keycloak manifests (without ingress since it's in separate file)
  kubectl apply -f keycloak-deployment.yaml
  
  # Wait for Kafka Bridge to be ready first
  log_info "Waiting for Kafka Bridge..."
  wait_for_deployment "$NAMESPACE" "kafka-bridge"
  

  # Wait for StatefulSet instead of Deployment
  wait_for_statefulset "$NAMESPACE" "keycloak"

  log_info "Keycloak deployment complete ✓"
}

verify_installation() {
  log_step "Verifying installation..."
  
  echo ""
  log_info "=== Component Status ==="
  
  # Check PostgreSQL
  if kubectl -n "$NAMESPACE" get statefulset keycloak-postgresql -o jsonpath='{.status.readyReplicas}' | grep -q "1"; then
    log_info "✓ PostgreSQL: Running"
  else
    log_warn "⚠ PostgreSQL: Not ready"
  fi
  
  # Check Kafka Bridge
  if kubectl -n "$NAMESPACE" get deployment kafka-bridge -o jsonpath='{.status.readyReplicas}' | grep -q "1"; then
    log_info "✓ Kafka Bridge: Running"
  else
    log_warn "⚠ Kafka Bridge: Not ready"
  fi
  
  # Check Keycloak
  if kubectl -n "$NAMESPACE" get deployment keycloak -o jsonpath='{.status.readyReplicas}' | grep -q "1"; then
    log_info "✓ Keycloak: Running"
  else
    log_warn "⚠ Keycloak: Not ready"
  fi
  
  echo ""
}

copy_kafka_ca_configmap() {
  log_step "Copying Kafka cluster CA certificate to Keycloak namespace..."
  # Wait for Strimzi CA secret to exist in kafka ns
  local ca_name="kafka-cluster-cluster-ca-cert"
  local tries=30
  while ! kubectl -n "$KAFKA_NAMESPACE" get secret "$ca_name" >/dev/null 2>&1; do
    ((tries--)) || { log_error "Timed out waiting for $ca_name in $KAFKA_NAMESPACE"; exit 1; }
    sleep 2
  done

  # Create/Update configmap with ca.crt under key 'ca.crt'
  kubectl -n "$NAMESPACE" create configmap kafka-ca-cert \
    --from-literal=ca.crt="$(kubectl -n "$KAFKA_NAMESPACE" get secret "$ca_name" -o jsonpath='{.data.ca\.crt}' | base64 -d)" \
    --dry-run=client -o yaml | kubectl apply -f -

  kubectl -n "$NAMESPACE" get configmap kafka-ca-cert >/dev/null 2>&1 \
    && log_info "Kafka CA ConfigMap created ✓" \
    || { log_error "Failed to create kafka-ca-cert ConfigMap"; exit 1; }
}


create_test_topic() {
  log_step "Creating Keycloak events topic in Kafka..."
  
  cat <<EOF | kubectl -n "$KAFKA_NAMESPACE" apply -f -
apiVersion: kafka.strimzi.io/v1beta2
kind: KafkaTopic
metadata:
  name: keycloak-events
  labels:
    strimzi.io/cluster: kafka-cluster
spec:
  partitions: 3
  replicas: 1
  config:
    retention.ms: 604800000  # 7 days
    compression.type: snappy
    cleanup.policy: delete
EOF
  
  log_info "Kafka topic created ✓"
}



copy_tls_certificate() {
  log_step "Copying TLS wildcard certificate to Keycloak namespace..."
  
  # Check if wildcard cert exists in skillsier namespace
  if kubectl get secret tls-wildcard -n skillsier >/dev/null 2>&1; then
    log_info "Found wildcard certificate in skillsier namespace"
    
    # Copy and rename the certificate
    kubectl get secret tls-wildcard -n skillsier -o yaml | \
      sed 's/namespace: skillsier/namespace: keycloak/' | \
      sed 's/name: tls-wildcard/name: keycloak-tls/' | \
      kubectl apply -f -
    
    # # Delete any failed certificate resources
    # kubectl -n "$NAMESPACE" delete certificate keycloak-tls --ignore-not-found=true 2>/dev/null || true
    
    log_info "TLS certificate copied successfully ✓"
  else
    log_warn "Wildcard certificate not found in skillsier namespace"
    log_warn "HTTPS access may not work until certificate is issued"
  fi
}



display_summary() {
  local KEYCLOAK_ADMIN_PASS="admin"
  local POSTGRES_PASS
  POSTGRES_PASS=$(kubectl -n "$NAMESPACE" get secret keycloak-postgresql -o jsonpath='{.data.password}' | base64 -d)
  
  local KAFKA_USER_PASS
  if kubectl -n "$KAFKA_NAMESPACE" get secret keycloak-user >/dev/null 2>&1; then
    KAFKA_USER_PASS=$(kubectl -n "$KAFKA_NAMESPACE" get secret keycloak-user -o jsonpath='{.data.password}' | base64 -d)
  else
    KAFKA_USER_PASS="<not available>"
  fi
  
  cat <<EOF

==========================================
Keycloak Installation Complete! 🎉
==========================================

📍 Namespace: $NAMESPACE

🔐 Keycloak Admin Console:
   URL: http://$DOMAIN (or https if you have cert-manager)
   Username: admin
   Password: $KEYCLOAK_ADMIN_PASS

🗄️  PostgreSQL:
   Host: keycloak-postgresql.$NAMESPACE.svc:5432
   Database: keycloak
   Username: keycloak
   Password: $POSTGRES_PASS

🌉 Kafka Bridge:
   Service: kafka-bridge.$NAMESPACE.svc:3000
   Endpoint: http://kafka-bridge.$NAMESPACE.svc:3000/events

📨 Kafka Integration:
   Broker: kafka-cluster-kafka-bootstrap.$KAFKA_NAMESPACE.svc:9094
   Topic: keycloak-events
   Username: keycloak-user
   Password: $KAFKA_USER_PASS

🔧 Useful Commands:
   # View Keycloak logs
   kubectl -n $NAMESPACE logs -f deployment/keycloak

   # View Kafka Bridge logs
   kubectl -n $NAMESPACE logs -f deployment/kafka-bridge

   # Check Keycloak health
   kubectl -n $NAMESPACE get pods -l app=keycloak

   # Port forward to access locally
   kubectl -n $NAMESPACE port-forward svc/keycloak 8080:8080

   # Consume Keycloak events from Kafka
   kubectl -n $KAFKA_NAMESPACE run kafka-consumer -ti --rm --restart=Never \\
     --image=quay.io/strimzi/kafka:0.47.0-kafka-3.9.1 -- \\
     bin/kafka-console-consumer.sh \\
     --bootstrap-server kafka-cluster-kafka-bootstrap:9092 \\
     --topic keycloak-events \\
     --from-beginning

   # Test webhook manually
   kubectl -n $NAMESPACE run curl --rm -ti --restart=Never \\
     --image=curlimages/curl -- \\
     -X POST http://kafka-bridge:3000/events \\
     -H "Content-Type: application/json" \\
     -H "Authorization: Basic $(echo -n 'webhook_user:webhook_secret_2024' | base64)" \\
     -d '{"type":"TEST","message":"Hello Kafka"}'

📚 Next Steps:
   1. Access Keycloak admin console at http://$DOMAIN
   2. Create a test realm and user
   3. Try logging in to generate events
   4. Monitor events in Kafka topic 'keycloak-events'

⚠️  Note: If using a custom domain, ensure DNS is configured and
    cert-manager is installed for TLS certificates.

==========================================
EOF
}

# -----------------------------
# Main Execution
# -----------------------------
main() {
  check_prerequisites
  check_files
  
  create_namespace
  install_postgresql
  copy_kafka_secret        # Copy BEFORE deploying Keycloak/Bridge
  copy_kafka_ca_configmap    # <-- add this call
  copy_tls_certificate
  update_ingress_domain
  deploy_keycloak
  create_test_topic
  
  verify_installation
  display_summary
  
  log_info "Installation completed successfully! 🚀"
}

# Run main function
main "$@"