package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"<module>/internal/config"
)

// NewConnection creates a new PostgreSQL database connection with connection pooling
func NewConnection(cfg config.DatabaseConfig) (*gorm.DB, error) {
	// Build DSN
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)

	// Configure GORM logger
	gormLogger := logger.Default
	if cfg.AutoMigrate {
		// Use info level when migrations are enabled to see SQL
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		// Use warn level in production
		gormLogger = logger.Default.LogMode(logger.Warn)
	}

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 gormLogger,
		SkipDefaultTransaction: true, // Disable default transaction for better performance
		PrepareStmt:            true, // Cache prepared statements
		NowFunc: func() time.Time {
			return time.Now().UTC() // Always use UTC
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying SQL DB for connection pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Printf("✓ Database connected: %s:%d/%s (pool: %d max open, %d max idle)",
		cfg.Host, cfg.Port, cfg.Database, cfg.MaxOpenConns, cfg.MaxIdleConns)

	return db, nil
}

// Close closes the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	log.Println("✓ Database connection closed")
	return nil
}

// Ping checks if the database connection is alive
func Ping(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

// Stats returns connection pool statistics
func Stats(db *gorm.DB) (map[string]interface{}, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	stats := sqlDB.Stats()
	return map[string]interface{}{
		"max_open_connections":  stats.MaxOpenConnections,
		"open_connections":      stats.OpenConnections,
		"in_use":                stats.InUse,
		"idle":                  stats.Idle,
		"wait_count":            stats.WaitCount,
		"wait_duration_ms":      stats.WaitDuration.Milliseconds(),
		"max_idle_closed":       stats.MaxIdleClosed,
		"max_idle_time_closed":  stats.MaxIdleTimeClosed,
		"max_lifetime_closed":   stats.MaxLifetimeClosed,
	}, nil
}