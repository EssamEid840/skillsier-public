package ginx

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"skillsier.dev/platform-shared/httpx"
)

// Timeout middleware enforces a timeout on request processing
func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		
		// Replace request context
		c.Request = c.Request.WithContext(ctx)
		
		// Channel to signal completion
		done := make(chan struct{})
		
		// Process request in goroutine
		go func() {
			c.Next()
			close(done)
		}()
		
		// Wait for completion or timeout
		select {
		case <-done:
			// Request completed successfully
			return
		case <-ctx.Done():
			// Timeout occurred
			if ctx.Err() == context.DeadlineExceeded {
				httpx.WriteError(c.Writer, c.Request, 
					httpx.NewHTTPError(http.StatusRequestTimeout, "request_timeout", "Request processing timeout"))
				c.Abort()
			}
		}
	}
}