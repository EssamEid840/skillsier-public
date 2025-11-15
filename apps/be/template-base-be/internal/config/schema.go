package config

import "time"

// Config holds all configuration for the <service>-be microservice
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
	Name        string `mapstructure:"name"`        // Application name (e.g., "<service>-be")
	Environment string `mapstructure:"environment"` // Environment: development, staging, production
	LogLevel    string `mapstructure:"log_level"`   // Log level: debug, info, warn, error
	Version     string `mapstructure:"version"`     // Application version
}

// ServerConfig contains HTTP server settings
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
	SASLMechanism    string        `mapstructure:"sasl_mechanism"`    // SASL mechanism: PLAIN, SCRAM-SHA-256, SCRAM-SHA-512
	SASLUsername     string        `mapstructure:"sasl_username"`     // SASL username
	SASLPassword     string        `mapstructure:"sasl_password"`     // SASL password
	TLSSkipVerify    bool          `mapstructure:"tls_skip_verify"`   // Skip TLS verification (dev only)
	ConsumerGroup    string        `mapstructure:"consumer_group"`    // Consumer group ID
	TopicPrefix      string        `mapstructure:"topic_prefix"`      // Topic prefix (e.g., "<service>")
	ProducerConfig   ProducerConfig `mapstructure:"producer"`         // Producer-specific config
	ConsumerConfig   ConsumerConfig `mapstructure:"consumer"`         // Consumer-specific config
}

// ProducerConfig contains Kafka producer settings
type ProducerConfig struct {
	RequiredAcks      int           `mapstructure:"required_acks"`      // Required acks: 0, 1, -1 (all)
	MaxMessageBytes   int           `mapstructure:"max_message_bytes"`  // Max message size
	Compression       string        `mapstructure:"compression"`        // Compression: none, gzip, snappy, lz4, zstd
	IdempotentWrites  bool          `mapstructure:"idempotent_writes"`  // Enable idempotent writes
	MaxRetries        int           `mapstructure:"max_retries"`        // Max produce retries
	RetryBackoff      time.Duration `mapstructure:"retry_backoff"`      // Retry backoff duration
}

// ConsumerConfig contains Kafka consumer settings
type ConsumerConfig struct {
	EnableAutoCommit  bool          `mapstructure:"enable_auto_commit"` // Enable auto-commit
	AutoCommitInterval time.Duration `mapstructure:"auto_commit_interval"` // Auto-commit interval
	SessionTimeout    time.Duration `mapstructure:"session_timeout"`    // Session timeout
	FetchMinBytes     int           `mapstructure:"fetch_min_bytes"`    // Min fetch bytes
	FetchMaxWait      time.Duration `mapstructure:"fetch_max_wait"`     // Max fetch wait time
	MaxProcessingTime time.Duration `mapstructure:"max_processing_time"`// Max message processing time
}

// KeycloakConfig contains Keycloak authentication settings
type KeycloakConfig struct {
	URL          string        `mapstructure:"url"`           // Keycloak base URL
	Realm        string        `mapstructure:"realm"`         // Keycloak realm
	ClientID     string        `mapstructure:"client_id"`     // Client ID
	ClientSecret string        `mapstructure:"client_secret"` // Client secret
	Timeout      time.Duration `mapstructure:"timeout"`       // HTTP client timeout
	JWKSCacheTTL time.Duration `mapstructure:"jwks_cache_ttl"`// JWKS cache TTL
}

// DaprConfig contains Dapr sidecar settings
type DaprConfig struct {
	Enabled        bool   `mapstructure:"enabled"`          // Enable Dapr integration
	HTTPPort       int    `mapstructure:"http_port"`        // Dapr HTTP port
	GRPCPort       int    `mapstructure:"grpc_port"`        // Dapr gRPC port
	PubSubName     string `mapstructure:"pubsub_name"`      // PubSub component name
	StateStoreName string `mapstructure:"state_store_name"` // State store component name
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

// GetDSN returns the PostgreSQL connection string
func (d *DatabaseConfig) GetDSN() string {
	return "host=" + d.Host +
		" port=" + string(rune(d.Port)) +
		" user=" + d.User +
		" password=" + d.Password +
		" dbname=" + d.Database +
		" sslmode=" + d.SSLMode
}

// IsDevelopment returns true if running in development environment
func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}

// IsStaging returns true if running in staging environment
func (c *Config) IsStaging() bool {
	return c.App.Environment == "staging"
}