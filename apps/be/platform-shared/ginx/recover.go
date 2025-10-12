package ginx

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"skillsier.dev/platform-shared/httpx"
	"skillsier.dev/platform-shared/logging"
)

// Recovery middleware recovers from panics and logs the error
func Recovery(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get stack trace
				stack := string(debug.Stack())
				
				// Get logger from context or use provided logger
				reqLogger := logging.FromContext(c.Request.Context())
				if reqLogger == nil {
					reqLogger = logger
				}
				
				// Log the panic
				reqLogger.
					WithField(logging.FieldError, fmt.Sprintf("%v", err)).
					WithField(logging.FieldErrorStack, stack).
					WithField(logging.FieldRequestID, GetRequestID(c)).
					WithField(logging.FieldMethod, c.Request.Method).
					WithField(logging.FieldPath, c.Request.URL.Path).
					Error("Panic recovered")
				
				// Return error response
				httpx.WriteError(c.Writer, c.Request, httpx.ErrInternal.WithDetails("An unexpected error occurred"))
				
				// Abort the request
				c.Abort()
			}
		}()
		
		c.Next()
	}
}