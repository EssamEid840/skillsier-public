# Database Migrations

This microservice uses **automatic migrations** as the primary migration method. However, this folder is provided for teams that prefer deterministic SQL migrations.

## Primary Method: Automatic Migrations

The service automatically creates and updates database tables on startup based on registered entity models.

**How it works:**
1. Entity models are registered in `internal/infrastructure/persistence/postgres/migrations.go`
2. On startup, GORM `AutoMigrate` creates/updates tables to match the models
3. Schema version is tracked in the `schema_versions` table

**To add a new entity:**
1. Create your entity model in `internal/domain/<entity_name>/entity.go`
2. Register it in `migrations.go`:
   ```go
   func init() {
       Register(
           &initial_entity.InitialEntity{},
           &your_entity.YourEntity{},  // Add this line
       )
   }
   ```
3. Restart the service - the table will be created automatically

**Production safety:**
- Auto-migrations are disabled by default in production
- Set `AUTO_MIGRATE=true` and `ALLOW_IN_PRODUCTION=true` to enable
- Destructive changes require `ALLOW_DESTRUCTIVE=true`

## Alternative: Deterministic SQL Migrations

If you prefer manual SQL migrations, you can place them in this folder:

```
migrations/
├── 000001_create_initial_entities_table.sql
├── 000002_create_outbox_table.sql
├── 000003_add_indexes.sql
└── README.md (this file)
```

### Migration File Format

Each migration file should be numbered sequentially:

```sql
-- 000001_create_initial_entities_table.sql
-- Migration: Create initial_entities table
-- Author: Your Name
-- Date: 2024-11-14

CREATE TABLE IF NOT EXISTS initial_entities (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    owner_id UUID NOT NULL,
    metadata_tags JSONB,
    metadata_properties JSONB,
    metadata_version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_initial_entities_name ON initial_entities(name);
CREATE INDEX idx_initial_entities_status ON initial_entities(status);
CREATE INDEX idx_initial_entities_owner_id ON initial_entities(owner_id);
CREATE INDEX idx_initial_entities_deleted_at ON initial_entities(deleted_at);
```

### Running Manual Migrations

If you choose to use manual migrations:

1. **Disable auto-migrations:**
   ```bash
   export <SERVICE>_BE_DATABASE_AUTO_MIGRATE=false
   ```

2. **Use a migration tool** like [golang-migrate](https://github.com/golang-migrate/migrate):
   ```bash
   migrate -path ./migrations -database "postgresql://user:pass@host:5432/dbname?sslmode=disable" up
   ```

3. **Or use psql:**
   ```bash
   psql -h host -U user -d dbname -f migrations/000001_create_initial_entities_table.sql
   ```

## Outbox Table Schema

The outbox table is required for reliable event publishing:

```sql
CREATE TABLE IF NOT EXISTS outbox_events (
    id UUID PRIMARY KEY,
    aggregate_id VARCHAR(255) NOT NULL,
    aggregate_type VARCHAR(100) NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    event_version INTEGER NOT NULL DEFAULT 1,
    payload JSONB NOT NULL,
    metadata JSONB,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    topic VARCHAR(100) NOT NULL,
    error_message TEXT,
    published_at BIGINT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    retry_count INTEGER DEFAULT 0,
    next_retry_at BIGINT
);

CREATE INDEX idx_outbox_events_aggregate_id ON outbox_events(aggregate_id);
CREATE INDEX idx_outbox_events_event_type ON outbox_events(event_type);
CREATE INDEX idx_outbox_events_status ON outbox_events(status);
CREATE INDEX idx_outbox_events_created_at ON outbox_events(created_at);
CREATE INDEX idx_outbox_events_next_retry_at ON outbox_events(next_retry_at);
CREATE INDEX idx_outbox_events_published_at ON outbox_events(published_at);
```

## Schema Versions Table

For tracking migration history:

```sql
CREATE TABLE IF NOT EXISTS schema_versions (
    id BIGSERIAL PRIMARY KEY,
    version INTEGER NOT NULL UNIQUE,
    description TEXT,
    applied_at BIGINT NOT NULL,
    applied_by VARCHAR(100)
);

CREATE INDEX idx_schema_versions_version ON schema_versions(version);
CREATE INDEX idx_schema_versions_applied_at ON schema_versions(applied_at);
```

## Best Practices

1. **Version Control**: Always commit migration files to git
2. **Never Modify**: Once a migration is applied, never modify it - create a new one
3. **Idempotency**: Use `IF NOT EXISTS`, `IF EXISTS` to make migrations idempotent
4. **Test First**: Test migrations on a staging database before production
5. **Backup**: Always backup production database before running migrations
6. **Rollback Plan**: Prepare rollback migrations for destructive changes

## Recommendation

For development and small teams, **use automatic migrations** - it's simpler and faster.

For production and large teams, consider **manual SQL migrations** for:
- Better control over schema changes
- Explicit migration history
- Easier code review
- More predictable deployments

You can also use a **hybrid approach**: automatic migrations in development, manual in production.