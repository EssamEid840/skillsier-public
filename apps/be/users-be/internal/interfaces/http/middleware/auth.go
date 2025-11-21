package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	
	"skillsier.dev/pkg/auth"
)

// AuthMiddleware creates authentication middleware using pkg/auth
type AuthMiddleware struct {
	verifier auth.TokenVerifier
}

// NewAuthMiddleware creates a new authentication middleware
func NewAuthMiddleware(verifier auth.TokenVerifier) *AuthMiddleware {
	return &AuthMiddleware{
		verifier: verifier,
	}
}

// Authenticate verifies the JWT token and adds principal to context
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "Missing authorization header",
			})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		// Verify token using pkg/auth
		principal, err := m.verifier.VerifyToken(c.Request.Context(), token)
		if err != nil {
			m.handleAuthError(c, err)
			c.Abort()
			return
		}

		// Add principal to context using pkg/auth helper
		ctx := auth.WithPrincipal(c.Request.Context(), principal)
		c.Request = c.Request.WithContext(ctx)

		// Also set in Gin context for convenience
		c.Set("principal", principal)
		c.Set("user_id", principal.UserID)
		c.Set("subject", principal.Subject)

		c.Next()
	}
}

// RequireRoles requires any of the specified roles
func (m *AuthMiddleware) RequireRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		principal, exists := c.Get("principal")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "No authentication context",
			})
			c.Abort()
			return
		}

		p := principal.(*auth.Principal)
		if !p.HasAnyRole(roles...) {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "forbidden",
				"message": "Insufficient permissions",
				"required_roles": roles,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAllRoles requires all of the specified roles
func (m *AuthMiddleware) RequireAllRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		principal, exists := c.Get("principal")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "No authentication context",
			})
			c.Abort()
			return
		}

		p := principal.(*auth.Principal)
		if !p.HasAllRoles(roles...) {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "forbidden",
				"message": "Insufficient permissions",
				"required_roles": roles,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequirePermissions requires any of the specified permissions
func (m *AuthMiddleware) RequirePermissions(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		principal, exists := c.Get("principal")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "No authentication context",
			})
			c.Abort()
			return
		}

		p := principal.(*auth.Principal)
		if !p.HasAnyPermission(permissions...) {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "forbidden",
				"message": "Insufficient permissions",
				"required_permissions": permissions,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// handleAuthError converts pkg/auth errors to HTTP responses
func (m *AuthMiddleware) handleAuthError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, auth.ErrExpiredToken):
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "token_expired",
			"message": "Token has expired",
		})
	case errors.Is(err, auth.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid_token",
			"message": "Invalid token",
		})
	case errors.Is(err, auth.ErrInvalidIssuer):
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid_issuer",
			"message": "Invalid token issuer",
		})
	case errors.Is(err, auth.ErrInvalidAudience):
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid_audience",
			"message": "Invalid token audience",
		})
	default:
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Authentication failed",
		})
	}
}
