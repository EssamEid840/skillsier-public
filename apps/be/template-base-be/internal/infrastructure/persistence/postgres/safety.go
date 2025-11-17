package postgres

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"<module>/internal/config"
)

// runSafetyChecks performs pre-migration safety checks
func runSafetyChecks(db *gorm.DB, cfg config.DatabaseConfig) error {
	log.Println("→ Running pre-migration safety checks...")

	// Check 1: Database connectivity
	if err := Ping(db); err != nil {
		return fmt.Errorf("database connectivity check failed: %w", err)
	}
	log.Println("  ✓ Database connectivity OK")

	// Check 2: Sufficient disk space (if possible)
	if err := checkDiskSpace(); err != nil {
		log.Printf("  ⚠ Could not verify disk space: %v", err)
	} else {
		log.Println("  ✓ Disk space OK")
	}

	// Check 3: Check if tables exist (for first-time migration)
	if err := checkTablesExist(db); err != nil {
		log.Printf("  ℹ First-time migration detected: %v", err)
	} else {
		log.Println("  ✓ Existing tables detected")
	}

	// Check 4: Environment validation
	if err := checkEnvironment(cfg); err != nil {
		return fmt.Errorf("environment validation failed: %w", err)
	}
	log.Println("  ✓ Environment validation OK")

	log.Println("✓ All safety checks passed")
	return nil
}

// checkDiskSpace checks if there is sufficient disk space
func checkDiskSpace() error {
	// This is a simplified check - in production you might want to use
	// platform-specific APIs to check actual disk space
	
	// For now, just check if we can create a temp file
	tmpFile, err := os.CreateTemp("", "diskcheck")
	if err != nil {
		return fmt.Errorf("cannot write to disk: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Try to write 1MB
	data := make([]byte, 1024*1024)
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("disk write test failed: %w", err)
	}

	return nil
}

// checkTablesExist checks if any tables exist in the database
func checkTablesExist(db *gorm.DB) error {
	var count int64
	
	// Query for any tables
	err := db.Raw(`
		SELECT COUNT(*) 
		FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_type = 'BASE TABLE'
	`).Scan(&count).Error
	
	if err != nil {
		return fmt.Errorf("failed to check for existing tables: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("no tables exist (first-time migration)")
	}

	return nil
}

// checkEnvironment validates the environment configuration
func checkEnvironment(cfg config.DatabaseConfig) error {
	// Ensure required environment variables are set
	if cfg.Host == "" {
		return fmt.Errorf("database host not configured")
	}

	if cfg.Database == "" {
		return fmt.Errorf("database name not configured")
	}

	if cfg.User == "" {
		return fmt.Errorf("database user not configured")
	}

	if cfg.Password == "" {
		return fmt.Errorf("database password not configured")
	}

	// Validate connection pool settings
	if cfg.MaxOpenConns < 1 {
		return fmt.Errorf("max_open_conns must be at least 1")
	}

	if cfg.MaxIdleConns < 1 {
		return fmt.Errorf("max_idle_conns must be at least 1")
	}

	if cfg.MaxIdleConns > cfg.MaxOpenConns {
		return fmt.Errorf("max_idle_conns cannot exceed max_open_conns")
	}

	return nil
}

// CheckDestructiveChanges checks if the migration would perform destructive changes
// This is a placeholder - in a real implementation, you would compare the current schema
// with the target schema to detect destructive changes
func CheckDestructiveChanges(db *gorm.DB, models []interface{}) (bool, error) {
	// TODO: Implement actual destructive change detection
	// This would involve:
	// 1. Getting the current schema from the database
	// 2. Comparing it with the schema that would be created by the models
	// 3. Detecting column drops, type changes, constraint changes, etc.
	
	// For now, we'll return false (no destructive changes detected)
	return false, nil
}

// BackupSchema creates a backup of the current schema
// This is a placeholder - in production, you would implement actual backup logic
func BackupSchema(db *gorm.DB) error {
	log.Println("  → Creating schema backup...")
	
	// TODO: Implement actual schema backup
	// This could involve:
	// 1. Exporting the schema using pg_dump
	// 2. Storing it in a versioned location
	// 3. Optionally backing up data as well
	
	log.Println("  ⚠ Schema backup not implemented (placeholder)")
	return nil
}

// VerifyMigrationIntegrity verifies that the migration completed successfully
func VerifyMigrationIntegrity(db *gorm.DB, models []interface{}) error {
	log.Println("→ Verifying migration integrity...")

	for _, model := range models {
		tableName := getTableName(db, model)
		
		// Check if table exists
		var exists bool
		err := db.Raw(`
			SELECT EXISTS (
				SELECT FROM information_schema.tables 
				WHERE table_schema = 'public' 
				AND table_name = ?
			)
		`, tableName).Scan(&exists).Error
		
		if err != nil {
			return fmt.Errorf("failed to verify table %s: %w", tableName, err)
		}

		if !exists {
			return fmt.Errorf("table %s was not created", tableName)
		}
	}

	log.Println("✓ Migration integrity verified")
	return nil
}