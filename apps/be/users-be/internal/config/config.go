package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	Kafka    KafkaConfig
	Keycloak KeycloakConfig
	Dapr     DaprConfig
	Outbox   OutboxConfig
}

// AppConfig contains application-level settings
type AppConfig struct {
	Name        string `mapstructure:"name"`        // Application name (e.g., "users-be")
	Environment string `mapstructure:"environment"` // Environment: development, staging, production
	LogLevel    string `mapstructure:"log_level"`   // Log level: debug, info, warn, error
	Version     string `mapstructure:"version"`     // Application version
}
// ServerConfig holds HTTP server settings
type ServerConfig struct {
	Port            string        `mapstructure:"port"`              // HTTP port (default: 8080)
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`      // Request read timeout
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`     // Response write timeout
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`      // Keep-alive idle timeout
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`  // Graceful shutdown timeout
	EnableCORS      bool          `mapstructure:"enable_cors"`       // Enable CORS middleware
	TrustedProxies  []string      `mapstructure:"trusted_proxies"`   // Trusted proxy IPs
}


// DatabaseConfig contains PostgreSQL settings
type DatabaseConfig struct {
	Host            string        `mapstructure:"host"`              // Database host
	Port            int           `mapstructure:"port"`              // Database port
	User            string        `mapstructure:"user"`              // Database user
	Password        string        `mapstructure:"password"`          // Database password
	Database        string        `mapstructure:"database"`          // Database name
	SSLMode         string        `mapstructure:"ssl_mode"`          // SSL mode: disable, require, verify-ca, verify-full
	MaxOpenConns    int           `mapstructure:"max_open_conns"`    // Maximum open connections
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`    // Maximum idle connections
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"` // Connection max lifetime
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`// Connection max idle time
	AutoMigrate     bool          `mapstructure:"auto_migrate"`      // Enable automatic migrations
	MigrationConfig MigrationConfig `mapstructure:"migration"`       // Migration-specific config
}

// MigrationConfig controls auto-migration behavior
type MigrationConfig struct {
	Enabled              bool   `mapstructure:"enabled"`                // Enable auto-migration
	AllowInProduction    bool   `mapstructure:"allow_in_production"`    // Allow migrations in production
	AllowDestructive     bool   `mapstructure:"allow_destructive"`      // Allow destructive changes
	BackupBeforeChange   bool   `mapstructure:"backup_before_change"`   // Backup before migrations
	LogMigrationSummary  bool   `mapstructure:"log_migration_summary"`  // Log what was migrated
}



// KafkaConfig contains Kafka settings
type KafkaConfig struct {
	Brokers          []string      `mapstructure:"brokers"`           // Kafka broker addresses
	SecurityProtocol string        `mapstructure:"security_protocol"` // Security protocol: PLAINTEXT, SASL_SSL
	EnableSASL      bool           `mapstructure:"enable_sasl"`       // Enable SASL authentication
	EnableTLS      bool            `mapstructure:"enable_tls"`        // Enable TLS
	SASLMechanism    string        `mapstructure:"sasl_mechanism"`    // SASL mechanism: PLAIN, SCRAM-SHA-256, SCRAM-SHA-512
	SASLUsername     string        `mapstructure:"sasl_username"`     // SASL username
	SASLPassword     string        `mapstructure:"sasl_password"`     // SASL password
	TLSSkipVerify    bool          `mapstructure:"tls_skip_verify"`   // Skip TLS verification (dev only)
	ConsumerGroup    string        `mapstructure:"consumer_group"`    // Consumer group ID
	TopicPrefix      string        `mapstructure:"topic_prefix"`      // Topic prefix (e.g., "users")
	ProducerConfig   ProducerConfig `mapstructure:"producer"`         // Producer-specific config
	ConsumerConfig   ConsumerConfig `mapstructure:"consumer"`         // Consumer-specific config
}


// KeycloakConfig contains Keycloak authentication settings
type KeycloakConfig struct {
	BaseURL      string        `mapstructure:"base_url"`      // Keycloak base URL
	URL          string        `mapstructure:"url"`           // Keycloak URL
	Realm        string        `mapstructure:"realm"`         // Keycloak realm
	ClientID     string        `mapstructure:"client_id"`     // Client ID
	ClientSecret string        `mapstructure:"client_secret"` // Client secret
	Timeout      time.Duration `mapstructure:"timeout"`       // HTTP client timeout
	JWKSCacheTTL time.Duration `mapstructure:"jwks_cache_ttl"`// JWKS cache TTL
}

// OutboxConfig contains outbox pattern settings
type OutboxConfig struct {
	Enabled            bool          `mapstructure:"enabled"`              // Enable outbox dispatcher
	PollInterval       time.Duration `mapstructure:"poll_interval"`        // Polling interval for pending events
	BatchSize          int           `mapstructure:"batch_size"`           // Max events to process per batch
	MaxRetries         int           `mapstructure:"max_retries"`          // Max retry attempts
	RetryBackoffBase   time.Duration `mapstructure:"retry_backoff_base"`   // Base backoff duration
	RetryBackoffMax    time.Duration `mapstructure:"retry_backoff_max"`    // Max backoff duration
	CleanupEnabled     bool          `mapstructure:"cleanup_enabled"`      // Enable cleanup of old events
	CleanupInterval    time.Duration `mapstructure:"cleanup_interval"`     // Cleanup interval
	CleanupRetentionDays int         `mapstructure:"cleanup_retention_days"` // Days to retain events
}

// Load loads configuration from environment variables
func Load() *Config {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("USERS_BE")

	return &Config{
		App: AppConfig{
			Name:        getEnv("APP_NAME", "users-be"),
			Version:     getEnv("APP_VERSION", "1.0.0"),
			Environment: getEnv("ENVIRONMENT", "development"),
			LogLevel:    getEnv("LOG_LEVEL", "info"),
		},
		Server: ServerConfig{
			Port:       getEnv("SERVER_PORT", "8080"),
			EnableCORS: getEnvBool("SERVER_ENABLE_CORS", true),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DATABASE_HOST", "173.212.218.251"),
			Port:            getEnvInt("DATABASE_PORT", 5432),
			Database:        getEnv("DATABASE_NAME", "usersdb"),
			User:            getEnv("DATABASE_USER", "usersuser"),
			Password:        getEnv("DATABASE_PASSWORD", "CHANGE_ME_IN_PRODUCTION"),
			SSLMode:         getEnv("DATABASE_SSL_MODE", "disable"),
			MaxOpenConns:    getEnvInt("DATABASE_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("DATABASE_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvDuration("DATABASE_CONN_MAX_LIFETIME", 5*time.Minute),
			AutoMigrate:     getEnvBool("DATABASE_AUTO_MIGRATE", true),
			MigrationConfig: MigrationConfig{
				Enabled:              getEnvBool("MIGRATION_ENABLED", true),
				AllowInProduction:    getEnvBool("MIGRATION_ALLOW_IN_PRODUCTION", false),
				AllowDestructive:     getEnvBool("MIGRATION_ALLOW_DESTRUCTIVE", false),
				LogMigrationSummary:  getEnvBool("MIGRATION_LOG_SUMMARY", true),
			},
		},
		Kafka: KafkaConfig{
			Brokers:      getEnvStringSlice("KAFKA_BROKERS", []string{"localhost:9092"}),
			TopicPrefix:  getEnv("KAFKA_TOPIC_PREFIX", "users"),
			EnableSASL:   getEnvBool("KAFKA_ENABLE_SASL", false),
			SASLUsername: getEnv("KAFKA_SASL_USERNAME", ""),
			SASLPassword: getEnv("KAFKA_SASL_PASSWORD", ""),
			EnableTLS:    getEnvBool("KAFKA_ENABLE_TLS", false),
		},
		Keycloak: KeycloakConfig{
			BaseURL:      getEnv("KEYCLOAK_BASE_URL", "http://localhost:8080"),
			Realm:        getEnv("KEYCLOAK_REALM", "skillsier"),
			ClientID:     getEnv("KEYCLOAK_CLIENT_ID", "users-be"),
			ClientSecret: getEnv("KEYCLOAK_CLIENT_SECRET", ""),
		},
		Outbox: OutboxConfig{
			PollInterval: getEnvDuration("OUTBOX_POLL_INTERVAL", 5*time.Second),
			BatchSize:    getEnvInt("OUTBOX_BATCH_SIZE", 100),
			MaxRetries:   getEnvInt("OUTBOX_MAX_RETRIES", 5),
		},
	}
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getEnvStringSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return []string{value}
	}
	return defaultValue
}

// DSN returns the PostgreSQL connection string
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode,
	)
}