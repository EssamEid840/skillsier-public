package postgres

import (
	"fmt"
	"time"

	"users-be/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewConnection creates a new PostgreSQL database connection
func NewConnection(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	// Configure GORM logger based on environment
	logLevel := logger.Info
	if cfg.SSLMode == "disable" {
		logLevel = logger.Warn // Less verbose in production
	}

	// Open database connection
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			// Use UTC for all timestamps
			return time.Now().UTC()
		},
		// Prepare statements for better performance
		PrepareStmt: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying *sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying database: %w", err)
	}

	// Configure connection pool
	// MaxOpenConns: Maximum number of open connections to the database
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	
	// MaxIdleConns: Maximum number of idle connections
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	
	// ConnMaxLifetime: Maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// Close closes the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}