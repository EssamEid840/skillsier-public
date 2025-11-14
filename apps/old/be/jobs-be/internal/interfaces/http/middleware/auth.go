package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT tokens from Keycloak
// For production, implement proper JWT validation with Keycloak public key
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		// Extract token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		token := parts[1]
		
		// TODO: Validate JWT token with Keycloak
		// For now, extract user_id from token claims
		// In production, use a proper JWT library and validate with Keycloak public key
		
		// Mock user_id extraction - REPLACE WITH REAL JWT VALIDATION
		userID := "mock-user-id" // Extract from validated JWT claims
		
		c.Set("user_id", userID)
		c.Next()
	}
}
