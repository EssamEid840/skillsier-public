package keycloak

import "skillsier.dev/pkg/auth"

// Config holds Keycloak-specific configuration
type Config struct {
	// Base auth config (issuer, audience, JWKS URL, etc.)
	Auth *auth.Config
	
	// Realm is the Keycloak realm name
	Realm string
	
	// ClientID is the Keycloak client ID
	ClientID string
	
	// ClientSecret is the Keycloak client secret (for service accounts)
	ClientSecret string
	
	// BaseURL is the Keycloak base URL (e.g., https://keycloak.example.com)
	BaseURL string
	
	// AdminUsername for Keycloak admin API (optional, for user management)
	AdminUsername string
	
	// AdminPassword for Keycloak admin API (optional, for user management)
	AdminPassword string
}

// NewConfig creates a Keycloak config from base URL and realm
func NewConfig(baseURL, realm, clientID, clientSecret string) *Config {
	issuer := baseURL + "/realms/" + realm
	jwksURL := issuer + "/protocol/openid-connect/certs"
	
	return &Config{
		Auth: &auth.Config{
			Issuer:            issuer,
			Audience:          clientID,
			JWKSURL:           jwksURL,
			AllowedAlgorithms: []string{"RS256"},
			CacheTTL:          auth.DefaultConfig().CacheTTL,
			ClockSkew:         auth.DefaultConfig().ClockSkew,
		},
		Realm:        realm,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
	}
}

// Validate checks if the Keycloak configuration is valid
func (c *Config) Validate() error {
	if c.Auth == nil {
		return auth.ErrInvalidConfig("auth config is required")
	}
	if err := c.Auth.Validate(); err != nil {
		return err
	}
	if c.Realm == "" {
		return auth.ErrInvalidConfig("realm is required")
	}
	if c.ClientID == "" {
		return auth.ErrInvalidConfig("client_id is required")
	}
	return nil
}