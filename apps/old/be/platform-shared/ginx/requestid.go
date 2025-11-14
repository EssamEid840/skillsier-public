package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"skillsier.dev/platform-shared/httpx"
)

const (
	// RequestIDHeader is the HTTP header name for request ID
	RequestIDHeader = "X-Request-ID"
	
	// RequestIDContextKey is the context key for request ID
	RequestIDContextKey = "request_id"
)

// RequestID middleware generates or extracts a request ID for each request
// The request ID is used for tracking requests across services and logs
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get request ID from header
		requestID := c.GetHeader(RequestIDHeader)
		
		// Generate new ID if not present
		if requestID == "" {
			requestID = uuid.New().String()
		}
		
		// Set in context
		c.Set(RequestIDContextKey, requestID)
		
		// Add to request context for downstream use
		ctx := httpx.WithRequestID(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)
		
		// Add to response header
		c.Header(RequestIDHeader, requestID)
		
		c.Next()
	}
}

// GetRequestID retrieves the request ID from Gin context
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(RequestIDContextKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}