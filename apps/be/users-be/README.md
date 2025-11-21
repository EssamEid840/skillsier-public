# Users-BE Microservice Template

A production-ready Golang/Gin microservice template for the Skillsier platform with automatic database migrations, outbox pattern, and full Kubernetes support.

## ✨ Features

- ✅ **Auto-Migrations** - Automatic database schema creation/updates at startup
- ✅ **Outbox Pattern** - Reliable event publishing with Kafka
- ✅ **Clean Architecture** - Domain-driven design with clear separation of concerns
- ✅ **Event-Driven** - Full event sourcing with protobuf schemas
- ✅ **Production-Ready** - Health checks, graceful shutdown, observability
- ✅ **K8s Native** - Complete Kubernetes manifests with Kustomize
- ✅ **SASL_SSL** - Secure Kafka connection with SCRAM authentication
- ✅ **Platform Integration** - Uses pkg/auth, platform-shared, contracts/events

## 🚀 Quick Start

### Prerequisites

- Go 1.23+
- PostgreSQL 16 (K8s)
- Kafka (K8s with SASL_SSL)
- Keycloak (K8s)
- kubectl configured

### 1. Instantiate Template

```bash
cd scripts
./instantiate-template jobs
# Follow prompts to configure service
```

### 1.1 Prepare PostgreSQL Deployment
- Change the NodePort
- Add this Port to the Global en => Prod of Scripts of Infra
- Run Prequestits to open this Port
- Run
```bash
make k8s-db.apply
```
- Connect with PGAdmin to make Sure Database is working

```bash
cd scripts
./instantiate-template jobs
# Follow prompts to configure service
```

### 2. Setup Dependencies

```bash
cd ../jobs-be
make deps
```

### 3.1 Prepare Environment
- In apps/be/Microservice/internal/config/config.go 
  - For the Datbase
    - Change the Host to Server IP
    - Change the Port to the Database Port
- We suppose to make all of that automaticlly later

### 3.2 Configure Environment

```bash
cp .env.example .env
# Edit .env with your configuration
# Or fetch from K8s:
source scripts/get-secrets.sh
```

### 4. Run Locally

```bash
make run-local
```

The service will:
- Connect to K8s PostgreSQL (NodePort 30432)
- Connect to K8s Kafka (NodePort 31691)
- Auto-create database tables
- Start outbox dispatcher
- Expose HTTP API on port 8080

### 5. Test

```bash
# Health check
curl http://localhost:8080/health

# Create entity
curl -X POST http://localhost:8080/v1/initial-entities \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Entity",
    "description": "Testing",
    "owner_id": "123e4567-e89b-12d3-a456-426614174000"
  }'
```

## 📚 Documentation

- **STRUCTURE.md** - Complete folder structure with comments
- **internal/config/docs/CONFIGURATION.md** - Environment variables
- **migrations/README.md** - Database migrations guide

## 🏗️ Architecture

```
template-users-be/
├── cmd/api/                    # Application entrypoint
├── internal/
│   ├── config/                 # Configuration management
│   ├── domain/                 # Domain entities & business logic
│   ├── application/            # Use cases & services
│   ├── infrastructure/         # DB, Kafka, outbox
│   ├── interfaces/http/        # HTTP handlers & routes
│   └── ioc/                    # Dependency injection
├── contracts/events/           # Protobuf event schemas
├── k8s/                        # Kubernetes manifests
├── tests/                      # Unit & integration tests
└── scripts/                    # Helper scripts
```

## 🔄 Adding a New Entity

1. **Create domain entity**:
   ```go
   // internal/domain/customer/entity.go
   type Customer struct {
       ID   uuid.UUID
       Name string
       // ...
   }
   ```

2. **Register for auto-migration**:
   ```go
   // internal/infrastructure/persistence/postgres/migrations.go
   func init() {
       Register(
           &initial_entity.InitialEntity{},
           &customer.Customer{},  // Add this line
       )
   }
   ```

3. **Restart service** - table created automatically!

4. **Implement repository, service, handlers** - follow existing patterns

See STRUCTURE.md for complete guide.

## 🧪 Testing

```bash
# Unit tests (no K8s required)
make test

# Integration tests (requires K8s)
INTEGRATION_TEST=true make test-integration

# Coverage report
make test-coverage
```

## 🐳 Docker

```bash
# Build image
make docker-build

# Push to registry
make docker-push
```

## ☸️ Kubernetes

```bash
# Deploy to K8s
make k8s-deploy

# View logs
make k8s-logs

# Check status
make k8s-status

# Delete deployment
make k8s-delete
```

## 🔧 Development

```bash
# Run locally against K8s
make run-local

# Format code
make fmt

# Run linter
make lint

# Generate protobuf
make generate-proto

# Port-forward PostgreSQL
make port-forward-postgres

# Port-forward Kafka
make port-forward-kafka
```

## 📊 Auto-Migration

**Development** (default):
- `AUTO_MIGRATE=true` - Migrations run on startup
- New entities automatically create tables

**Production** (default):
- `AUTO_MIGRATE=false` - Migrations disabled
- Set `ALLOW_IN_PRODUCTION=true` to override

**How it works**:
1. Entities registered in `migrations.go`
2. GORM AutoMigrate runs on startup
3. Schema version tracked in `schema_versions` table
4. Safe guards prevent destructive changes in production

## 🔐 Security

- SASL_SSL for Kafka (SCRAM-SHA-512)
- TLS for PostgreSQL (production)
- Keycloak JWT authentication ready
- Non-root Docker user
- Health checks for K8s probes

## 📈 Observability

- Structured logging (zerolog ready)
- Health endpoints (/health, /ready, /live)
- Request ID tracking
- Database connection stats
- Outbox statistics API

## 🎯 Production Checklist

- [ ] Configure secrets in K8s
- [ ] Set `AUTO_MIGRATE=false` in production
- [ ] Enable TLS for PostgreSQL
- [ ] Configure resource limits
- [ ] Setup monitoring/alerting
- [ ] Configure backup strategy
- [ ] Review security settings

## 🤝 Contributing

This is a template. Customize for your service needs!

## 📝 License

Internal Skillsier Platform Template

---

**Generated from**: template-service-be v1.0.0  
**Compatible with**: Skillsier Platform Phase 1+  
**Last updated**: 2024-11-15