#!/usr/bin/env bash
set -Eeuo pipefail

# ---------------------------------------
# Config
# ---------------------------------------
NAMESPACE="${NAMESPACE:-kafka}"
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")"/../../.. && pwd)"
SHARED_DIR="${ROOT_DIR}/infra/3_data/event-streaming/shared"
TOPICS_DIR="${SHARED_DIR}/topics"
KAFKA_DIR="${SHARED_DIR}/kafka"
TOOLS_DIR="${SHARED_DIR}/kafka/tools"


KAFKA_CR="${ROOT_DIR}/infra/3_data/event-streaming/shared/kafka/kafka-persistent.yaml"

# ── Determine Kafka CLUSTER_NAME once, before first use ────
if command -v yq >/dev/null 2>&1 && [ -f "$KAFKA_CR" ]; then
  CLUSTER_NAME="$(yq eval '.metadata.name // "small-kafka"' "$KAFKA_CR" 2>/dev/null || echo small-kafka)"
else
  CLUSTER_NAME="${CLUSTER_NAME:-small-kafka}"
fi
: "${CLUSTER_NAME:=small-kafka}"

echo "ℹ️  Cleaning Kafka cluster name: ${CLUSTER_NAME}"


# Inputs (environment flags)
# FULL_RESET=true  => wipe PVCs and PVs + delete CRs (Kafka + NodePool)
# KEEP_NAMESPACE=1 => do not delete namespace (we never delete it by default)
FULL_RESET="${FULL_RESET:-false}"

# ---------------------------------------
# Utils
# ---------------------------------------
info(){ echo -e "ℹ️  $*"; }
ok(){   echo -e "✅ $*"; }
warn(){ echo -e "⚠️  $*"; }
die(){  echo -e "❌ $*" >&2; exit 1; }

need_bin() {
  command -v "$1" >/dev/null 2>&1 || die "Missing dependency: $1"
}

need_bin kubectl
need_bin yq || warn "yq not installed; we'll fall back when needed"

# ---------------------------------------
# Resolve cluster name
# ---------------------------------------
# 1) prefer live CR
CLUSTER_NAME="$(kubectl -n "$NAMESPACE" get kafka -o jsonpath='{.items[0].metadata.name}' 2>/dev/null || true)"
# 2) fallback to file (if exists)
if [[ -z "${CLUSTER_NAME}" && -f "${KAFKA_DIR}/kafka-persistent.yaml" && "$(command -v yq)" ]]; then
  CLUSTER_NAME="$(yq eval '.metadata.name' "${KAFKA_DIR}/kafka-persistent.yaml" 2>/dev/null || true)"
fi
# 3) default
CLUSTER_NAME="${CLUSTER_NAME:-small-kafka}"

info "Using namespace: ${NAMESPACE}"
info "Detected cluster name: ${CLUSTER_NAME}"

# ---------------------------------------
# Helper: scale nodepool
# ---------------------------------------
scale_nodepool() {
  local replicas="$1"
  if kubectl -n "$NAMESPACE" get kafkanodepool default >/dev/null 2>&1; then
    info "Scaling KafkaNodePool/default -> replicas=${replicas}"
    kubectl -n "$NAMESPACE" patch kafkanodepool default --type merge -p "{\"spec\":{\"replicas\":${replicas}}}" || true
  fi
}

# ---------------------------------------
# Helper: wait for broker pod gone
# ---------------------------------------
wait_broker_gone() {
  info "Waiting for broker pod deletion…"
  kubectl -n "$NAMESPACE" wait pod -l "strimzi.io/cluster=${CLUSTER_NAME},strimzi.io/name=${CLUSTER_NAME}-kafka" \
    --for=delete --timeout=5m || true
}

# ---------------------------------------
# Helper: wipe broker PVC contents safely
# ---------------------------------------
wipe_broker_pvc() {
  info "🧹 Wiping broker PVC data (non-destructive to the PVC/PV objects)…"
  # Ensure broker is stopped
  scale_nodepool 0
  wait_broker_gone

  # Run the wipe pod
  kubectl -n "$NAMESPACE" delete pod "wipe-${CLUSTER_NAME}-default-0" --ignore-not-found || true
  kubectl -n "$NAMESPACE" apply -f "${TOOLS_DIR}/wipe-broker-pvc.yaml"
  kubectl -n "$NAMESPACE" wait pod/"wipe-${CLUSTER_NAME}-default-0" --for=condition=Ready --timeout=90s || true
  kubectl -n "$NAMESPACE" logs "pod/wipe-${CLUSTER_NAME}-default-0" --tail=200 || true
  kubectl -n "$NAMESPACE" delete pod "wipe-${CLUSTER_NAME}-default-0" --ignore-not-found || true
}

# ---------------------------------------
# Helper: delete exporter and hardening
# ---------------------------------------
delete_exporter_and_hardening() {
  info "Deleting Kafka exporter (Deployment/Service) if exists…"
  kubectl -n "$NAMESPACE" delete deploy kafka-exporter --ignore-not-found
  kubectl -n "$NAMESPACE" delete svc kafka-exporter --ignore-not-found

  info "Deleting PDBs/NetworkPolicy…"
  kubectl -n "$NAMESPACE" delete pdb kafka-pdb --ignore-not-found
  kubectl -n "$NAMESPACE" delete pdb entity-operator-pdb --ignore-not-found
  kubectl -n "$NAMESPACE" delete networkpolicy kafka-restrict-access --ignore-not-found
}

# ---------------------------------------
# Helper: delete topics & users created by apply script
# ---------------------------------------
delete_topics_and_users() {
  info "Deleting Kafka topics created by the apply script (if present)…"
  kubectl -n "$NAMESPACE" delete kafkatopic users-topic --ignore-not-found
  kubectl -n "$NAMESPACE" delete kafkatopic payments-topic --ignore-not-found

  # Optional: Keycloak topics file (only if you applied it)
  if [[ -f "${TOPICS_DIR}/keycloak-topics.yaml" ]]; then
    warn "Found ${TOPICS_DIR}/keycloak-topics.yaml → deleting resources from it (ignore errors)…"
    kubectl -n "$NAMESPACE" delete -f "${TOPICS_DIR}/keycloak-topics.yaml" --ignore-not-found || true
  fi

  info "Deleting Kafka users created by the apply script (if present)…"
  kubectl -n "$NAMESPACE" delete kafkauser client-user --ignore-not-found
  kubectl -n "$NAMESPACE" delete kafkauser client-user-tls --ignore-not-found
  kubectl -n "$NAMESPACE" delete kafkauser keycloak-producer --ignore-not-found

  info "Deleting user secrets (if any remain)…"
  kubectl -n "$NAMESPACE" delete secret client-user --ignore-not-found
  kubectl -n "$NAMESPACE" delete secret client-user-tls --ignore-not-found
  kubectl -n "$NAMESPACE" delete secret keycloak-producer --ignore-not-found
}

# ---------------------------------------
# Helper: remove Kafka CRs (non-destructive by default)
# ---------------------------------------
delete_kafka_crs() {
  info "Deleting Kafka CR and NodePool (garbage collection will remove SPS/EO, etc.)…"
  kubectl -n "$NAMESPACE" delete kafka "${CLUSTER_NAME}" --ignore-not-found
  kubectl -n "$NAMESPACE" delete kafkanodepool default --ignore-not-found
}

# ---------------------------------------
# Helper: remove SPS/pod if they linger
# ---------------------------------------
delete_lingering_runtime() {
  info "Cleaning StrimziPodSet + broker pod if still present…"
  kubectl -n "$NAMESPACE" delete sps "${CLUSTER_NAME}-default" --ignore-not-found || true
  kubectl -n "$NAMESPACE" delete pod "${CLUSTER_NAME}-default-0" --ignore-not-found || true
}

# ---------------------------------------
# Helper: nuke persistent state (when FULL_RESET=true)
# ---------------------------------------
nuke_persistent() {
  warn "FULL_RESET=true → removing persistent state (PVC/PV) and CA-related secrets"
  # delete PVC used by broker
  kubectl -n "$NAMESPACE" delete pvc -l "strimzi.io/cluster=${CLUSTER_NAME},strimzi.io/name=${CLUSTER_NAME}-kafka" --ignore-not-found || true

  # delete cluster/user CA certs (they will be re-created by operator next run)
  kubectl -n "$NAMESPACE" delete secret "${CLUSTER_NAME}-cluster-ca-cert" --ignore-not-found || true
  kubectl -n "$NAMESPACE" delete secret "${CLUSTER_NAME}-clients-ca-cert" --ignore-not-found || true
}

# ---------------------------------------
# Main
# ---------------------------------------
info "Switching to prod context (assuming already configured)…"
# (If you need to force context: kubectl config use-context skillsier-prod )

delete_exporter_and_hardening
delete_topics_and_users

# If you're trying to clear cluster.id problems without losing PVC:
if [[ "${FULL_RESET}" != "true" ]]; then
  warn "FULL_RESET is not set → leaving CRs in place and wiping broker data only."
  wipe_broker_pvc
  ok "PVC contents wiped. You can re-run the apply script now."
  exit 0
fi

# FULL_RESET=true → delete CRs and persistent state
delete_kafka_crs
delete_lingering_runtime
nuke_persistent

ok "Clean completed. Strimzi operator left running in '${NAMESPACE}'."
