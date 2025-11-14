package outbox

import "time"

// Config holds outbox pattern configuration
type Config struct {
	// PollInterval is how often to poll for pending events
	PollInterval time.Duration
	
	// BatchSize is the maximum number of events to process in one batch
	BatchSize int
	
	// RetryBaseDelay is the base delay for exponential backoff
	RetryBaseDelay time.Duration
	
	// RetryMaxDelay is the maximum delay between retries
	RetryMaxDelay time.Duration
	
	// MaxAttempts is the maximum number of publish attempts
	MaxAttempts int
	
	// CleanupEnabled indicates if old published events should be cleaned up
	CleanupEnabled bool
	
	// CleanupInterval is how often to run cleanup
	CleanupInterval time.Duration
	
	// CleanupAge is how old published events must be to be cleaned up (in days)
	CleanupAge int
}

// DefaultConfig returns a configuration with sensible defaults
func DefaultConfig() *Config {
	return &Config{
		PollInterval:    5 * time.Second,
		BatchSize:       100,
		RetryBaseDelay:  1 * time.Second,
		RetryMaxDelay:   5 * time.Minute,
		MaxAttempts:     5,
		CleanupEnabled:  true,
		CleanupInterval: 24 * time.Hour,
		CleanupAge:      7, // 7 days
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.PollInterval <= 0 {
		return ErrInvalidConfig("poll_interval must be positive")
	}
	if c.BatchSize <= 0 {
		return ErrInvalidConfig("batch_size must be positive")
	}
	if c.RetryBaseDelay <= 0 {
		return ErrInvalidConfig("retry_base_delay must be positive")
	}
	if c.RetryMaxDelay < c.RetryBaseDelay {
		return ErrInvalidConfig("retry_max_delay must be >= retry_base_delay")
	}
	if c.MaxAttempts <= 0 {
		return ErrInvalidConfig("max_attempts must be positive")
	}
	if c.CleanupAge < 0 {
		return ErrInvalidConfig("cleanup_age must be non-negative")
	}
	return nil
}

// ErrInvalidConfig returns a configuration error
func ErrInvalidConfig(message string) error {
	return &ConfigError{
		Code:    "invalid_outbox_config",
		Message: message,
	}
}

// ConfigError represents a configuration error
type ConfigError struct {
	Code    string
	Message string
}

// Error implements the error interface
func (e *ConfigError) Error() string {
	return e.Code + ": " + e.Message
}