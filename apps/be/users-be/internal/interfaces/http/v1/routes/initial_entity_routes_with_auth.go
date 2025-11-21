package routes

import (
	"github.com/gin-gonic/gin"
	
	"skillsier.dev/apps/be/users-be/internal/interfaces/http/middleware"
	"skillsier.dev/apps/be/users-be/internal/interfaces/http/v1/handlers"
)

// RegisterInitialEntityRoutes registers InitialEntity routes with optional auth
func RegisterInitialEntityRoutes(
	router *gin.RouterGroup,
	handler *handlers.InitialEntityHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Public routes (no auth)
	publicRoutes := router.Group("/initial-entities")
	{
		publicRoutes.GET("/:id", handler.Get)
		publicRoutes.GET("", handler.List)
	}
	
	// Protected routes (requires authentication)
	protectedRoutes := router.Group("/initial-entities")
	protectedRoutes.Use(authMiddleware.Authenticate())
	{
		// Any authenticated user can create
		protectedRoutes.POST("", handler.Create)
		
		// Only owner or admin can update/delete
		protectedRoutes.PUT("/:id", handler.Update)
		protectedRoutes.DELETE("/:id", handler.Delete)
	}
	
	// Admin only routes
	adminRoutes := router.Group("/initial-entities/admin")
	adminRoutes.Use(authMiddleware.Authenticate())
	adminRoutes.Use(authMiddleware.RequireRoles("admin"))
	{
		// Admin endpoints here
	}
	
	// Role-specific examples
	freelancerRoutes := router.Group("/initial-entities/freelancer")
	freelancerRoutes.Use(authMiddleware.Authenticate())
	freelancerRoutes.Use(authMiddleware.RequireRoles("freelancer", "premium_freelancer"))
	{
		// Freelancer-specific endpoints
	}
	
	// Permission-based examples
	permissionRoutes := router.Group("/initial-entities/manage")
	permissionRoutes.Use(authMiddleware.Authenticate())
	permissionRoutes.Use(authMiddleware.RequirePermissions("initial_entities:manage"))
	{
		// Permission-based endpoints
	}
}