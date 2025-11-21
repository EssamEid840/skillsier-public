package postgres

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// SchemaVersion represents the schema_versions table for tracking migration history
type SchemaVersion struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Version     int    `gorm:"not null;uniqueIndex"`
	Description string `gorm:"type:text"`
	AppliedAt   int64  `gorm:"not null;index"`
	AppliedBy   string `gorm:"type:varchar(100)"`
}

// TableName returns the table name for GORM
func (SchemaVersion) TableName() string {
	return "schema_versions"
}

// ensureSchemaVersionTable creates the schema_versions table if it doesn't exist
func ensureSchemaVersionTable(db *gorm.DB) error {
	return db.AutoMigrate(&SchemaVersion{})
}

// GetSchemaVersion retrieves the current schema version
func GetSchemaVersion(db *gorm.DB) (int, error) {
	// Ensure table exists
	if err := ensureSchemaVersionTable(db); err != nil {
		return 0, fmt.Errorf("failed to ensure schema_versions table: %w", err)
	}

	var version SchemaVersion
	err := db.Order("version DESC").First(&version).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil // No versions recorded yet
		}
		return 0, fmt.Errorf("failed to get schema version: %w", err)
	}

	return version.Version, nil
}

// SetSchemaVersion records a new schema version
func SetSchemaVersion(db *gorm.DB, version int) error {
	// Ensure table exists
	if err := ensureSchemaVersionTable(db); err != nil {
		return fmt.Errorf("failed to ensure schema_versions table: %w", err)
	}

	schemaVersion := SchemaVersion{
		Version:     version,
		Description: fmt.Sprintf("Auto-migration to version %d", version),
		AppliedAt:   time.Now().Unix(),
		AppliedBy:   "auto-migration",
	}

	if err := db.Create(&schemaVersion).Error; err != nil {
		return fmt.Errorf("failed to set schema version: %w", err)
	}

	return nil
}

// GetMigrationHistory retrieves the full migration history
func GetMigrationHistory(db *gorm.DB) ([]SchemaVersion, error) {
	// Ensure table exists
	if err := ensureSchemaVersionTable(db); err != nil {
		return nil, fmt.Errorf("failed to ensure schema_versions table: %w", err)
	}

	var versions []SchemaVersion
	err := db.Order("version DESC").Find(&versions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get migration history: %w", err)
	}

	return versions, nil
}

// IsMigrated checks if a specific version has been applied
func IsMigrated(db *gorm.DB, version int) (bool, error) {
	// Ensure table exists
	if err := ensureSchemaVersionTable(db); err != nil {
		return false, fmt.Errorf("failed to ensure schema_versions table: %w", err)
	}

	var count int64
	err := db.Model(&SchemaVersion{}).Where("version = ?", version).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check migration status: %w", err)
	}

	return count > 0, nil
}

// GetLatestVersion returns the latest schema version
func GetLatestVersion(db *gorm.DB) (*SchemaVersion, error) {
	// Ensure table exists
	if err := ensureSchemaVersionTable(db); err != nil {
		return nil, fmt.Errorf("failed to ensure schema_versions table: %w", err)
	}

	var version SchemaVersion
	err := db.Order("version DESC").First(&version).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No versions recorded yet
		}
		return nil, fmt.Errorf("failed to get latest version: %w", err)
	}

	return &version, nil
}

// PrintMigrationHistory prints the migration history
func PrintMigrationHistory(db *gorm.DB) error {
	versions, err := GetMigrationHistory(db)
	if err != nil {
		return err
	}

	if len(versions) == 0 {
		fmt.Println("No migration history found")
		return nil
	}

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("               MIGRATION HISTORY")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Printf("%-10s %-30s %-20s\n", "Version", "Applied At", "Applied By")
	fmt.Println("───────────────────────────────────────────────────────────")

	for _, v := range versions {
		appliedAt := time.Unix(v.AppliedAt, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("%-10d %-30s %-20s\n", v.Version, appliedAt, v.AppliedBy)
		if v.Description != "" {
			fmt.Printf("           %s\n", v.Description)
		}
	}

	fmt.Println("═══════════════════════════════════════════════════════════")
	return nil
}