#!/usr/bin/env bash
# scaffold-empty-users-be.sh
set -euo pipefail

ROOT="${1:-users-be}"
FORCE="${FORCE:-0}"  # 0 = keep existing files as-is, 1 = truncate to empty

# --- helper: mkdir -p for each path ---
mkd() { mkdir -p "$1"; }

# --- helper: create empty file (or truncate if FORCE=1) ---
mkempty() {
  local path="$1"
  mkd "$(dirname "$path")"
  if [[ -e "$path" && "$FORCE" != "1" ]]; then
    echo "↪︎ exists (kept): $path"
  else
    : > "$path"   # create or truncate to empty
    echo "✓ empty: $path"
  fi
}

echo "→ Scaffolding empty structure under: $ROOT"

# --- directories ---
DIRS=(
  "cmd/api"
  "internal/domain/user"
  "internal/domain/outbox"
  "internal/application/user"
  "internal/application/eventhandler"
  "internal/infrastructure/persistence/postgres"
  "internal/infrastructure/messaging/kafka"
  "internal/infrastructure/outbox"
  "internal/interfaces/http/handlers"
  "internal/interfaces/http/middleware"
  "internal/config"
  "deployments/k8s"
  "deployments/local/dapr-components"
  "scripts"
)
for d in "${DIRS[@]}"; do mkd "$ROOT/$d"; done

# --- files (all created empty) ---
FILES=(
  # cmd/
  "cmd/api/main.go"

  # internal/domain
  "internal/domain/user/entity.go"
  "internal/domain/user/repository.go"
  "internal/domain/outbox/entity.go"
  "internal/domain/outbox/repository.go"

  # internal/application
  "internal/application/user/service.go"
  "internal/application/user/dto.go"
  "internal/application/user/mapper.go"
  "internal/application/eventhandler/keycloak_handler.go"

  # infrastructure/persistence/postgres
  "internal/infrastructure/persistence/postgres/connection.go"
  "internal/infrastructure/persistence/postgres/migrations.go"
  "internal/infrastructure/persistence/postgres/user_repository.go"
  "internal/infrastructure/persistence/postgres/outbox_repository.go"

  # infrastructure/messaging/kafka
  "internal/infrastructure/messaging/kafka/consumer.go"
  "internal/infrastructure/messaging/kafka/producer.go"
  "internal/infrastructure/messaging/kafka/scram.go"

  # infrastructure/outbox
  "internal/infrastructure/outbox/processor.go"

  # interfaces/http
  "internal/interfaces/http/handlers/user_handler.go"
  "internal/interfaces/http/handlers/health_handler.go"
  "internal/interfaces/http/middleware/logging.go"
  "internal/interfaces/http/middleware/cors.go"
  "internal/interfaces/http/router.go"

  # config
  "internal/config/config.go"

  # deployments/k8s
  "deployments/k8s/deployment.yaml"
  "deployments/k8s/service.yaml"
  "deployments/k8s/configmap.yaml"
  "deployments/k8s/secret.yaml"
  "deployments/k8s/dapr-component-kafka.yaml"
  "deployments/k8s/dapr-component-state.yaml"

  # deployments/local
  "deployments/local/dapr-components/.gitkeep"

  # scripts
  "scripts/setup-local.sh"
  "scripts/get-secrets.sh"

  # repo root
  ".gitignore"
  ".env.example"
  "Dockerfile"
  "Makefile"
  "go.mod"
  "go.sum"
  "README.md"
  "GETTING_STARTED.md"
  "FILE_STRUCTURE.md"
)

for f in "${FILES[@]}"; do mkempty "$ROOT/$f"; done

echo "→ Done."
command -v tree >/dev/null 2>&1 && { echo; tree -a "$ROOT"; } || true
