package http

import (
	"users-be/internal/interfaces/http/handlers"
	"users-be/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all HTTP routes
func SetupRouter(
	userHandler *handlers.UserHandler,
	healthHandler *handlers.HealthHandler,
) *gin.Engine {
	// Create router
	router := gin.New()

	// Global middleware
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// Health check endpoints (no authentication required)
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/ready", healthHandler.ReadinessCheck)
	router.GET("/live", healthHandler.LivenessCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
			users.GET("/:id", userHandler.GetUser)
			users.GET("/keycloak/:keycloak_id", userHandler.GetUserByKeycloakID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return router
}