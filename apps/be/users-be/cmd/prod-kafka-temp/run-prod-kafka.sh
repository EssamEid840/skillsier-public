#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"

# adjust paths as needed
ROOT_DIR="$(git rev-parse --show-toplevel)"
CA_FILE="${ROOT_DIR}/kafka-cluster-ca.crt"

export APP_KAFKA__ENABLED=true
export APP_KAFKA__HOST="173.212.218.251:32629"
export APP_KAFKA__TOPIC="keycloak.realm-events"
export APP_KAFKA__CONSUMER_GROUP="users-be-keycloak-consumer"
export APP_KAFKA__SECURITY_PROTOCOL="SASL_SSL"
export APP_KAFKA__SASL_MECHANISM="SCRAM-SHA-512"
export APP_KAFKA__SASL_USERNAME="users-be-consumer"
export APP_KAFKA__SASL_PASSWORD="$(kubectl -n kafka get secret users-be-consumer -o jsonpath='{.data.password}' | base64 -d)"
export APP_KAFKA__SSL_CA_LOCATION="${CA_FILE}"

cd "${ROOT_DIR}/apps/be/users-be"
go run ./cmd/all-in-one
