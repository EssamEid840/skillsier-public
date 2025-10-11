#!/usr/bin/env bash
set -Eeuo pipefail

echo "==========================================
Kafka Strimzi Installation Script (KRaft Mode)
==========================================

This script will install:
- Strimzi Kafka Operator
- Kafka Cluster with KRaft mode
- SASL/SCRAM authentication
- Topics and Users for Keycloak
- (Optional) kcat producer/consumer examples
==========================================
"

# -----------------------------
# Configuration (tweak if needed)
# -----------------------------
NAMESPACE="${NAMESPACE:-kafka}"
RELEASE_NAME="${RELEASE_NAME:-strimzi-kafka-operator}"
STRIMZI_VERSION="${STRIMZI_VERSION:-0.47.0}"
TIMEOUT="${TIMEOUT:-600s}"

# Optional auto-fixes for small VPS
AUTO_BUMP_KAFKA_VERSION="${AUTO_BUMP_KAFKA_VERSION:-true}"   # if Kafka version < 3.9, change to 3.9.1 in the YAML
AUTO_FIX_STORAGE_CLASS="${AUTO_FIX_STORAGE_CLASS:-true}"     # if StorageClass not found, swap to 'local-path'

# Optional: apply kcat manifests if present
INSTALL_KCAT_EXAMPLES="${INSTALL_KCAT_EXAMPLES:-true}"
JOB_WAIT_TIMEOUT="${JOB_WAIT_TIMEOUT:-180s}"                 # wait time for kcat Jobs to complete

# -----------------------------
# Pretty logging
# -----------------------------
RED=$'\033[0;31m'
GREEN=$'\033[0;32m'
YELLOW=$'\033[1;33m'
NC=$'\033[0m'

log_info()  { echo -e "${GREEN}[INFO]${NC} $*"; }
log_warn()  { echo -e "${YELLOW}[WARN]${NC} $*"; }
log_error() { echo -e "${RED}[ERROR]${NC} $*"; }

trap 'log_error "An error occurred on line $LINENO. Exiting."' ERR

# ── Switch to prod context ──────────────────────────────────
echo "🔄 Switching to prod context…"
kubectx skillsier-prod 2>/dev/null || kubectl config use-context skillsier-prod


# -----------------------------
# Helpers
# -----------------------------
ensure_file() {
  local f="$1"
  if [[ ! -s "$f" ]]; then
    log_error "Required file not found or empty: $f (expected next to this script)"
    exit 1
  fi
}

retry_helm_repo_update() {
  local tries=0 max=5
  until helm repo update; do
    tries=$((tries+1))
    if (( tries >= max )); then
      log_error "helm repo update failed after $max attempts"
      return 1
    fi
    log_warn "helm repo update failed, retrying in 3s… ($tries/$max)"
    sleep 3
  done
}

wait_for_deploy() {
  local ns="$1" name="$2"
  kubectl -n "$ns" rollout status "deployment/$name" --timeout="$TIMEOUT"
}

# Extract Kubernetes server version (robust; prevents trap chatter)
get_kube_major_minor() {
  set +Eeuo pipefail
  local json v_major v_minor
  json="$(kubectl version -o json 2>/dev/null)"
  if echo "$json" | grep -q '"emulationMajor"'; then
    v_major="$(printf '%s' "$json" | sed -n 's/.*"emulationMajor":"\{0,1\}\([0-9]\+\).*/\1/p' | head -n1)"
    v_minor="$(printf '%s' "$json" | sed -n 's/.*"emulationMinor":"\{0,1\}\([0-9]\+\).*/\1/p' | head -n1)"
  else
    v_major="$(printf '%s' "$json" | sed -n 's/.*"major":"\{0,1\}\([0-9]\+\).*/\1/p' | head -n1)"
    v_minor="$(printf '%s' "$json" | sed -n 's/.*"minor":"\{0,1\}\([0-9]\+\).*/\1/p' | head -n1)"
  fi
  if [[ -z "$v_major" || -z "$v_minor" ]]; then
    echo "1 27"
  else
    echo "$v_major $v_minor"
  fi
  set -Eeuo pipefail
}

align_release_name_if_owned() {
  local owner
  owner="$(kubectl get clusterrole strimzi-cluster-operator-namespaced \
    -o jsonpath='{.metadata.annotations.meta\.helm\.sh/release-name}' 2>/dev/null || true)"
  if [[ -n "$owner" ]]; then
    log_info "Detected existing Strimzi cluster roles owned by Helm release: ${owner}"
    RELEASE_NAME="$owner"
  fi
  log_info "Using Helm release name: ${RELEASE_NAME}"
}

helm_upgrade_install_with_retry() {
  local tries=0 max=5 rc=0
  while true; do
    set +e
    helm upgrade --install "${RELEASE_NAME}" strimzi/strimzi-kafka-operator \
      --namespace "${NAMESPACE}" \
      --create-namespace \
      --version "${STRIMZI_VERSION}" \
      --wait \
      --atomic \
      --timeout "${TIMEOUT}"
    rc=$?
    set -e
    if (( rc == 0 )); then
      return 0
    fi
    tries=$((tries+1))
    if (( tries < max )); then
      log_warn "Helm reported a transient error (code $rc). Retrying in 10s… ($tries/$max)"
      sleep 10
      continue
    fi
    log_error "Helm upgrade/install failed after $max attempts (last code $rc)"
    return $rc
  done
}

guard_kafka_version() {
  local f="kafka-cluster.yaml"
  local current
  current="$(grep -E '^[[:space:]]*version:[[:space:]]*' "$f" | head -n1 | awk '{print $2}' || true)"
  if [[ -z "$current" ]]; then
    log_warn "Could not detect Kafka version in $f"
    return 0
  fi
  case "$current" in
    3.9.*|4.0.*)
      log_info "Kafka version in $f is $current (supported by Strimzi ${STRIMZI_VERSION})."
      ;;
    *)
      log_warn "Kafka version in $f is $current (unsupported by Strimzi ${STRIMZI_VERSION}). Supported: 3.9.0/3.9.1/4.0.0"
      if [[ "$AUTO_BUMP_KAFKA_VERSION" == "true" ]]; then
        log_info "Auto-bumping Kafka version to 3.9.1 in $f"
        sed -i '0,/^[[:space:]]*version:[[:space:]]*.*/s//    version: 3.9.1/' "$f"
      else
        log_warn "AUTO_BUMP_KAFKA_VERSION=false; leaving as-is (install may fail)."
      fi
      ;;
  esac
}

guard_storage_class() {
  local f="kafka-cluster.yaml"
  local sc
  sc="$(awk '/kind:[[:space:]]*KafkaNodePool/{flag=1} flag && /class:/{print $2; exit}' "$f" || true)"
  if [[ -z "$sc" ]]; then
    log_info "No explicit StorageClass in KafkaNodePool (will use default)."
    return 0
  fi
  if kubectl get storageclass "$sc" >/dev/null 2>&1; then
    log_info "StorageClass '$sc' exists."
  else
    log_warn "StorageClass '$sc' not found in cluster."
    if [[ "$AUTO_FIX_STORAGE_CLASS" == "true" ]]; then
      if kubectl get storageclass local-path >/dev/null 2>&1; then
        log_info "Patching StorageClass to 'local-path' in $f"
        sed -i "0,/class:[[:space:]]*$sc/s//        class: local-path/" "$f"
      else
        log_warn "'local-path' StorageClass not present; leaving as-is."
      fi
    else
      log_warn "AUTO_FIX_STORAGE_CLASS=false; leaving as-is (PVC may stay Pending)."
    fi
  fi
}

# -----------------------------
# Steps
# -----------------------------
check_prerequisites() {
  log_info "Checking prerequisites…"
  command -v kubectl >/dev/null || { log_error "kubectl is not installed"; exit 1; }
  command -v helm   >/dev/null || { log_error "helm is not installed"; exit 1; }
  kubectl cluster-info >/dev/null || { log_error "Cannot connect to Kubernetes cluster"; exit 1; }
  log_info "Prerequisites check passed"
}

create_namespace() {
  log_info "Ensuring namespace ${NAMESPACE} exists…"
  kubectl create namespace "${NAMESPACE}" --dry-run=client -o yaml | kubectl apply -f -
}

install_strimzi_operator() {
  log_info "Adding Strimzi Helm repo…"
  helm repo add strimzi https://strimzi.io/charts/ || true
  retry_helm_repo_update

  align_release_name_if_owned

  log_info "Installing/Upgrading Strimzi Operator (version ${STRIMZI_VERSION}) into namespace ${NAMESPACE}…"
  helm_upgrade_install_with_retry

  read -r KMAJOR KMINOR < <(get_kube_major_minor)
  log_info "Setting STRIMZI_KUBERNETES_VERSION=major=${KMAJOR},minor=${KMINOR} on the operator deployment"
  kubectl -n "${NAMESPACE}" set env deployment/strimzi-cluster-operator \
    STRIMZI_KUBERNETES_VERSION="major=${KMAJOR},minor=${KMINOR}" --overwrite

  log_info "Waiting for Strimzi operator deployment to be Ready…"
  wait_for_deploy "${NAMESPACE}" "strimzi-cluster-operator"
}

wait_for_kafka_ready() {
  local ns="$1" name="$2"
  log_info "Waiting for Kafka/$name (namespace: $ns) to be Ready…"
  local end=$(( $(date +%s) + ${TIMEOUT%s} ))
  while true; do
    local st
    st="$(kubectl get kafka "$name" -n "$ns" -o jsonpath='{.status.conditions[?(@.type=="Ready")].status}' 2>/dev/null || true)"
    if [[ "$st" == "True" ]]; then
      log_info "Kafka cluster is Ready"
      return 0
    fi
    if (( $(date +%s) >= end )); then
      log_error "Timed out waiting for Kafka/$name to be Ready"
      kubectl -n "$ns" get kafka "$name" -o yaml || true
      exit 1
    fi
    echo -n "."
    sleep 10
  done
}

wait_for_kafkauser_ready() {
  local ns="$1" user="$2"
  log_info "Waiting for KafkaUser/$user to be Ready…"
  local end=$(( $(date +%s) + ${TIMEOUT%s} ))
  while true; do
    local st
    st="$(kubectl get kafkauser "$user" -n "$ns" -o jsonpath='{.status.conditions[?(@.type=="Ready")].status}' 2>/dev/null || true)"
    if [[ "$st" == "True" ]]; then
      log_info "KafkaUser/$user is Ready"
      return 0
    fi
    if (( $(date +%s) >= end )); then
      log_error "Timed out waiting for KafkaUser/$user to be Ready"
      kubectl -n "$ns" get kafkauser "$user" -o yaml || true
      exit 1
    fi
    echo -n "."
    sleep 5
  done
}

deploy_kafka_cluster() {
  ensure_file "kafka-cluster.yaml"
  guard_kafka_version
  guard_storage_class
  log_info "Applying Kafka cluster (KRaft) CR…"
  kubectl -n "${NAMESPACE}" apply -f "kafka-cluster.yaml"
  wait_for_kafka_ready "${NAMESPACE}" "kafka-cluster"
}

create_users_and_topics() {
  ensure_file "kafka-users.yaml"
  log_info "Applying Kafka users/topics…"
  kubectl -n "${NAMESPACE}" apply -f "kafka-users.yaml"

  for u in keycloak-user admin-user; do
    wait_for_kafkauser_ready "${NAMESPACE}" "$u"
  done

  log_info "Extracting user credentials (if available)…"
  if kubectl -n "${NAMESPACE}" get secret keycloak-user >/dev/null 2>&1; then
    kubectl -n "${NAMESPACE}" get secret keycloak-user -o jsonpath='{.data.password}' | base64 -d > keycloak-user-password.txt || true
    log_info "Saved: keycloak-user-password.txt"
  else
    log_warn "Secret keycloak-user not found (yet)."
  fi

  if kubectl -n "${NAMESPACE}" get secret admin-user >/dev/null 2>&1; then
    kubectl -n "${NAMESPACE}" get secret admin-user -o jsonpath='{.data.password}' | base64 -d > admin-user-password.txt || true
    log_info "Saved: admin-user-password.txt"
  else
    log_warn "Secret admin-user not found (yet)."
  fi
}

# ---- kcat helpers (Jobs) ----
_wait_for_job() {
  local job="$1"
  log_info "Waiting for Job/$job to complete (timeout ${JOB_WAIT_TIMEOUT})…"
  set +e
  kubectl -n "${NAMESPACE}" wait --for=condition=complete "job/${job}" --timeout="${JOB_WAIT_TIMEOUT}"
  local rc=$?
  set -e
  if (( rc != 0 )); then
    log_warn "Job/${job} did not complete successfully within ${JOB_WAIT_TIMEOUT}"
    log_info "Describe Job/${job}:"
    kubectl -n "${NAMESPACE}" describe "job/${job}" || true
    local pod
    pod="$(kubectl -n "${NAMESPACE}" get pods --no-headers -o custom-columns=":metadata.name" --selector=job-name="${job}" | head -n1 || true)"
    if [[ -n "$pod" ]]; then
      log_info "Last 200 lines from pod/$pod:"
      kubectl -n "${NAMESPACE}" logs "$pod" --tail=200 || true
    fi
    return 1
  fi
  return 0
}

_stream_job_logs() {
  local job="$1"
  local pod
  pod="$(kubectl -n "${NAMESPACE}" get pods --no-headers -o custom-columns=":metadata.name" --selector=job-name="${job}" | head -n1 || true)"
  if [[ -n "$pod" ]]; then
    log_info "Streaming logs for pod/$pod (Job/${job})…"
    kubectl -n "${NAMESPACE}" logs "$pod" --tail=100 -f || true
  else
    log_warn "No pod found for Job/${job} to stream logs."
  fi
}

apply_kcat_examples() {
  if [[ "${INSTALL_KCAT_EXAMPLES}" != "true" ]]; then
    log_info "INSTALL_KCAT_EXAMPLES=false; skipping kcat example manifests."
    return 0
  fi

  local any=0
  if [[ -s "kcat-producer.yaml" ]]; then
    log_info "Applying kcat-producer.yaml…"
    kubectl -n "${NAMESPACE}" apply -f "kcat-producer.yaml"
    any=1
  else
    log_warn "kcat-producer.yaml not found; skipping."
  fi

  if [[ -s "kcat-consumer.yaml" ]]; then
    log_info "Applying kcat-consumer.yaml…"
    kubectl -n "${NAMESPACE}" apply -f "kcat-consumer.yaml"
    any=1
  else
    log_warn "kcat-consumer.yaml not found; skipping."
  fi

  if (( any == 1 )); then
    log_info "kcat example manifests applied."

    # Known default job names (adjust here if you change metadata.name in your YAMLs)
    local jobs=()
    kubectl -n "${NAMESPACE}" get jobs -o name | grep -E '^job.batch/kcat-(producer|consumer)$' >/dev/null 2>&1 && \
      jobs=("kcat-producer" "kcat-consumer")

    # Stream logs quickly (non-blocking-ish) then wait for completion
    for j in "${jobs[@]}"; do
      _stream_job_logs "$j" || true
    done

    local failures=0
    for j in "${jobs[@]}"; do
      _wait_for_job "$j" || failures=$((failures+1))
    done

    if (( failures > 0 )); then
      log_warn "One or more kcat jobs failed or timed out (${failures}). See logs above."
    else
      log_info "All kcat jobs completed successfully."
    fi

    cat <<'HINT'
Quick hints:
- List pods:      kubectl -n kafka get pods -l app=kcat
- Check Job logs: kubectl -n kafka logs job/kcat-producer --all-containers
                  kubectl -n kafka logs job/kcat-consumer --all-containers
- Re-run:         kubectl -n kafka delete job kcat-producer kcat-consumer && \
                  kubectl -n kafka apply -f kcat-producer.yaml -f kcat-consumer.yaml
HINT
  fi
}

summary() {
  cat <<EOF

==========================================
Kafka Strimzi Installation Complete!

Cluster Details:
- Namespace: ${NAMESPACE}
- Cluster Name: kafka-cluster
- Mode: KRaft (No ZooKeeper)
- Bootstrap Service: kafka-cluster-kafka-bootstrap.${NAMESPACE}.svc
- Authentication: SASL/SCRAM (per your CR + users)

Users Created:
- keycloak-user (password -> keycloak-user-password.txt if secret present)
- admin-user    (password -> admin-user-password.txt if secret present)

Notes:
- If your YAML still had Kafka 3.6.x, it was auto-bumped to 3.9.1 (supported by Strimzi ${STRIMZI_VERSION}).
- If your NodePool referenced a missing StorageClass, it was switched to 'local-path'.
- kcat examples applied: ${INSTALL_KCAT_EXAMPLES}

==========================================
EOF
}

start_live_consumer() {
  kubectl -n "$NAMESPACE" apply -f kcat-live-consumer.yaml || true
  kubectl -n "$NAMESPACE" wait --for=condition=Ready pod/kcat-live-consumer --timeout=60s || true
}

main() {
  # ensure_file "kafka-cluster.yaml"
  # ensure_file "kafka-users.yaml"
  # # kcat files are optional; we’ll apply them if present (and INSTALL_KCAT_EXAMPLES=true)

  check_prerequisites
  create_namespace
  install_strimzi_operator
  deploy_kafka_cluster
  create_users_and_topics
  start_live_consumer
  apply_kcat_examples
  summary
}

main "$@"
