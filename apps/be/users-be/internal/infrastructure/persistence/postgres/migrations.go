package postgres

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"

	"skillsier.dev/apps/be/users-be/internal/config"
	"skillsier.dev/apps/be/users-be/internal/domain/initial_entity"
	
	platformOutboxPostgres "skillsier.dev/platform-shared/outbox/postgres"
)

// init registers all entity models for auto-migration
func init() {
	// Register domain entities here
	// When you add a new entity, add it to this list:
	Register(
		&initial_entity.InitialEntity{},
		// Add your new entities here, e.g.:
		// &your_entity.YourEntity{},
	)

	// Register outbox table from platform-shared
	Register(&platformOutboxPostgres.OutboxEvent{})
}

// AutoMigrate performs automatic database migrations based on registered entity models
func AutoMigrate(db *gorm.DB, cfg config.DatabaseConfig) error {
	migrationConfig := cfg.MigrationConfig

	// Check if auto-migration is enabled
	if !cfg.AutoMigrate || !migrationConfig.Enabled {
		log.Println("⊗ Auto-migration is disabled")
		return nil
	}

	// Production safety check
	if err := checkProductionSafety(cfg); err != nil {
		return fmt.Errorf("production safety check failed: %w", err)
	}

	// Pre-migration safety checks
	if err := runSafetyChecks(db, cfg); err != nil {
		return fmt.Errorf("safety checks failed: %w", err)
	}

	// Get registered models
	models := GetRegisteredModels()
	if len(models) == 0 {
		log.Println("⚠ No entity models registered for auto-migration")
		return nil
	}

	log.Printf("→ Starting auto-migration for %d entity models...", len(models))

	// Print registered models
	if migrationConfig.LogMigrationSummary {
		PrintRegisteredModels()
	}

	// Track schema version before migration
	versionBefore, err := GetSchemaVersion(db)
	if err != nil {
		log.Printf("⚠ Could not get schema version before migration: %v", err)
		versionBefore = 0
	}

	// Perform auto-migration
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("auto-migration failed: %w", err)
	}

	// Update schema version
	versionAfter := versionBefore + 1
	if err := SetSchemaVersion(db, versionAfter); err != nil {
		log.Printf("⚠ Could not update schema version: %v", err)
	}

	// Log migration summary
	if migrationConfig.LogMigrationSummary {
		logMigrationSummary(db, models, versionBefore, versionAfter)
	}

	log.Printf("✓ Auto-migration completed successfully (schema version: %d → %d)", versionBefore, versionAfter)

	return nil
}

// checkProductionSafety ensures auto-migration is safe to run in production
func checkProductionSafety(cfg config.DatabaseConfig) error {
	// Check if running in production
	isProduction := strings.ToLower(cfg.SSLMode) == "require" || 
		strings.ToLower(cfg.SSLMode) == "verify-ca" ||
		strings.ToLower(cfg.SSLMode) == "verify-full"

	if isProduction {
		// In production, auto-migrate must be explicitly allowed
		if !cfg.MigrationConfig.AllowInProduction {
			return fmt.Errorf("auto-migration is disabled in production (set ALLOW_IN_PRODUCTION=true to override)")
		}

		log.Println("⚠ Running auto-migration in production environment!")

		// Warn about destructive changes
		if cfg.MigrationConfig.AllowDestructive {
			log.Println("⚠ WARNING: Destructive schema changes are ALLOWED in production!")
		}
	}

	return nil
}

// logMigrationSummary logs what was migrated
func logMigrationSummary(db *gorm.DB, models []interface{}, versionBefore, versionAfter int) {
	log.Println("═══════════════════════════════════════════════════════════")
	log.Println("               MIGRATION SUMMARY")
	log.Println("═══════════════════════════════════════════════════════════")
	log.Printf("Schema Version:    %d → %d", versionBefore, versionAfter)
	log.Printf("Models Migrated:   %d", len(models))
	log.Println("───────────────────────────────────────────────────────────")
	log.Println("Tables:")

	for _, model := range models {
		tableName := getTableName(db, model)
		log.Printf("  • %s (%T)", tableName, model)
	}

	log.Println("═══════════════════════════════════════════════════════════")
}

// getTableName gets the table name for a model
func getTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil {
		return "unknown"
	}
	return stmt.Schema.Table
}