package ginx

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSConfig holds CORS configuration
type CORSConfig struct {
	// AllowOrigins is a list of allowed origins (e.g., ["https://app.skillsier.com"])
	// Use ["*"] to allow all origins (not recommended for production)
	AllowOrigins []string
	
	// AllowMethods is a list of allowed HTTP methods
	AllowMethods []string
	
	// AllowHeaders is a list of allowed headers
	AllowHeaders []string
	
	// ExposeHeaders is a list of headers exposed to the client
	ExposeHeaders []string
	
	// AllowCredentials indicates whether credentials are allowed
	AllowCredentials bool
	
	// MaxAge indicates how long preflight results can be cached (in seconds)
	MaxAge int
}

// DefaultCORSConfig returns a CORS config with sensible defaults
func DefaultCORSConfig() *CORSConfig {
	return &CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept",
			"Authorization",
			"X-Request-ID",
			"X-CSRF-Token",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"X-Request-ID",
		},
		AllowCredentials: true,
		MaxAge:           3600, // 1 hour
	}
}

// CORS middleware handles Cross-Origin Resource Sharing
func CORS(config *CORSConfig) gin.HandlerFunc {
	if config == nil {
		config = DefaultCORSConfig()
	}
	
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// Check if origin is allowed
		if origin != "" && isOriginAllowed(origin, config.AllowOrigins) {
			c.Header("Access-Control-Allow-Origin", origin)
			
			if config.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
			
			if len(config.ExposeHeaders) > 0 {
				c.Header("Access-Control-Expose-Headers", strings.Join(config.ExposeHeaders, ", "))
			}
		}
		
		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			if len(config.AllowMethods) > 0 {
				c.Header("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ", "))
			}
			
			if len(config.AllowHeaders) > 0 {
				c.Header("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))
			}
			
			if config.MaxAge > 0 {
				c.Header("Access-Control-Max-Age", string(rune(config.MaxAge)))
			}
			
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	}
}

// isOriginAllowed checks if an origin is in the allowed list
func isOriginAllowed(origin string, allowed []string) bool {
	for _, o := range allowed {
		if o == "*" {
			return true
		}
		if o == origin {
			return true
		}
		// Support wildcard subdomains (e.g., "*.skillsier.com")
		if strings.HasPrefix(o, "*.") {
			domain := o[2:]
			if strings.HasSuffix(origin, domain) {
				return true
			}
		}
	}
	return false
}