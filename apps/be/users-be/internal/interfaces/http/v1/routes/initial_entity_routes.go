package routes

import (
	"github.com/gin-gonic/gin"

	"skillsier.dev/apps/be/users-be/internal/interfaces/http/v1/handlers"
)

// RegisterInitialEntityRoutes registers all InitialEntity routes
func RegisterInitialEntityRoutesNoAuth(router *gin.RouterGroup, handler *handlers.InitialEntityHandler) {
	entities := router.Group("/initial-entities")
	{
		// CRUD operations
		entities.POST("", handler.Create)           // Create new entity
		entities.GET("", handler.List)              // List entities with pagination
		entities.GET("/:id", handler.Get)           // Get entity by ID
		entities.PUT("/:id", handler.Update)        // Update entity
		entities.DELETE("/:id", handler.Delete)     // Soft delete entity
		
		// Additional operations
		entities.POST("/:id/restore", handler.Restore)              // Restore deleted entity
		entities.GET("/owner/:owner_id", handler.GetByOwner)        // Get entities by owner
	}
}