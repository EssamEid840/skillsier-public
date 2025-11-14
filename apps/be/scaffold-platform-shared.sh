#!/usr/bin/env bash
# Scaffold script: creates platform-shared directory tree with empty files
# Run from repository root: ./scaffold-platform-shared.sh
set -euo pipefail

ROOT_DIR=platform-shared

mkdir -p "$ROOT_DIR"

# top-level files
mkdir -p "$ROOT_DIR"
: > "$ROOT_DIR/go.mod"
: > "$ROOT_DIR/go.sum"
: > "$ROOT_DIR/README.md"

# logging
mkdir -p "$ROOT_DIR/logging"
: > "$ROOT_DIR/logging/logger.go"
: > "$ROOT_DIR/logging/context.go"
: > "$ROOT_DIR/logging/fields.go"
: > "$ROOT_DIR/logging/config.go"

# tracing
mkdir -p "$ROOT_DIR/tracing"
: > "$ROOT_DIR/tracing/otel.go"
: > "$ROOT_DIR/tracing/tracer.go"
: > "$ROOT_DIR/tracing/span.go"
: > "$ROOT_DIR/tracing/config.go"

# metrics
mkdir -p "$ROOT_DIR/metrics"
: > "$ROOT_DIR/metrics/metrics.go"
: > "$ROOT_DIR/metrics/http.go"
: > "$ROOT_DIR/metrics/collectors.go"
: > "$ROOT_DIR/metrics/config.go"

# httpx
mkdir -p "$ROOT_DIR/httpx"
: > "$ROOT_DIR/httpx/errors.go"
: > "$ROOT_DIR/httpx/pagination.go"
: > "$ROOT_DIR/httpx/response.go"
: > "$ROOT_DIR/httpx/validator.go"

# ginx
mkdir -p "$ROOT_DIR/ginx"
: > "$ROOT_DIR/ginx/requestid.go"
: > "$ROOT_DIR/ginx/logging.go"
: > "$ROOT_DIR/ginx/recover.go"
: > "$ROOT_DIR/ginx/otel.go"
: > "$ROOT_DIR/ginx/cors.go"
: > "$ROOT_DIR/ginx/timeout.go"

# outbox
mkdir -p "$ROOT_DIR/outbox/postgres"
: > "$ROOT_DIR/outbox/publisher.go"
: > "$ROOT_DIR/outbox/forwarder.go"
: > "$ROOT_DIR/outbox/scheduler.go"
: > "$ROOT_DIR/outbox/entity.go"
: > "$ROOT_DIR/outbox/repository.go"
: > "$ROOT_DIR/outbox/postgres/repository.go"
: > "$ROOT_DIR/outbox/config.go"

# inbox
mkdir -p "$ROOT_DIR/inbox/postgres"
: > "$ROOT_DIR/inbox/checker.go"
: > "$ROOT_DIR/inbox/marker.go"
: > "$ROOT_DIR/inbox/entity.go"
: > "$ROOT_DIR/inbox/repository.go"
: > "$ROOT_DIR/inbox/postgres/repository.go"

# idempotency
mkdir -p "$ROOT_DIR/idempotency/postgres"
: > "$ROOT_DIR/idempotency/middleware.go"
: > "$ROOT_DIR/idempotency/handler.go"
: > "$ROOT_DIR/idempotency/entity.go"
: > "$ROOT_DIR/idempotency/repository.go"
: > "$ROOT_DIR/idempotency/postgres/repository.go"

# final message
echo "Scaffold created under: $ROOT_DIR"

echo "Files created (sample):"
find "$ROOT_DIR" -type f | sed 's|^| - |' | sed -n '1,200p'

exit 0
