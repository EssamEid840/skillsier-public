# template-<service>-be - Complete Folder Structure

## Overview
This is a production-ready Golang/Gin microservice template for the Skillsier platform. It implements clean architecture, automatic database migrations, outbox pattern, and full integration with platform-shared modules.

---

## Complete Folder Tree with Comments

```
template-<service>-be/
│
├── cmd/
│   └── api/
│       └── main.go                                    # Application entrypoint - bootstraps DI container, registers routes, starts Gin server
│
├── internal/
│   │
│   # =============================
│   # 🔧 CONFIGURATION (LOAD FIRST)
│   # =============================
│   ├── config/
│   │   ├── schema.go                              # Typed Config struct (App, Server, Database, Kafka, Keycloak, Dapr)
│   │   ├── loader.go                              # Viper-based loader (env vars → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md                   # Environment variable documentation with examples
│   │
│   # =============================
│   # 🏛️ DOMAIN LAYER (DDD)
│   # =============================
│   ├── domain/
│   │   └── initial_entity/
│   │       ├── entity.go                          # InitialEntity aggregate root (ID, Name, Status, Timestamps) with validation
│   │       ├── enums.go                           # Status enum (Active, Inactive, Archived)
│   │       ├── errors.go                          # Domain-specific errors (EntityNotFound, InvalidStatus, etc.)
│   │       ├── repository.go                      # Repository interface (CRUD methods, CreateWithOutbox)
│   │       └── events.go                          # Domain events (InitialEntityCreated, InitialEntityUpdated, etc.)
│   │
│   # =============================
│   # 🧠 APPLICATION LAYER (USE CASES)
│   # =============================
│   ├── application/
│   │   └── initial_entity/
│   │       ├── service.go                         # Business logic orchestration (Create, Update, Delete, Get, List)
│   │       ├── commands.go                        # Command handlers (CreateInitialEntityCommand, UpdateInitialEntityCommand)
│   │       ├── queries.go                         # Query handlers (GetInitialEntityQuery, ListInitialEntitiesQuery)
│   │       ├── dto.go                             # Data Transfer Objects (CreateDTO, UpdateDTO, ResponseDTO)
│   │       ├── mapper.go                          # Entity ↔ DTO mapping functions
│   │       └── validators.go                      # Business validation rules (name length, status transitions)
│   │
│   # =============================
│   # 🔌 INFRASTRUCTURE LAYER
│   # =============================
│   ├── infrastructure/
│   │   │
│   │   # =========================
│   │   # 🗄️ PERSISTENCE (POSTGRES)
│   │   # =========================
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── connection.go              # Database connection setup (DSN, pooling, timeouts)
│   │   │       ├── transaction.go             # Transaction helpers (WithTx, rollback handling)
│   │   │       ├── migrations.go              # **AUTO-MIGRATION ENGINE** (entity registration, startup migration)
│   │   │       ├── registry.go                # **ENTITY REGISTRY** (Register models for auto-migration)
│   │   │       ├── safety.go                  # Production safety guards (ENV checks, destructive change protection)
│   │   │       ├── version.go                 # Schema version tracking table
│   │   │       ├── initial_entity_repository.go  # InitialEntity repository implementation (CRUD + CreateWithOutbox)
│   │   │       └── outbox_store.go            # Outbox table access within same DB transaction
│   │   │
│   │   # =========================
│   │   # 📨 MESSAGING (KAFKA + OUTBOX)
│   │   # =========================
│   │   ├── messaging/
│   │   │   ├── outbox/
│   │   │   │   ├── dispatcher.go              # Background outbox processor (polls pending events, publishes to Kafka)
│   │   │   │   └── scheduler.go               # Retry logic with exponential backoff and DLQ handling
│   │   │   │
│   │   │   └── kafka/
│   │   │       ├── producer.go                # Kafka producer (SASL_SSL with SCRAM-SHA-512)
│   │   │       ├── consumer.go                # Kafka consumer skeleton (for consuming events from other services)
│   │   │       └── scram.go                   # SCRAM authentication helper
│   │   │
│   │   # =========================
│   │   # 🔐 EXTERNAL CLIENTS
│   │   # =========================
│   │   └── clients/
│   │       └── keycloak_client.go             # Keycloak integration (if needed for service-specific operations)
│   │
│   # =============================
│   # 🌐 INTERFACES LAYER (HTTP)
│   # =============================
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/
│   │           ├── handlers/
│   │           │   ├── initial_entity_handler.go  # HTTP handlers for InitialEntity endpoints (Create, Update, Get, List, Delete)
│   │           │   └── health_handler.go          # Health check handlers (/health, /ready, /live)
│   │           │
│   │           └── routes/
│   │               ├── initial_entity_routes.go    # Route registration for InitialEntity (POST, GET, PUT, DELETE)
│   │               └── health_routes.go            # Health check route registration
│   │
│   # =============================
│   # 🧩 DEPENDENCY INJECTION
│   # =============================
│   └── ioc/
│       ├── container.go                           # DI container builder (creates all dependencies: DB, repos, services, handlers)
│       └── wiring.go                              # Wiring logic (connects repos → services → handlers)
│
├── contracts/
│   └── events/
│       └── initialentity/
│           └── v1/
│               └── initial_entity_events.proto    # Protobuf event schema (InitialEntityCreated, InitialEntityUpdated, etc.)
│
├── migrations/                                    # Optional deterministic SQL migrations (for teams preferring manual migrations)
│   ├── 000001_create_initial_entities_table.sql  # Example SQL migration
│   ├── 000002_create_outbox_table.sql            # Outbox table creation
│   └── README.md                                  # Migration usage notes (optional, primary method is auto-migrate)
│
├── k8s/
│   ├── db/
│   │   ├── postgres-secret.yaml
│   │   ├── postgres-statefulset.yaml
│   │   ├── postgres-service.yaml
│   │   └── postgres-nodeport.yaml
│   │   └── postgres-networkpolicy.yaml
│   ├── base/
│   │   ├── deployment.yaml                        # Kubernetes Deployment with placeholders (<service>, <docker_image>, etc.)
│   │   ├── service.yaml                           # Kubernetes Service (ClusterIP)
│   │   ├── configmap.yaml                         # ConfigMap with non-sensitive configuration
│   │   ├── secret.yaml                            # Secret placeholders (DB password, Kafka password, etc.)
│   │   └── kustomization.yaml                     # Kustomize base configuration
│   │
│   └── overlays/
│       └── local/
│           ├── kustomization.yaml                 # Local development overlay (NodePort services, debug settings)
│           └── deployment-patch.yaml              # Local-specific patches (port-forwards, env overrides)
│
├── scripts/
│   ├── instantiate-template                       # Token replacement script - converts template to concrete service
│   │                                             # Usage: ./scripts/instantiate-template <service-name>
│   └── run-local.sh                               # Helper script to run local binary against K8s infra (port-forwards)
│
├── tests/
│   ├── unit/
│   │   └── initial_entity_test.go                 # Unit tests with mocks (no K8s required)
│   │
│   └── integration/
│       └── initial_entity_integration_test.go     # Integration tests (guarded by INTEGRATION_TEST=true env var)
│
├── Makefile                                       # Build, test, run, deploy targets with inline comments
├── Dockerfile                                     # Multi-stage Docker build
├── .dockerignore                                  # Docker ignore patterns
├── .env.example                                   # Template environment variables with examples
├── go.mod                                         # Go module definition
├── go.sum                                         # Go module checksums (generated)
└── STRUCTURE.md                                   # This file - complete folder tree with comments

```

---

## How to Add a New Entity (Developer Guide)

When you instantiate this template and later need to add a new domain entity (e.g., `Customer`, `Order`, etc.), follow these steps:

### 1. **Create Domain Files**
```
internal/domain/<entity_name>/
├── entity.go          # Domain model with validation
├── enums.go           # Status enums or other domain enums
├── errors.go          # Domain-specific errors
├── repository.go      # Repository interface
└── events.go          # Domain events
```

### 2. **Register Entity for Auto-Migration**
In `internal/infrastructure/persistence/postgres/migrations.go`:
```go
// Add your entity to the registry in init() function
func init() {
    migrations.Register(&initial_entity.InitialEntity{})
    migrations.Register(&your_entity.YourEntity{})  // <- ADD THIS LINE
}
```
**That's it!** The next time you restart the service (or run `make run-local`), the auto-migration will detect the new entity and create its table automatically.

### 3. **Implement Repository**
Create `internal/infrastructure/persistence/postgres/<entity_name>_repository.go`:
- Implement all methods from the repository interface
- Include `CreateWithOutbox(ctx, tx, entity, eventType, payload, topic)` method

### 4. **Create Application Layer**
```
internal/application/<entity_name>/
├── service.go         # Business logic
├── commands.go        # Command handlers
├── queries.go         # Query handlers
├── dto.go             # Data Transfer Objects
├── mapper.go          # Entity ↔ DTO mapping
└── validators.go      # Business validation
```

### 5. **Create HTTP Interface**
```
internal/interfaces/http/v1/handlers/<entity_name>_handler.go   # HTTP handlers
internal/interfaces/http/v1/routes/<entity_name>_routes.go       # Route registration
```

### 6. **Wire in DI Container**
In `internal/ioc/container.go`:
```go
// Add repository
yourRepo := postgres.NewYourEntityRepository(container.DB)

// Add service
yourService := your_entity.NewService(yourRepo, container.OutboxPublisher)

// Add handler
yourHandler := handlers.NewYourEntityHandler(yourService)

// Register routes
your_entity_routes.RegisterRoutes(container.Router, yourHandler, container.Authorizer)
```

### 7. **Add Event Schema**
Create `contracts/events/<entity_name>/v1/<entity_name>_events.proto`:
```protobuf
syntax = "proto3";
package <entity_name>.v1;

message YourEntityCreated {
  string event_id = 1;
  google.protobuf.Timestamp event_timestamp = 2;
  string your_entity_id = 3;
  // ... other fields
}
```

### 8. **Write Tests**
- Unit tests: `tests/unit/<entity_name>_test.go`
- Integration tests: `tests/integration/<entity_name>_integration_test.go`

### 9. **Update Kubernetes Manifests (if needed)**
If your new entity requires additional configuration (DB name, topic name), add to:
- `k8s/base/configmap.yaml` - for non-sensitive config
- `k8s/base/secret.yaml` - for sensitive config

---

## Verification Checklist

After adding a new entity, verify the following:

- [ ] **Build succeeds**: `go build ./...` completes without errors
- [ ] **Auto-migration works**: Start service with `AUTO_MIGRATE=true`, check logs for new table creation
- [ ] **Unit tests pass**: `make test` runs successfully
- [ ] **HTTP endpoints work**: `POST /v1/<entities>` returns 201 and creates entity + outbox row
- [ ] **Outbox processes**: Dispatcher picks up pending outbox event and publishes to Kafka
- [ ] **Integration tests pass**: `INTEGRATION_TEST=true make test-integration` succeeds
- [ ] **K8s deployment works**: `make k8s-deploy` applies manifests successfully
- [ ] **Health checks respond**: `GET /health`, `GET /ready`, `GET /live` return 200
- [ ] **No hard-coded values**: All service names use placeholders (check `grep -r "template-" .`)
- [ ] **Event schema valid**: `buf lint` and `buf generate` succeed
- [ ] **Production safety**: With `ENV=production` and `AUTO_MIGRATE=false`, migrations don't run
- [ ] **Idempotency works**: Duplicate POST requests with same `Idempotency-Key` header return cached response

---

## Token Placeholders

The following placeholders are used throughout the template and must be replaced when instantiating:

| Placeholder | Description | Example |
|-------------|-------------|---------|
| `<service>` | Service name (kebab-case) | `jobs`, `proposals`, `contracts` |
| `<SERVICE>` | Service name (UPPER_CASE) | `JOBS`, `PROPOSALS`, `CONTRACTS` |
| `<Service>` | Service name (PascalCase) | `Jobs`, `Proposals`, `Contracts` |
| `<module>` | Go module path | `skillsier.dev/apps/be/jobs-be` |
| `<db_name>` | PostgreSQL database name | `jobsdb`, `proposalsdb` |
| `<db_user>` | PostgreSQL username | `jobsuser`, `proposalsuser` |
| `<postgres_dsn>` | PostgreSQL connection string | `postgres://user:pass@host:5432/dbname` |
| `<kafka_brokers>` | Kafka broker addresses | `173.212.218.251:31691` |
| `<topic_prefix>` | Kafka topic prefix | `jobs`, `proposals` |
| `<service_port>` | HTTP service port | `8080`, `8081` |
| `<docker_image>` | Docker image name | `skillsier/jobs-be:latest` |
| `<k8s_namespace>` | Kubernetes namespace | `skillsier` |
| `<keycloak_realm>` | Keycloak realm name | `skillsier` |
| `<keycloak_client_id>` | Keycloak client ID | `jobs-be`, `proposals-be` |
| `<idempotency_header>` | Idempotency header name | `Idempotency-Key` |

---

## Missing Platform-Shared APIs

During the scan, the following APIs were identified as missing from platform-shared and have been implemented in this template. **These should be extracted to platform-shared:**

### 1. **Auto-Migration System** (CRITICAL - NEW)
**Location to add**: `platform-shared/migrations/`

**Files to create**:
```
platform-shared/migrations/
├── registry.go          # Entity registration system
├── migrator.go          # Auto-migration engine
├── safety.go            # Production guards
└── version.go           # Schema version tracking
```

**API Signatures**:
```go
package migrations

// Registry manages entity models for auto-migration
type Registry interface {
    Register(models ...interface{})
    GetModels() []interface{}
}

// Migrator performs automatic database migrations
type Migrator interface {
    AutoMigrate(db *gorm.DB, config MigrationConfig) error
    GetVersion(db *gorm.DB) (int, error)
    SetVersion(db *gorm.DB, version int) error
}

// MigrationConfig controls migration behavior
type MigrationConfig struct {
    Enabled                bool   // Enable/disable auto-migration
    AllowInProduction      bool   // Allow migrations in production
    AllowDestructive       bool   // Allow destructive changes
    Environment            string // development, staging, production
    BackupBeforeChange     bool   // Create backup before migrations
}

// Functions
func NewRegistry() Registry
func NewMigrator() Migrator
func Register(models ...interface{})        // Global registration
func AutoMigrate(db *gorm.DB, config MigrationConfig) error
```

### 2. **Outbox Store Integration** (ENHANCEMENT)
**Location to add**: `platform-shared/outbox/postgres/store.go`

**API Additions**:
```go
package postgres

// CreateWithTx creates an outbox event within an existing transaction
func (r *Repository) CreateWithTx(tx *gorm.DB, event *outbox.Event) error

// CreateBatchWithTx creates multiple outbox events in one transaction
func (r *Repository) CreateBatchWithTx(tx *gorm.DB, events []*outbox.Event) error
```

### 3. **Transaction Helpers** (ENHANCEMENT)
**Location to add**: `platform-shared/persistence/postgres/transaction.go`

**API Additions**:
```go
package postgres

// WithTx executes a function within a database transaction
func WithTx(db *gorm.DB, fn func(*gorm.DB) error) error

// WithTxContext executes a function within a transaction with context
func WithTxContext(ctx context.Context, db *gorm.DB, fn func(context.Context, *gorm.DB) error) error

// SavePoint creates a savepoint within a transaction
func SavePoint(tx *gorm.DB, name string) error

// RollbackTo rolls back to a savepoint
func RollbackTo(tx *gorm.DB, name string) error
```

### 4. **Kafka SCRAM Authentication** (NEW)
**Location to add**: `platform-shared/kafka/scram.go`

**API**:
```go
package kafka

import "crypto/sha256"

// SCRAM256HashGenerator generates SCRAM-SHA-256 hashes
type SCRAM256HashGenerator struct{}

func (s *SCRAM256HashGenerator) Hash(username, password, nonce string) []byte

// NewSCRAMClient creates a new SCRAM client for Kafka
func NewSCRAMClient(mechanism string) sarama.SCRAMClient
```

---

## Notes

- **Primary migration method**: Automatic at startup (GORM AutoMigrate)
- **Optional SQL migrations**: Available in `migrations/` folder for teams preferring manual control
- **Production safety**: Auto-migrations disabled by default in production (set `AUTO_MIGRATE=true` to override)
- **Idempotency**: Built-in using platform-shared/idempotency middleware (header: `Idempotency-Key`)
- **Observability**: Integrated with platform-shared/logging (structured logs with zerolog)
- **Auth**: Ready for pkg/auth integration (uncomment in handlers and DI container)

---

## Quick Start Commands

```bash
# Install dependencies
make deps

# Run locally against K8s infra
make run-local

# Run unit tests
make test

# Run integration tests (requires K8s)
INTEGRATION_TEST=true make test-integration

# Build Docker image
make docker-build

# Deploy to K8s
make k8s-deploy

# View logs
make k8s-logs

# Instantiate template for a new service
./scripts/instantiate-template jobs
```

---

**Generated**: 2024-11-14  
**Template Version**: 1.0.0  
**Compatible with**: Skillsier Platform Phase 1+