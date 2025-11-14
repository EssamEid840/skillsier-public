package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Kafka    KafkaConfig
	Keycloak KeycloakConfig
	Dapr     DaprConfig
}

// AppConfig contains general application settings
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"` // local, development, production
	Port        int    `mapstructure:"port"`
	LogLevel    string `mapstructure:"log_level"` // debug, info, warn, error
}

// DatabaseConfig contains PostgreSQL connection settings
type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	SSLMode         string `mapstructure:"ssl_mode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"` // minutes
}

// KafkaConfig contains Kafka connection settings
type KafkaConfig struct {
	Brokers              []string `mapstructure:"brokers"`
	SASLUsername         string   `mapstructure:"sasl_username"`
	SASLPassword         string   `mapstructure:"sasl_password"`
	SASLMechanism        string   `mapstructure:"sasl_mechanism"` // SCRAM-SHA-512
	SecurityProtocol     string   `mapstructure:"security_protocol"` // SASL_SSL
	SkipVerify           bool     `mapstructure:"skip_verify"`
	ConsumerGroup        string   `mapstructure:"consumer_group"`
	KeycloakEventsTopic  string   `mapstructure:"keycloak_events_topic"`
	UserEventsTopic      string   `mapstructure:"user_events_topic"`
}

// KeycloakConfig contains Keycloak settings
type KeycloakConfig struct {
	URL          string `mapstructure:"url"`
	Realm        string `mapstructure:"realm"`
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
}

// DaprConfig contains Dapr settings
type DaprConfig struct {
	Enabled        bool   `mapstructure:"enabled"`
	HTTPPort       int    `mapstructure:"http_port"`
	GRPCPort       int    `mapstructure:"grpc_port"`
	PubSubName     string `mapstructure:"pubsub_name"`
	StateStoreName string `mapstructure:"state_store_name"`
}

// Load reads configuration from environment variables and config files
// Priority: Environment variables > config file > defaults
func Load() (*Config, error) {
	v := viper.New()

	// Set default values
	setDefaults(v)

	// Configure viper to read from environment variables
	v.SetEnvPrefix("CONTRACTS_BE")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Read from config file if it exists
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	v.AddConfigPath("/etc/users-be")

	// Try to read config file (it's okay if it doesn't exist)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Unmarshal configuration
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Override with environment variables for sensitive data
	overrideFromEnv(&config)

	// Validate configuration
	if err := validate(&config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &config, nil
}

// setDefaults sets default configuration values
func setDefaults(v *viper.Viper) {
	// App defaults
	v.SetDefault("app.name", "users-be")
	v.SetDefault("app.environment", "development")
	v.SetDefault("app.port", 8080)
	v.SetDefault("app.log_level", "info")

	// Database defaults
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "usersuser")
	v.SetDefault("database.database", "usersdb")
	v.SetDefault("database.ssl_mode", "disable")
	v.SetDefault("database.max_open_conns", 50)  // Increased from 25
	v.SetDefault("database.max_idle_conns", 10)  // Increased from 5
	v.SetDefault("database.conn_max_lifetime", 30)

	// Kafka defaults
	v.SetDefault("kafka.brokers", []string{"173.212.218.251:31691"})
	v.SetDefault("kafka.sasl_mechanism", "SCRAM-SHA-512")
	v.SetDefault("kafka.security_protocol", "SASL_SSL")
	v.SetDefault("kafka.skip_verify", true)
	v.SetDefault("kafka.consumer_group", "contracts-be-group")
	v.SetDefault("kafka.keycloak_events_topic", "keycloak-events")
	v.SetDefault("kafka.user_events_topic", "user-events")

	// Keycloak defaults
	v.SetDefault("keycloak.url", "http://173.212.218.251:30080")
	v.SetDefault("keycloak.realm", "skillsier")

	// Dapr defaults
	v.SetDefault("dapr.enabled", true)
	v.SetDefault("dapr.http_port", 3500)
	v.SetDefault("dapr.grpc_port", 50001)
	v.SetDefault("dapr.pubsub_name", "kafka-pubsub")
	v.SetDefault("dapr.state_store_name", "statestore")
}

// overrideFromEnv overrides sensitive configuration from environment variables
func overrideFromEnv(cfg *Config) {
	// Database password
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		cfg.Database.Password = pass
	}
	if pass := os.Getenv("POSTGRES_PASSWORD"); pass != "" {
		cfg.Database.Password = pass
	}

	// Kafka credentials
	if user := os.Getenv("KAFKA_USERNAME"); user != "" {
		cfg.Kafka.SASLUsername = user
	}
	if pass := os.Getenv("KAFKA_PASSWORD"); pass != "" {
		cfg.Kafka.SASLPassword = pass
	}

	// Keycloak credentials
	if clientID := os.Getenv("KEYCLOAK_CLIENT_ID"); clientID != "" {
		cfg.Keycloak.ClientID = clientID
	}
	if clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET"); clientSecret != "" {
		cfg.Keycloak.ClientSecret = clientSecret
	}
}

// validate checks if the configuration is valid
func validate(cfg *Config) error {
	if cfg.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if cfg.Database.User == "" {
		return fmt.Errorf("database user is required")
	}
	if cfg.Database.Database == "" {
		return fmt.Errorf("database name is required")
	}

	if len(cfg.Kafka.Brokers) == 0 {
		return fmt.Errorf("kafka brokers are required")
	}

	if cfg.Keycloak.URL == "" {
		return fmt.Errorf("keycloak URL is required")
	}
	if cfg.Keycloak.Realm == "" {
		return fmt.Errorf("keycloak realm is required")
	}

	return nil
}

// DSN returns the PostgreSQL connection string
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Database, d.SSLMode,
	)
}