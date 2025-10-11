#!/usr/bin/env bash
set -Eeuo pipefail

# Simple cluster/VPS resource diagnostic focused on Strimzi + Kafka in one namespace.
# Usage: ./cluster-diagnose.sh [namespace]
NS="${1:-kafka}"

section() {
  echo
  echo "=============================================="
  echo "$1"
  echo "=============================================="
}

has_cmd() { command -v "$1" >/dev/null 2>&1; }

section "[0] Context"
kubectl config current-context || true
echo "Namespace: $NS"

section "[1] Nodes overview"
kubectl get nodes -o wide || true

section "[2] Node conditions & allocatable"
# Show all nodes conditions & allocatable resources
for n in $(kubectl get nodes -o name | cut -d'/' -f2); do
  echo "--- Node: $n ---"
  kubectl describe node "$n" | egrep -i "Name:|Roles:|OS Image:|Kernel Version:|Container Runtime:|Allocatable:|Capacity:|MemoryPressure|DiskPressure|PIDPressure" || true
done

section "[3] Metrics (if metrics-server installed)"
if kubectl get apiservices | grep -qi metrics.k8s.io; then
  echo "metrics.k8s.io APIService detected"
  kubectl top nodes || true
  kubectl -n "$NS" top pods || true
else
  echo "metrics.k8s.io APIService NOT detected (install metrics-server to see 'kubectl top' output)"
fi

section "[4] Namespace resource usage / quotas / limits"
kubectl -n "$NS" get resourcequota,limitrange || true

section "[5] Strimzi operator & Kafka pods status"
kubectl -n "$NS" get pods -o wide || true
echo
echo "Recent events:"
kubectl -n "$NS" get events --sort-by=.lastTimestamp | tail -n 50 || true

OP_LABEL="name=strimzi-cluster-operator"
section "[6] Operator describe & probe failures"
kubectl -n "$NS" get pod -l "$OP_LABEL" || true
for p in $(kubectl -n "$NS" get pod -l "$OP_LABEL" -o name 2>/dev/null || true); do
  echo "--- $p ---"
  kubectl -n "$NS" describe "$p" | egrep -i "Reason:|Message:|Liveness|Readiness|Last State|Exit Code|OOMKilled" || true
done

section "[7] Kafka pods & PVCs"
# Kafka broker StatefulSet pods are labeled strimzi.io/name like <cluster>-kafka
kubectl -n "$NS" get pods -l 'strimzi.io/name in (kafka-cluster-kafka)' -o wide || true
echo
echo "PVCs in namespace:"
kubectl -n "$NS" get pvc || true
echo
echo "StorageClasses:"
kubectl get sc || true

section "[8] Common blocking conditions quick check"
echo "- Any pods Pending due to resources?"
kubectl -n "$NS" get pods --field-selector=status.phase=Pending || true
echo
echo "- Any pods Waiting with 'Insufficient cpu/memory' or ImagePullBackOff?"
kubectl -n "$NS" get pods | grep -E "Pending|CrashLoopBackOff|OOMKilled|ImagePullBackOff|ErrImagePull|CreateContainerConfigError" || true

section "[9] Operator logs (last 200 lines)"
if kubectl -n "$NS" get deploy/strimzi-cluster-operator >/dev/null 2>&1; then
  kubectl -n "$NS" logs deploy/strimzi-cluster-operator --tail=200 || true
else
  echo "Operator deployment not found"
fi

section "[10] Kafka CR status"
if kubectl -n "$NS" get kafka.kafka.strimzi.io kafka-cluster >/dev/null 2>&1; then
  kubectl -n "$NS" get kafka.kafka.strimzi.io kafka-cluster -o jsonpath='{.status.conditions}' || true
  echo
else
  echo "Kafka CR kafka-cluster not found"
fi

section "[11] Disk space on nodes (if using k3s or single-node, may reflect VPS disk)"
# Try to exec into an operator pod to run df, if permitted (often not); otherwise, print hint.
OP_POD=$(kubectl -n "$NS" get pod -l "$OP_LABEL" -o jsonpath='{.items[0].metadata.name}' 2>/dev/null || true)
if [[ -n "${OP_POD:-}" ]]; then
  echo "Trying to run 'df -h' inside operator pod container (may fail if restricted)..."
  kubectl -n "$NS" exec "$OP_POD" -- df -h 2>/dev/null || echo "Could not exec df -h inside pod."
fi
echo "Hint: also check VPS host: 'df -h', 'free -h', 'uptime'."

section "[12] Summary hints"
cat <<'HINTS'
- If you see MemoryPressure/DiskPressure on the node, reduce pod memory limits/requests or free space on the VPS.
- If PVCs are Pending with 'no storageclass', change 'class:' in kafka-cluster.yaml to an existing StorageClass.
- If pods show OOMKilled, increase memory limit or lower JVM -Xmx.
- If CrashLoopBackOff with liveness probe failing, check container logs and lower CPU/memory requests.
- For single-node Kafka, make sure replication factors and min.insync.replicas are 1.
HINTS

echo
echo "Done."
