# Configuration Documentation

This document describes all environment variables used by the users-be microservice.

## Configuration Priority

Configuration is loaded in the following order (highest priority first):

1. **Environment Variables** - Prefixed with `USERS_BE_`
2. **Config File** - `config.yaml` (optional)
3. **Default Values** - Built into the application

## Environment Variables

### Application Settings

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `USERS_BE_APP_NAME` | string | `users-be` | Application name |
| `USERS_BE_APP_ENVIRONMENT` | string | `development` | Environment: development, staging, production |
| `USERS_BE_APP_VERSION` | string | `1.0.0` | Application version |
| `USERS_BE_APP_LOG_LEVEL` | string | `info` | Log level: debug, info, warn, error |

### HTTP Server Settings

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `USERS_BE_SERVER_PORT` | string | `8080` | HTTP server port |
| `USERS_BE_SERVER_READ_TIMEOUT` | duration | `30s` | Request read timeout |
| `USERS_BE_SERVER_WRITE_TIMEOUT` | duration | `30s` | Response write timeout |
| `USERS_BE_SERVER_IDLE_TIMEOUT` | duration | `120s` | Keep-alive idle timeout |
| `USERS_BE_SERVER_SHUTDOWN_TIMEOUT` | duration | `10s` | Graceful shutdown timeout |
| `USERS_BE_SERVER_ENABLE_CORS` | bool | `true` | Enable CORS middleware |

### Database Settings

| Variable | Type | Default | Description | Required |
|----------|------|---------|-------------|----------|
| `USERS_BE_DATABASE_HOST` | string | `localhost` | PostgreSQL host | ✅ |
| `USERS_BE_DATABASE_PORT` | int | `5432` | PostgreSQL port | ✅ |
| `USERS_BE_DATABASE_USER` | string | `usersuser` | PostgreSQL username | ✅ |
| `USERS_BE_DATABASE_PASSWORD` | string | - | PostgreSQL password | ✅ |
| `USERS_BE_DATABASE_DATABASE` | string | `usersdb` | Database name | ✅ |
| `USERS_BE_DATABASE_SSL_MODE` | string | `disable` | SSL mode: disable, require, verify-ca, verify-full | |
| `USERS_BE_DATABASE_MAX_OPEN_CONNS` | int | `25` | Max open connections | |
| `USERS_BE_DATABASE_MAX_IDLE_CONNS` | int | `5` | Max idle connections | |
| `USERS_BE_DATABASE_CONN_MAX_LIFETIME` | duration | `30m` | Connection max lifetime | |
| `USERS_BE_DATABASE_CONN_MAX_IDLE_TIME` | duration | `10m` | Connection max idle time | |

### Database Migration Settings

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `USERS_BE_DATABASE_AUTO_MIGRATE` | bool | `true` | Enable automatic migrations on startup |
| `USERS_BE_DATABASE_MIGRATION_ENABLED` | bool | `true` | Enable migration system |
| `USERS_BE_DATABASE_MIGRATION_ALLOW_IN_PRODUCTION` | bool | `false` | Allow auto-migrations in production |
| `USERS_BE_DATABASE_MIGRATION_ALLOW_DESTRUCTIVE` | bool | `false` | Allow destructive schema changes |
| `USERS_BE_DATABASE_MIGRATION_BACKUP_BEFORE_CHANGE` | bool | `true` | Backup before destructive changes |
| `USERS_BE_DATABASE_MIGRATION_LOG_MIGRATION_SUMMARY` | bool | `true` | Log migration summary |

**⚠️ Production Safety**: In production, auto-migrations are disabled by default. Set `ALLOW_IN_PRODUCTION=true` to override.

### Kafka Settings

| Variable | Type | Default | Description | Required |
|----------|------|---------|-------------|----------|
| `USERS_BE_KAFKA_BROKERS` | []string | `localhost:9092` | Kafka broker addresses (comma-separated) | ✅ |
| `USERS_BE_KAFKA_SECURITY_PROTOCOL` | string | `SASL_SSL` | Security protocol: PLAINTEXT, SASL_SSL | ✅ |
| `USERS_BE_KAFKA_SASL_MECHANISM` | string | `SCRAM-SHA-512` | SASL mechanism | ✅ |
| `USERS_BE_KAFKA_SASL_USERNAME` | string | `admin-user` | SASL username | ✅ |
| `USERS_BE_KAFKA_SASL_PASSWORD` | string | - | SASL password | ✅ |
| `USERS_BE_KAFKA_TLS_SKIP_VERIFY` | bool | `false` | Skip TLS verification (dev only) | |
| `USERS_BE_KAFKA_CONSUMER_GROUP` | string | `users-be-group` | Consumer group ID | |
| `USERS_BE_KAFKA_TOPIC_PREFIX` | string | `users` | Topic prefix for published events | |

#### Kafka Producer Settings

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `USERS_BE_KAFKA_PRODUCER_REQUIRED_ACKS` | int | `-1` | Required acks: 0 (none), 1 (leader), -1 (all) |
| `USERS_BE_KAFKA_PRODUCER_MAX_MESSAGE_BYTES` | int | `1000000` | Max message size (1MB) |
| `USERS_BE_KAFKA_PRODUCER_COMPRESSION` | string | `snappy` | Compression: none, gzip, snappy, lz4, zstd |
| `USERS_BE_KAFKA_PRODUCER_IDEMPOTENT_WRITES` | bool | `true` | Enable idempotent writes |
| `USERS_BE_KAFKA_PRODUCER_MAX_RETRIES` | int | `5` | Max produce retries |
| `USERS_BE_KAFKA_PRODUCER_RETRY_BACKOFF` | duration | `100ms` | Retry backoff duration |

#### Kafka Consumer Settings

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `USERS_BE_KAFKA_CONSUMER_ENABLE_AUTO_COMMIT` | bool | `false` | Enable auto-commit |
| `USERS_BE_KAFKA_CONSUMER_AUTO_COMMIT_INTERVAL` | duration | `1s` | Auto-commit interval |
| `USERS_BE_KAFKA_CONSUMER_SESSION_TIMEOUT` | duration | `10s` | Session timeout |
| `USERS_BE_KAFKA_CONSUMER_FETCH_MIN_BYTES` | int | `1` | Min fetch bytes |
| `USERS_BE_KAFKA_CONSUMER_FETCH_MAX_WAIT` | duration | `500ms` | Max fetch wait time |
| `USERS_BE_KAFKA_CONSUMER_MAX_PROCESSING_TIME` | duration | `30s` | Max message processing time |

### Keycloak Settings

| Variable | Type | Default | Description | Required |
|----------|------|---------|-------------|----------|
| `USERS_BE_KEYCLOAK_URL` | string | `http://localhost:8080` | Keycloak base URL | ✅ |
| `USERS_BE_KEYCLOAK_REALM` | string | `skillsier` | Keycloak realm | ✅ |
| `USERS_BE_KEYCLOAK_CLIENT_ID` | string | `users-be` | Client ID | ✅ |
| `USERS_BE_KEYCLOAK_CLIENT_SECRET` | string | - | Client secret | |
| `USERS_BE_KEYCLOAK_TIMEOUT` | duration | `10s` | HTTP client timeout | |
| `USERS_BE_KEYCLOAK_JWKS_CACHE_TTL` | duration | `24h` | JWKS cache TTL | |

### Dapr Settings (Optional)

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `USERS_BE_DAPR_ENABLED` | bool | `false` | Enable Dapr integration |
| `USERS_BE_DAPR_HTTP_PORT` | int | `3500` | Dapr HTTP port |
| `USERS_BE_DAPR_GRPC_PORT` | int | `50001` | Dapr gRPC port |
| `USERS_BE_DAPR_PUBSUB_NAME` | string | `kafka-pubsub` | PubSub component name |
| `USERS_BE_DAPR_STATE_STORE_NAME` | string | `statestore` | State store component name |

### Outbox Pattern Settings

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `USERS_BE_OUTBOX_ENABLED` | bool | `true` | Enable outbox dispatcher |
| `USERS_BE_OUTBOX_POLL_INTERVAL` | duration | `5s` | Poll interval for pending events |
| `USERS_BE_OUTBOX_BATCH_SIZE` | int | `100` | Max events per batch |
| `USERS_BE_OUTBOX_MAX_RETRIES` | int | `10` | Max retry attempts |
| `USERS_BE_OUTBOX_RETRY_BACKOFF_BASE` | duration | `1s` | Base backoff duration |
| `USERS_BE_OUTBOX_RETRY_BACKOFF_MAX` | duration | `5m` | Max backoff duration |
| `USERS_BE_OUTBOX_CLEANUP_ENABLED` | bool | `true` | Enable cleanup of old events |
| `USERS_BE_OUTBOX_CLEANUP_INTERVAL` | duration | `1h` | Cleanup interval |
| `USERS_BE_OUTBOX_CLEANUP_RETENTION_DAYS` | int | `7` | Days to retain events |

## Fetching Secrets from Kubernetes

For local development against Kubernetes infrastructure:

```bash
# Database password
export USERS_BE_DATABASE_PASSWORD=$(kubectl get secret users-be-postgres -n skillsier -o jsonpath='{.data.POSTGRES_PASSWORD}' | base64 -d)

# Kafka password
export USERS_BE_KAFKA_SASL_PASSWORD=$(kubectl get secret admin-user -n kafka -o jsonpath='{.data.password}' | base64 -d)

# Keycloak client secret
export USERS_BE_KEYCLOAK_CLIENT_SECRET=$(kubectl get secret keycloak-client-users-be -n keycloak -o jsonpath='{.data.client-secret}' | base64 -d)
```

Or use the provided script:
```bash
source scripts/get-secrets.sh
```

## Example Configurations

### Development (Local)
```bash
USERS_BE_APP_ENVIRONMENT=development
USERS_BE_APP_LOG_LEVEL=debug
USERS_BE_DATABASE_HOST=173.212.218.251
USERS_BE_DATABASE_PORT=30432
USERS_BE_DATABASE_AUTO_MIGRATE=true
USERS_BE_KAFKA_TLS_SKIP_VERIFY=true
```

### Production
```bash
USERS_BE_APP_ENVIRONMENT=production
USERS_BE_APP_LOG_LEVEL=info
USERS_BE_DATABASE_SSL_MODE=require
USERS_BE_DATABASE_AUTO_MIGRATE=false
USERS_BE_DATABASE_MIGRATION_ALLOW_IN_PRODUCTION=false
USERS_BE_KAFKA_TLS_SKIP_VERIFY=false
```

## Config File (Optional)

Instead of environment variables, you can use a `config.yaml` file:

```yaml
app:
  name: users-be
  environment: development
  log_level: info

server:
  port: "8080"
  enable_cors: true

database:
  host: 173.212.218.251
  port: 30432
  user: usersuser
  database: usersdb
  auto_migrate: true

kafka:
  brokers:
    - 173.212.218.251:31691
  topic_prefix: users
```

Environment variables take precedence over config file values.