// package postgres

// import (
// 	"fmt"

// 	"users-be/internal/domain/outbox"
// 	"users-be/internal/domain/user"

// 	"gorm.io/gorm"
// )

// // AutoMigrate runs automatic database migrations
// // This creates/updates tables based on the struct definitions
// func AutoMigrate(db *gorm.DB) error {
// 	// List of all models to migrate
// 	// Add new models here as your application grows
// 	models := []interface{}{
// 		&user.User{},
// 		&outbox.Event{},
// 	}

// 	// Run auto-migration for all models
// 	if err := db.AutoMigrate(models...); err != nil {
// 		return fmt.Errorf("failed to run auto-migration: %w", err)
// 	}

// 	// Create indexes for better query performance
// 	if err := createIndexes(db); err != nil {
// 		return fmt.Errorf("failed to create indexes: %w", err)
// 	}

// 	return nil
// }

// // createIndexes creates additional indexes not covered by GORM tags
// func createIndexes(db *gorm.DB) error {
// 	// Users table indexes
// 	// These indexes improve query performance for common operations
	
// 	// Index for keycloak_id lookups (most common query)
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_users_keycloak_id_not_deleted 
// 		ON users(keycloak_id) 
// 		WHERE deleted_at IS NULL
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create users keycloak_id index: %w", err)
// 	}

// 	// Index for email lookups
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_users_email_not_deleted 
// 		ON users(email) 
// 		WHERE deleted_at IS NULL
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create users email index: %w", err)
// 	}

// 	// Index for username lookups
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_users_username_not_deleted 
// 		ON users(username) 
// 		WHERE deleted_at IS NULL
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create users username index: %w", err)
// 	}

// 	// Composite index for active users lookup
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_users_active_created 
// 		ON users(is_active, created_at DESC) 
// 		WHERE deleted_at IS NULL
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create users active index: %w", err)
// 	}

// 	// Index for profile type filtering
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_users_profile_type 
// 		ON users(profile_type) 
// 		WHERE deleted_at IS NULL AND is_active = true
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create users profile type index: %w", err)
// 	}

// 	// Outbox events indexes
// 	// These indexes optimize the outbox processor queries
	
// 	// Composite index for finding pending events (most important!)
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_outbox_pending_events 
// 		ON outbox_events(status, created_at ASC) 
// 		WHERE status IN ('pending', 'failed')
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create outbox pending events index: %w", err)
// 	}

// 	// Index for retry logic
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_outbox_retry 
// 		ON outbox_events(status, next_retry_at, retry_count) 
// 		WHERE status = 'failed' AND next_retry_at IS NOT NULL
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create outbox retry index: %w", err)
// 	}

// 	// Index for event sourcing queries (by aggregate)
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_outbox_aggregate 
// 		ON outbox_events(aggregate_type, aggregate_id, created_at DESC)
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create outbox aggregate index: %w", err)
// 	}

// 	// Index for cleanup queries (finding old published events)
// 	if err := db.Exec(`
// 		CREATE INDEX IF NOT EXISTS idx_outbox_cleanup 
// 		ON outbox_events(status, published_at) 
// 		WHERE status = 'published'
// 	`).Error; err != nil {
// 		return fmt.Errorf("failed to create outbox cleanup index: %w", err)
// 	}

// 	return nil
// }

// // DropAllTables drops all tables (use with caution!)
// // This is useful for testing or resetting the database
// func DropAllTables(db *gorm.DB) error {
// 	return db.Migrator().DropTable(
// 		&user.User{},
// 		&outbox.Event{},
// 	)
// }