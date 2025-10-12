package logging

// Config holds logging configuration
type Config struct {
	// Level is the minimum log level (debug, info, warn, error, fatal, panic)
	Level string
	
	// Format is the log output format (json, pretty)
	Format string
	
	// ServiceName is added to all log entries
	ServiceName string
	
	// Output is the output destination (stdout, stderr, file path)
	Output string
}

// DefaultConfig returns a configuration with sensible defaults
func DefaultConfig(serviceName string) *Config {
	return &Config{
		Level:       "info",
		Format:      "json",
		ServiceName: serviceName,
		Output:      "stdout",
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	validLevels := map[string]bool{
		"debug": true, "info": true, "warn": true,
		"warning": true, "error": true, "fatal": true, "panic": true,
	}
	
	if !validLevels[c.Level] {
		return ErrInvalidLogLevel(c.Level)
	}
	
	validFormats := map[string]bool{"json": true, "pretty": true}
	if !validFormats[c.Format] {
		return ErrInvalidLogFormat(c.Format)
	}
	
	if c.ServiceName == "" {
		return ErrMissingServiceName
	}
	
	return nil
}

// ErrInvalidLogLevel returns an error for invalid log level
func ErrInvalidLogLevel(level string) error {
	return &LogError{
		Code:    "invalid_log_level",
		Message: "invalid log level: " + level,
	}
}

// ErrInvalidLogFormat returns an error for invalid log format
func ErrInvalidLogFormat(format string) error {
	return &LogError{
		Code:    "invalid_log_format",
		Message: "invalid log format: " + format,
	}
}

// ErrMissingServiceName is returned when service name is not provided
var ErrMissingServiceName = &LogError{
	Code:    "missing_service_name",
	Message: "service name is required",
}

// LogError represents a logging configuration error
type LogError struct {
	Code    string
	Message string
}

// Error implements the error interface
func (e *LogError) Error() string {
	return e.Code + ": " + e.Message
}