package idempotency

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// IdempotencyKeyHeader is the HTTP header name for idempotency key
	IdempotencyKeyHeader = "Idempotency-Key"
	
	// DefaultTTL is the default time-to-live for idempotency records
	DefaultTTL = 24 * time.Hour
)

// responseWriter wraps http.ResponseWriter to capture the response
type responseWriter struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func newResponseWriter(w gin.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		body:           &bytes.Buffer{},
		statusCode:     http.StatusOK,
	}
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Middleware provides idempotency for POST, PUT, and PATCH requests
func Middleware(repo Repository, ttl time.Duration) gin.HandlerFunc {
	if ttl == 0 {
		ttl = DefaultTTL
	}
	
	return func(c *gin.Context) {
		// Only apply to mutating methods
		if c.Request.Method != "POST" && c.Request.Method != "PUT" && c.Request.Method != "PATCH" {
			c.Next()
			return
		}
		
		// Get idempotency key from header
		key := c.GetHeader(IdempotencyKeyHeader)
		if key == "" {
			// No idempotency key, process normally
			c.Next()
			return
		}
		
		// Check if request with this key was already processed
		record, err := repo.Get(c.Request.Context(), key)
		if err == nil && record != nil && !record.IsExpired() {
			// Return cached response
			for k, v := range record.ResponseHeaders {
				c.Header(k, v)
			}
			c.Data(record.StatusCode, "application/json", record.ResponseBody)
			c.Abort()
			return
		}
		
		// Wrap response writer to capture response
		writer := newResponseWriter(c.Writer)
		c.Writer = writer
		
		// Process request
		c.Next()
		
		// Store response for future requests with same key
		headers := make(map[string]string)
		for k, v := range c.Writer.Header() {
			if len(v) > 0 {
				headers[k] = v[0]
			}
		}
		
		record = NewRecord(
			key,
			writer.statusCode,
			writer.body.Bytes(),
			headers,
			ttl,
		)
		
		// Store record (ignore errors to not fail the request)
		_ = repo.Create(c.Request.Context(), record)
	}
}