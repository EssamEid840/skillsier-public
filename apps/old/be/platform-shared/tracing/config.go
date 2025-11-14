package tracing

// Config holds tracing configuration
type Config struct {
	// Enabled enables tracing
	Enabled bool

	// Endpoint is the Jaeger collector endpoint
	Endpoint string

	// ServiceName is the service name
	ServiceName string

	// SamplingRate is the sampling rate (0.0 to 1.0)
	SamplingRate float64
}

// DefaultConfig returns default tracing configuration
func DefaultConfig() *Config {
	return &Config{
		Enabled:      true,
		Endpoint:     "http://localhost:14268/api/traces",
		ServiceName:  "",
		SamplingRate: 1.0,
	}
}