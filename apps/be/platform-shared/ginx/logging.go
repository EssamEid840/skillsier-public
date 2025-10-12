package ginx

import (
	"time"

	"github.com/gin-gonic/gin"
	"skillsier.dev/platform-shared/logging"
)

// Logging middleware logs all HTTP requests with structured logging
func Logging(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		
		// Add logger to context
		reqLogger := logger.
			WithRequestID(GetRequestID(c)).
			WithField(logging.FieldMethod, c.Request.Method).
			WithField(logging.FieldPath, path)
		
		ctx := logging.WithLogger(c.Request.Context(), reqLogger)
		c.Request = c.Request.WithContext(ctx)
		
		// Process request
		c.Next()
		
		// Calculate latency
		latency := time.Since(start)
		
		// Get client IP
		clientIP := c.ClientIP()
		
		// Get status code
		statusCode := c.Writer.Status()
		
		// Build log entry
		logEntry := reqLogger.
			WithField(logging.FieldStatusCode, statusCode).
			WithField(logging.FieldDuration, latency.Milliseconds()).
			WithField(logging.FieldClientIP, clientIP)
		
		if raw != "" {
			logEntry = logEntry.WithField("query", raw)
		}
		
		// Add user agent
		if userAgent := c.Request.UserAgent(); userAgent != "" {
			logEntry = logEntry.WithField(logging.FieldUserAgent, userAgent)
		}
		
		// Add error if present
		if len(c.Errors) > 0 {
			logEntry = logEntry.WithField(logging.FieldError, c.Errors.String())
		}
		
		// Log based on status code
		message := "HTTP request completed"
		switch {
		case statusCode >= 500:
			logEntry.Error(message)
		case statusCode >= 400:
			logEntry.Warn(message)
		default:
			logEntry.Info(message)
		}
	}
}