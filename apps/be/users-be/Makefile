# Makefile — Infra in Docker, app locally
# Usage: `make help`

# ---------- Config (override via env or `make VAR=value`) ----------
PROJECT_NAME ?= users-outbox
COMPOSE_FILE ?= docker-compose.yml
DC ?= docker compose -f $(COMPOSE_FILE)

INFRA_SERVICES ?= postgres redpanda kafka-ui

GO ?= go

# Database (matches docker-compose defaults)
DB_USER ?= user
DB_PASS ?= pass
DB_HOST ?= localhost          # talk to Postgres from host
DB_PORT ?= 5432
DB_NAME ?= users
DB_SSLMODE ?= disable
DB_DSN ?= postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Kafka (Redpanda) broker from host
KAFKA_BROKER ?= localhost:9092

# ---------- Phony ----------
.PHONY: help create-config \
        start-infra stop-infra down-infra logs-infra \
        db-shell db-pgcrypto \
        migrate-up-local migrate-down-1 migrate-steps migrate-version migrate-force \
        run-app-local run-consumer-local env \
        start build up

# ---------- Help ----------
help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1mTargets\033[0m\n"} /^[a-zA-Z0-9_.-]+:.*?##/ { printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
	@echo ""
	@echo "Key vars: DB_DSN=$(DB_DSN)  KAFKA_BROKER=$(KAFKA_BROKER)"

# ---------- Config ----------
create-config: ## Copy sample config to configs/.env.yml (idempotent)
	@[ -f configs/.env.yml ] || cp configs/.env.yml.dist configs/.env.yml
	@echo "✔ configs/.env.yml ready"

# ---------- Infra (Docker) ----------
start-infra: create-config ## Up Postgres, Redpanda, Kafka UI
	$(DC) up -d $(INFRA_SERVICES)

stop-infra: ## Stop Postgres, Redpanda, Kafka UI (containers kept)
	$(DC) stop $(INFRA_SERVICES)

down-infra: ## Down & remove volumes for all services (careful)
	$(DC) down -v

logs-infra: ## Follow logs for infra
	$(DC) logs -f $(INFRA_SERVICES)

# ---------- Database helpers ----------
db-shell: ## psql shell into Postgres
	$(DC) exec postgres psql -U $(DB_USER) -d $(DB_NAME)

db-pgcrypto: ## Ensure pgcrypto extension (gen_random_uuid)
	$(DC) exec postgres psql -U $(DB_USER) -d $(DB_NAME) -c 'CREATE EXTENSION IF NOT EXISTS "pgcrypto";'

# ---------- Migrations (run from local Go) ----------
migrate-up-local: ## Apply all up migrations using local Go migrator
	@echo "Using DB_DSN=$(DB_DSN)"
	DB_DSN='$(DB_DSN)' $(GO) run ./cmd/migrator -action up

migrate-down-1: ## Roll back one migration
	@echo "Using DB_DSN=$(DB_DSN)"
	DB_DSN='$(DB_DSN)' $(GO) run ./cmd/migrator -action down

migrate-steps: ## Run N steps (use: make migrate-steps N=-2)
	@echo "Using DB_DSN=$(DB_DSN)  steps=$(N)"
	DB_DSN='$(DB_DSN)' $(GO) run ./cmd/migrator -action steps -n $(N)

migrate-version: ## Show current migration version
	@echo "Using DB_DSN=$(DB_DSN)"
	DB_DSN='$(DB_DSN)' $(GO) run ./cmd/migrator -action version

migrate-force: ## Force version (use: make migrate-force V=20250504001817)
	@echo "Using DB_DSN=$(DB_DSN)  version=$(V)"
	DB_DSN='$(DB_DSN)' $(GO) run ./cmd/migrator -action force -version $(V)

# ---------- Run locally ----------
run-app-local: ## Run API locally against Docker infra
	@echo "KAFKA_BROKER=$(KAFKA_BROKER)  DB_DSN=$(DB_DSN)"
	KAFKA_BROKER='$(KAFKA_BROKER)' DATABASE_URL='$(DB_DSN)' $(GO) run ./cmd/app

run-consumer-local: ## Run consumer locally against Docker infra
	@echo "KAFKA_BROKER=$(KAFKA_BROKER)  DB_DSN=$(DB_DSN)"
	KAFKA_BROKER='$(KAFKA_BROKER)' DATABASE_URL='$(DB_DSN)' $(GO) run ./cmd/consumer

env: ## Print key env exports to use in your shell
	@echo "export DATABASE_URL='$(DB_DSN)'"
	@echo "export KAFKA_BROKER='$(KAFKA_BROKER)'"

# ---------- Back-compat targets (deprecated) ----------
start: ## (Deprecated) Old target; use start-infra
	@$(MAKE) start-infra

build: ## (Deprecated) Old target; prefer start-infra then run-app-local
	$(DC) up -d --build

up: ## (Deprecated) Old target; prefer start-infra
	$(DC) up -d

# ----------------- Local DAPR ---------------------------
DAPR_RES ?= ./deploy/local-dapr

dapr-run-app-local:
	APP_OUTBOX__MODE=cdc APP_DAPR__ENABLED=true \
	dapr run --app-id users-outbox \
	  --app-port 8080 \
	  --dapr-http-port 3501 \
	  --resources-path $(DAPR_RES) \
	  -- go run ./cmd/app

dapr-run-consumer-local:
	APP_DAPR__ENABLED=true \
	APP_DAPR__CONSUMER__APP_PORT=8082 \
	dapr run --app-id users-outbox-consumer \
	  --app-port 8082 \
	  --dapr-http-port 3502 \
	  --resources-path $(DAPR_RES) \
	  -- go run ./cmd/consumer

# ----------------- Remote Development ---------------------------
PF_NS ?= kafka
PF_CTX ?= skillsier-prod

.PHONY: remote-dev-pf-kafka
remote-dev-pf-kafka:
	kubectl config use-context $(PF_CTX)
	@echo "Port-forwarding prod Kafka 9094 → localhost:19094..."
	kubectl -n $(PF_NS) port-forward svc/kafka-cluster-kafka-bootstrap 19094:9094

.PHONY: remote-dev-api
remote-dev-api:
	CONFIG_ENV=prod-local \
	APP_DAPR__ENABLED=true \
	APP_DAPR__PUBSUB=kafka-pubsub \
	APP_KAFKA__ENABLED=false \
	dapr run \
	  --app-id users-api \
	  --app-port 8080 \
	  --dapr-http-port 3500 \
	  --components-path ./remote-dev \
	  -- \
	  go run ./cmd/app

.PHONY: remote-dev-consumer
remote-dev-consumer:
	CONFIG_ENV=prod-local \
	APP_DAPR__ENABLED=true \
	APP_DAPR__PUBSUB=kafka-pubsub \
	APP_DAPR__HTTP_PORT=3501 \
	APP_DAPR__CONSUMER__APP_PORT=8082 \
	APP_KAFKA__ENABLED=false \
	dapr run \
	  --app-id users-consumer \
	  --app-port 8082 \
	  --dapr-http-port 3501 \
	  --components-path ./remote-dev \
	  -- \
	  go run ./cmd/consumer
