package ginx

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// OTel returns OpenTelemetry middleware for Gin
// This middleware automatically creates spans for each HTTP request
func OTel(serviceName string) gin.HandlerFunc {
	return otelgin.Middleware(serviceName)
}