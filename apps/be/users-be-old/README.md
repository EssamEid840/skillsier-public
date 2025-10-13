# Users-BE Microservice

A clean architecture microservice for user management in the Skillsier platform. Built with Go, Gin, GORM, and event-driven architecture using Kafka and the Outbox pattern.

## Features

- 🏗️ **Clean Architecture** - Separation of concerns with clear layers
- 📦 **Event-Driven** - Uses Kafka for event streaming
- 📮 **Outbox Pattern** - Reliable event publishing with Watermill
- 🔐 **Keycloak Integration** - Automatic user creation from Keycloak events
- 🗄️ **PostgreSQL** - Relational database with GORM ORM
- 🚀 **Auto-Migration** - Database schema automatically updated
- 🔄 **Dual Mode** - Run locally or in Kubernetes
- 📊 **Health Checks** - Kubernetes-ready probes

## Architecture

```
users-be/
├── cmd/api/              # Application entry point
├── internal/
│   ├── domain/          # Business entities and repository interfaces
│   ├── application/     # Use cases and business logic
│   ├── infrastructure/  # External implementations (DB, Kafka, etc.)
│   ├── interfaces/      # HTTP handlers and routes
│   └── config/          # Configuration management
├── pkg/                 # Shared utilities
├── deployments/         # Kubernetes manifests
└── scripts/             # Helper scripts
```

## Prerequisites

- **Go 1.23+**
- **PostgreSQL 16** (running in K8s)
- **Kafka** (running in K8s with SASL_SSL)
- **Keycloak** (running in K8s) - Client name: `users-be`
- **kubectl** (configured to access your cluster)
- **Docker** (for containerization)

## Quick Start

### 1. Setup Keycloak Client

First, create a client named `users-be` in Keycloak. See [KEYCLOAK_SETUP.md](KEYCLOAK_SETUP.md) for detailed instructions.

**Quick steps:**
1. Go to Keycloak admin console → Clients → Create client
2. Client ID: `users-be`
3. Client authentication: **ON**
4. Service accounts roles: **ON**
5. Copy the client secret from Credentials tab

### 2. Setup Local Development

```bash
# Clone the repository
git clone <your-repo>
cd users-be

# Install dependencies
make deps

# Export Keycloak credentials (get these from Keycloak admin console)
export KEYCLOAK_CLIENT_ID=users-be
export KEYCLOAK_CLIENT_SECRET=your-secret-from-keycloak

# Setup local environment and create K8s secret
source scripts/get-secrets.sh
```

The script will:
- ✅ Verify your Keycloak credentials
- ✅ Create the K8s secret `keycloak-client-users-be`
- ✅ Fetch all other secrets (DB, Kafka)
- ✅ Export environment variables

### 3. Run Locally

```bash
# Run the application
make run

# Or build and run
make build-local
./bin/users-be
```

The service will:
- Connect to PostgreSQL in K8s (via NodePort 30432)
- Connect to Kafka in K8s (via NodePort 31691)
- Listen for Keycloak events
- Publish user events to Kafka
- Expose HTTP API on port 8080

### 3. Test the API

```bash
# Health check
curl http://localhost:8080/health

# Create a user (normally done via Keycloak events)
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "keycloak_id": "test-keycloak-id",
    "username": "testuser",
    "email": "test@example.com",
    "first_name": "Test",
    "last_name": "User"
  }'

# Get user by Keycloak ID
curl http://localhost:8080/api/v1/users/keycloak/test-keycloak-id

# List users
curl http://localhost:8080/api/v1/users?page=1&page_size=20

# Update user
curl -X PUT http://localhost:8080/api/v1/users/{user-id} \
  -H "Content-Type: application/json" \
  -d '{
    "profile_type": "freelancer",
    "profession": "Software Engineer",
    "hourly_rate": 50.00
  }'
```

## Configuration

Configuration is loaded from:
1. Environment variables (highest priority)
2. Configuration file (`config.yaml`)
3. Default values

### Environment Variables

```bash
# Database
export USERS_BE_DATABASE_HOST=173.212.218.251
export USERS_BE_DATABASE_PORT=30432
export DB_PASSWORD=your-password

# Kafka
export USERS_BE_KAFKA_BROKERS=173.212.218.251:31691
export KAFKA_PASSWORD=your-kafka-password

# Keycloak
# export USERS_BE_KEYCLOAK_URL=http://173.212.218.251:30080
export USERS_BE_KEYCLOAK_URL=https://keycloak.skillsier.com
export KEYCLOAK_CLIENT_ID=myapp
export KEYCLOAK_CLIENT_SECRET=your-secret
```

See `.env.example` for all available options.

## Database

### Auto-Migration

The service automatically creates and updates database tables on startup:

- `users` - User profiles with additional information
- `outbox_events` - Event outbox for reliable publishing

### Manual Operations

```bash
# Reset database (WARNING: Deletes all data!)
make db-reset
```

## Events

### Consuming Keycloak Events

The service listens to the `keycloak-events` topic and automatically:
- Creates users when they register in Keycloak
- Updates user information on profile changes
- Tracks login events

### Publishing User Events

User operations publish events to the `user-events` topic:
- `user.created` - New user registered
- `user.updated` - User profile updated
- `user.deleted` - User soft-deleted

Events use the Outbox pattern for reliable delivery.

## Deployment

### Build Docker Image

```bash
make docker-build
```

### Deploy to Kubernetes

```bash
# Update secrets first
kubectl apply -f deployments/k8s/secret.yaml

# Deploy the application
make k8s-deploy
```

### Check Deployment

```bash
# View status
make k8s-status

# View logs
make k8s-logs

# Restart deployment
make k8s-restart
```

## Development

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Run vet
make vet
```

### Testing

```bash
# Run unit tests
make test

# Run integration tests
make test-integration

# View coverage
go tool cover -html=coverage.out
```

### Hot Reload

Install Air for hot reloading:

```bash
make install-tools
make dev
```

## API Endpoints

### Health Checks
- `GET /health` - Overall health status
- `GET /ready` - Readiness probe
- `GET /live` - Liveness probe

### User Management
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - List users (paginated)
- `GET /api/v1/users/:id` - Get user by ID
- `GET /api/v1/users/keycloak/:keycloak_id` - Get user by Keycloak ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user (soft delete)

## Monitoring

### Metrics

The service exposes:
- HTTP request metrics (via middleware)
- Database connection pool metrics
- Kafka consumer lag (via Kafka metrics)
- Outbox processing metrics

### Logs

Structured logging with:
- Request/response logging
- Event processing logs
- Error tracking
- Performance metrics

## Troubleshooting

### Cannot connect to PostgreSQL

```bash
# Verify service is exposed
kubectl get svc users-be-postgres-external -n skillsier

# Test connection
psql -h 173.212.218.251 -p 30432 -U usersuser -d usersdb
```

### Cannot connect to Kafka

```bash
# Verify Kafka is running
kubectl get pods -n kafka

# Test connection
echo "test" | kcat -P -b 173.212.218.251:31691 -t smoke
```

### Events not being consumed

```bash
# Check consumer group
kafka-consumer-groups.sh --bootstrap-server 173.212.218.251:31691 \
  --group users-be-group --describe

# View application logs
make k8s-logs
```

## Contributing

1. Follow clean architecture principles
2. Write tests for new features
3. Update documentation
4. Run linters before committing

## License

Proprietary - Skillsier Platform