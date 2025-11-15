package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Load reads configuration from environment variables, config file, and defaults
// Priority order: 1. Environment variables, 2. Config file, 3. Defaults
func Load() *Config {
	v := viper.New()

	// Set config file properties
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("/etc/<service>-be")

	// Enable environment variables
	v.SetEnvPrefix("<SERVICE>_BE")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Set defaults
	setDefaults(v)

	// Read config file (optional - don't fail if missing)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Warning: Error reading config file: %v\n", err)
		}
	}

	// Unmarshal into Config struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		os.Exit(1)
	}

	// Validate required fields
	if err := validate(&config); err != nil {
		fmt.Printf("Configuration validation failed: %v\n", err)
		os.Exit(1)
	}

	return &config
}

// setDefaults sets default configuration values
func setDefaults(v *viper.Viper) {
	// App defaults
	v.SetDefault("app.name", "<service>-be")
	v.SetDefault("app.environment", "development")
	v.SetDefault("app.log_level", "info")
	v.SetDefault("app.version", "1.0.0")

	// Server defaults
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.read_timeout", 30*time.Second)
	v.SetDefault("server.write_timeout", 30*time.Second)
	v.SetDefault("server.idle_timeout", 120*time.Second)
	v.SetDefault("server.shutdown_timeout", 10*time.Second)
	v.SetDefault("server.enable_cors", true)
	v.SetDefault("server.trusted_proxies", []string{})

	// Database defaults
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "<db_user>")
	v.SetDefault("database.database", "<db_name>")
	v.SetDefault("database.ssl_mode", "disable")
	v.SetDefault("database.max_open_conns", 25)
	v.SetDefault("database.max_idle_conns", 5)
	v.SetDefault("database.conn_max_lifetime", 30*time.Minute)
	v.SetDefault("database.conn_max_idle_time", 10*time.Minute)
	v.SetDefault("database.auto_migrate", true)

	// Migration defaults
	v.SetDefault("database.migration.enabled", true)
	v.SetDefault("database.migration.allow_in_production", false)
	v.SetDefault("database.migration.allow_destructive", false)
	v.SetDefault("database.migration.backup_before_change", true)
	v.SetDefault("database.migration.log_migration_summary", true)

	// Kafka defaults
	v.SetDefault("kafka.brokers", []string{"localhost:9092"})
	v.SetDefault("kafka.security_protocol", "SASL_SSL")
	v.SetDefault("kafka.sasl_mechanism", "SCRAM-SHA-512")
	v.SetDefault("kafka.sasl_username", "admin-user")
	v.SetDefault("kafka.tls_skip_verify", false)
	v.SetDefault("kafka.consumer_group", "<service>-be-group")
	v.SetDefault("kafka.topic_prefix", "<service>")

	// Kafka producer defaults
	v.SetDefault("kafka.producer.required_acks", -1) // Wait for all replicas
	v.SetDefault("kafka.producer.max_message_bytes", 1000000) // 1MB
	v.SetDefault("kafka.producer.compression", "snappy")
	v.SetDefault("kafka.producer.idempotent_writes", true)
	v.SetDefault("kafka.producer.max_retries", 5)
	v.SetDefault("kafka.producer.retry_backoff", 100*time.Millisecond)

	// Kafka consumer defaults
	v.SetDefault("kafka.consumer.enable_auto_commit", false)
	v.SetDefault("kafka.consumer.auto_commit_interval", 1*time.Second)
	v.SetDefault("kafka.consumer.session_timeout", 10*time.Second)
	v.SetDefault("kafka.consumer.fetch_min_bytes", 1)
	v.SetDefault("kafka.consumer.fetch_max_wait", 500*time.Millisecond)
	v.SetDefault("kafka.consumer.max_processing_time", 30*time.Second)

	// Keycloak defaults
	v.SetDefault("keycloak.url", "http://localhost:8080")
	v.SetDefault("keycloak.realm", "skillsier")
	v.SetDefault("keycloak.client_id", "<service>-be")
	v.SetDefault("keycloak.timeout", 10*time.Second)
	v.SetDefault("keycloak.jwks_cache_ttl", 24*time.Hour)

	// Dapr defaults
	v.SetDefault("dapr.enabled", false)
	v.SetDefault("dapr.http_port", 3500)
	v.SetDefault("dapr.grpc_port", 50001)
	v.SetDefault("dapr.pubsub_name", "kafka-pubsub")
	v.SetDefault("dapr.state_store_name", "statestore")

	// Outbox defaults
	v.SetDefault("outbox.enabled", true)
	v.SetDefault("outbox.poll_interval", 5*time.Second)
	v.SetDefault("outbox.batch_size", 100)
	v.SetDefault("outbox.max_retries", 10)
	v.SetDefault("outbox.retry_backoff_base", 1*time.Second)
	v.SetDefault("outbox.retry_backoff_max", 5*time.Minute)
	v.SetDefault("outbox.cleanup_enabled", true)
	v.SetDefault("outbox.cleanup_interval", 1*time.Hour)
	v.SetDefault("outbox.cleanup_retention_days", 7)
}

// validate checks that required configuration fields are present
func validate(config *Config) error {
	// Validate database
	if config.Database.Host == "" {
		return fmt.Errorf("database.host is required")
	}
	if config.Database.User == "" {
		return fmt.Errorf("database.user is required")
	}
	if config.Database.Password == "" {
		return fmt.Errorf("database.password is required")
	}
	if config.Database.Database == "" {
		return fmt.Errorf("database.database is required")
	}

	// Validate Kafka
	if len(config.Kafka.Brokers) == 0 {
		return fmt.Errorf("kafka.brokers is required")
	}
	if config.Kafka.SASLUsername == "" {
		return fmt.Errorf("kafka.sasl_username is required")
	}
	if config.Kafka.SASLPassword == "" {
		return fmt.Errorf("kafka.sasl_password is required")
	}

	// Validate Keycloak
	if config.Keycloak.URL == "" {
		return fmt.Errorf("keycloak.url is required")
	}
	if config.Keycloak.Realm == "" {
		return fmt.Errorf("keycloak.realm is required")
	}
	if config.Keycloak.ClientID == "" {
		return fmt.Errorf("keycloak.client_id is required")
	}

	// Validate environment-specific settings
	if config.IsProduction() {
		// In production, auto-migrate must be explicitly enabled
		if config.Database.AutoMigrate && !config.Database.MigrationConfig.AllowInProduction {
			return fmt.Errorf("auto_migrate is enabled in production but allow_in_production is false")
		}

		// In production, TLS verification should be enabled
		if config.Kafka.TLSSkipVerify {
			fmt.Println("Warning: TLS verification is disabled in production - this is insecure!")
		}

		// In production, destructive migrations should be disabled
		if config.Database.MigrationConfig.AllowDestructive {
			fmt.Println("Warning: Destructive migrations are allowed in production - this is dangerous!")
		}
	}

	return nil
}