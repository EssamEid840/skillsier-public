package metrics

// Config holds metrics configuration
type Config struct {
	// Enabled enables metrics collection
	Enabled bool

	// Namespace is the metrics namespace (e.g., "skillsier")
	Namespace string

	// Subsystem is the metrics subsystem (e.g., "users_be")
	Subsystem string

	// Path is the HTTP path for metrics endpoint (default: "/metrics")
	Path string
}

// DefaultConfig returns default metrics configuration
func DefaultConfig() *Config {
	return &Config{
		Enabled:   true,
		Namespace: "skillsier",
		Subsystem: "",
		Path:      "/metrics",
	}
}