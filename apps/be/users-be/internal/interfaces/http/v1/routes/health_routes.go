package routes

import (
	"github.com/gin-gonic/gin"

	"skillsier.dev/apps/be/users-be/internal/interfaces/http/v1/handlers"
)

// RegisterHealthRoutes registers all health check routes
func RegisterHealthRoutes(router *gin.Engine, handler *handlers.HealthHandler) {
	// Health check endpoints (no versioning)
	router.GET("/health", handler.Health)  // Comprehensive health check
	router.GET("/ready", handler.Ready)    // Kubernetes readiness probe
	router.GET("/live", handler.Live)      // Kubernetes liveness probe
}