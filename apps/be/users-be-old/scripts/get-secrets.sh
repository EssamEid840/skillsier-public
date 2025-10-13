#!/usr/bin/env bash
# Source-safe get-secrets.sh
# - If sourced: exports variables into current shell; never exits the terminal.
# - If executed: validates/prints summary; use --print-env to emit export lines.

# ── detect sourced vs executed ─────────────────────────────────────────────
if [[ "${BASH_SOURCE[0]}" != "$0" ]]; then __SOURCED=1; else __SOURCED=0; fi
# Only harden shell in executed mode (so we don't kill parent shell on errors)
if (( ! __SOURCED )); then set -euo pipefail; fi

# ── cli flags ──────────────────────────────────────────────────────────────
PRINT_ENV=0
NO_VERIFY=0
for a in "$@"; do
  case "$a" in
    --print-env) PRINT_ENV=1 ;;
    --no-verify) NO_VERIFY=1 ;;
  esac
done

# ── helpers ────────────────────────────────────────────────────────────────
color() { local c=$1; shift; printf "\033[%sm%s\033[0m" "$c" "$*"; }
info()  { echo "$(color '36' '[INFO]') $*"; }
ok()    { echo "$(color '32' '[OK]  ') $*"; }
warn()  { echo "$(color '33' '[WARN]') $*"; }
err()   { echo "$(color '31' '[ERR] ') $*"; }
mask()  { local s="${1:-}"; [[ -z "$s" ]] && echo "" || echo "${s:0:5}***"; }
die()   { err "$*"; if (( __SOURCED )); then return 1; else exit 1; fi; }

need_bin() { command -v "$1" >/dev/null 2>&1 || die "Missing dependency: $1"; }

k8s_get_secret_val() {
  # args: <ns> <secret> <key>
  local ns="$1" name="$2" key="$3"
  kubectl -n "$ns" get secret "$name" -o jsonpath="{.data.$key}" 2>/dev/null | base64 -d || true
}
k8s_secret_exists() { kubectl -n "$1" get secret "$2" >/dev/null 2>&1; }

# ── env load (optional .env/.env.local nearby) ─────────────────────────────
ENV_FILE="${ENV_FILE:-}"
if [[ -z "$ENV_FILE" ]]; then
  for f in ".env.local" ".env" "../.env" "../../.env"; do [[ -f "$f" ]] && ENV_FILE="$f" && break; done
fi
if [[ -n "$ENV_FILE" && -f "$ENV_FILE" ]]; then
  info "Loading env from $ENV_FILE"
  # shellcheck disable=SC2046
  export $(grep -E '^[A-Za-z_][A-Za-z0-9_]*=' "$ENV_FILE" | sed 's/#.*//')
fi

# ── defaults ───────────────────────────────────────────────────────────────
NS_SKILLSIER="${NS_SKILLSIER:-skillsier}"
NS_KEYCLOAK="${NS_KEYCLOAK:-keycloak}"
NS_KAFKA="${NS_KAFKA:-kafka}"

DB_SECRET_NAME="${DB_SECRET_NAME:-users-be-postgres}"
DB_SECRET_KEY="${DB_SECRET_KEY:-POSTGRES_PASSWORD}"

KAFKA_USERNAME="${KAFKA_USERNAME:-${KAFKA_USER:-}}"
KAFKA_USER_SECRET="${KAFKA_USER_SECRET:-${KAFKA_USERNAME:-}}"
KAFKA_PASSWORD_KEY="${KAFKA_PASSWORD_KEY:-password}"

KC_SECRET_NAME="${KC_SECRET_NAME:-keycloak-client-users-be}"
KC_ID_KEY="client-id"
KC_SECRET_KEY="client-secret"

USERS_BE_KEYCLOAK_URL="${USERS_BE_KEYCLOAK_URL:-}"
USERS_BE_KEYCLOAK_REALM="${USERS_BE_KEYCLOAK_REALM:-}"
KEYCLOAK_CLIENT_ID="${KEYCLOAK_CLIENT_ID:-}"
KEYCLOAK_CLIENT_SECRET="${KEYCLOAK_CLIENT_SECRET:-}"

need_bin kubectl
need_bin curl

[[ -n "$USERS_BE_KEYCLOAK_URL" ]] && USERS_BE_KEYCLOAK_URL="${USERS_BE_KEYCLOAK_URL%/}"

# ── Keycloak helpers ───────────────────────────────────────────────────────
get_keycloak_client_id()      { k8s_get_secret_val "$NS_KEYCLOAK" "$KC_SECRET_NAME" "$KC_ID_KEY"; }
get_keycloak_client_secret()  { k8s_get_secret_val "$NS_KEYCLOAK" "$KC_SECRET_NAME" "$KC_SECRET_KEY"; }

verify_keycloak_credentials() {
  # args: <client_id> <client_secret>
  local cid="$1" csec="$2"
  if (( NO_VERIFY )) || [[ -z "$USERS_BE_KEYCLOAK_URL" || -z "$USERS_BE_KEYCLOAK_REALM" ]]; then
    warn "Skipping Keycloak verification."
    return 0
  fi
  local token_url="${USERS_BE_KEYCLOAK_URL}/realms/${USERS_BE_KEYCLOAK_REALM}/protocol/openid-connect/token"
  info "Testing Keycloak client_credentials…"
  local code
  code=$(curl -ksS -o /dev/null -w "%{http_code}" \
    -X POST "$token_url" -H "Content-Type: application/x-www-form-urlencoded" \
    --data "grant_type=client_credentials&client_id=${cid}&client_secret=${csec}")
  [[ "$code" == "200" ]] && { ok "Credentials are valid"; return 0; }
  err "Keycloak validation failed (HTTP $code)"; return 1
}

create_or_update_kc_secret_from_env() {
  [[ -z "$KEYCLOAK_CLIENT_ID" || -z "$KEYCLOAK_CLIENT_SECRET" ]] && \
    die "KEYCLOAK_CLIENT_ID/KEYCLOAK_CLIENT_SECRET not set in env."
  info "Applying secret '$KC_SECRET_NAME' in ns '$NS_KEYCLOAK' from env…"
  kubectl -n "$NS_KEYCLOAK" create secret generic "$KC_SECRET_NAME" \
    --from-literal="${KC_ID_KEY}=${KEYCLOAK_CLIENT_ID}" \
    --from-literal="${KC_SECRET_KEY}=${KEYCLOAK_CLIENT_SECRET}" \
    --dry-run=client -o yaml | kubectl apply -f -
  ok "Secret applied."
}

reconcile_keycloak_client_secret() {
  echo
  info "Step 1: Checking Keycloak client credentials in Kubernetes…"
  if k8s_secret_exists "$NS_KEYCLOAK" "$KC_SECRET_NAME"; then
    info "Secret '$KC_SECRET_NAME' found."
    local cid csec
    cid="$(get_keycloak_client_id || true)"
    csec="$(get_keycloak_client_secret || true)"
    if [[ -n "$cid" && -n "$csec" ]]; then
      info "Verifying existing credentials…"
      if verify_keycloak_credentials "$cid" "$csec"; then
        export KEYCLOAK_CLIENT_ID="$cid" KEYCLOAK_CLIENT_SECRET="$csec"
        ok "Using credentials from K8s secret."
      else
        warn "Invalid creds in secret; recreating from env…"
        kubectl -n "$NS_KEYCLOAK" delete secret "$KC_SECRET_NAME" --ignore-not-found
        verify_keycloak_credentials "$KEYCLOAK_CLIENT_ID" "$KEYCLOAK_CLIENT_SECRET" || die "Env creds invalid."
        create_or_update_kc_secret_from_env
        export KEYCLOAK_CLIENT_ID="$(get_keycloak_client_id)"
        export KEYCLOAK_CLIENT_SECRET="$(get_keycloak_client_secret)"
        ok "Refreshed credentials."
      fi
    else
      warn "Secret exists but missing keys; recreating from env…"
      create_or_update_kc_secret_from_env
      export KEYCLOAK_CLIENT_ID="$(get_keycloak_client_id)"
      export KEYCLOAK_CLIENT_SECRET="$(get_keycloak_client_secret)"
      ok "Credentials retrieved."
    fi
  else
    warn "Secret '$KC_SECRET_NAME' not found."
    info "Creating from env:"
    echo "  KEYCLOAK_CLIENT_ID: ${KEYCLOAK_CLIENT_ID:-<unset>}"
    echo "  KEYCLOAK_CLIENT_SECRET: $(mask "$KEYCLOAK_CLIENT_SECRET")"
    verify_keycloak_credentials "$KEYCLOAK_CLIENT_ID" "$KEYCLOAK_CLIENT_SECRET" || die "Env creds invalid."
    create_or_update_kc_secret_from_env
    export KEYCLOAK_CLIENT_ID="$(get_keycloak_client_id)"
    export KEYCLOAK_CLIENT_SECRET="$(get_keycloak_client_secret)"
    ok "Credentials stored and loaded."
  fi
}

# ── other secrets ──────────────────────────────────────────────────────────
get_postgres_password() { k8s_get_secret_val "$NS_SKILLSIER" "$DB_SECRET_NAME" "$DB_SECRET_KEY"; }
get_kafka_password() {
  [[ -n "${KAFKA_PASSWORD:-}" ]] && { echo "$KAFKA_PASSWORD"; return 0; }
  if [[ -n "$KAFKA_USER_SECRET" ]]; then
    local v; v="$(k8s_get_secret_val "$NS_KAFKA" "$KAFKA_USER_SECRET" "$KAFKA_PASSWORD_KEY")"
    [[ -n "$v" ]] && { echo "$v"; return 0; }
  fi
  local v2; v2="$(k8s_get_secret_val "$NS_KAFKA" "dapr-pubsub" "password" 2>/dev/null || true)"
  [[ -n "$v2" ]] && { echo "$v2"; return 0; }
  echo ""
}

# ── main ───────────────────────────────────────────────────────────────────
need_bin kubectl; need_bin curl
reconcile_keycloak_client_secret

DB_PASSWORD="$(get_postgres_password || true)"
KAFKA_PASSWORD="$(get_kafka_password || true)"
export DB_PASSWORD POSTGRES_PASSWORD="$DB_PASSWORD" KAFKA_PASSWORD
export KEYCLOAK_CLIENT_ID KEYCLOAK_CLIENT_SECRET

if (( PRINT_ENV )); then
  # emit export lines (use with: source <(./scripts/get-secrets.sh --print-env))
  cat <<EOF
export KEYCLOAK_CLIENT_ID='${KEYCLOAK_CLIENT_ID}'
export KEYCLOAK_CLIENT_SECRET='${KEYCLOAK_CLIENT_SECRET}'
export DB_PASSWORD='${DB_PASSWORD}'
export POSTGRES_PASSWORD='${DB_PASSWORD}'
export KAFKA_PASSWORD='${KAFKA_PASSWORD}'
EOF
else
  echo
  echo "========================================="
  echo "✓ All secrets loaded successfully!"
  echo "-----------------------------------------"
  echo "  Keycloak Client ID:  ${KEYCLOAK_CLIENT_ID}"
  echo "  Keycloak Secret:     $(mask "${KEYCLOAK_CLIENT_SECRET}")"
  echo "  DB Password:         $(mask "${DB_PASSWORD}")"
  echo "  Kafka Password:      $(mask "${KAFKA_PASSWORD}")"
  echo "========================================="
fi

# In sourced mode, never exit the terminal
if (( __SOURCED )); then return 0; fi
