package auth

import "time"

// Config holds authentication configuration
type Config struct {
	// Issuer is the expected issuer of the JWT token (e.g., Keycloak realm URL)
	Issuer string
	
	// Audience is the expected audience claim in the JWT
	Audience string
	
	// JWKSURL is the URL to fetch the JSON Web Key Set for token verification
	JWKSURL string
	
	// AllowedAlgorithms specifies which signing algorithms are allowed (e.g., ["RS256"])
	AllowedAlgorithms []string
	
	// CacheTTL is how long to cache JWKS keys before refreshing
	CacheTTL time.Duration
	
	// ClockSkew allows for time drift between systems
	ClockSkew time.Duration
}

// DefaultConfig returns a configuration with sensible defaults
func DefaultConfig() *Config {
	return &Config{
		AllowedAlgorithms: []string{"RS256"},
		CacheTTL:          10 * time.Minute,
		ClockSkew:         60 * time.Second,
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Issuer == "" {
		return ErrInvalidConfig("issuer is required")
	}
	if c.Audience == "" {
		return ErrInvalidConfig("audience is required")
	}
	if c.JWKSURL == "" {
		return ErrInvalidConfig("jwks_url is required")
	}
	if len(c.AllowedAlgorithms) == 0 {
		return ErrInvalidConfig("at least one allowed algorithm is required")
	}
	return nil
}