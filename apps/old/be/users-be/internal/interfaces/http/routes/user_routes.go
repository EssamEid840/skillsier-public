// internal/interfaces/http/routes/user_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	
	"users-be/internal/interfaces/http/handlers"
	"users-be/internal/interfaces/http/middleware"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(router *gin.Engine, handler *handlers.UserHandler, authMiddleware *middleware.AuthMiddleware, rbacMiddleware *middleware.RBACMiddleware) {
	// Public routes (no authentication required)
	public := router.Group("/api/v1")
	{
		public.GET("/users/username/:username", handler.GetUserByUsername)
		public.GET("/users/:id/public", handler.GetPublicProfile)
	}
	
	// Authenticated routes
	authenticated := router.Group("/api/v1")
	authenticated.Use(authMiddleware.RequireAuth())
	{
		// User CRUD operations
		users := authenticated.Group("/users")
		{
			users.POST("", handler.CreateUser)
			users.GET("", handler.ListUsers)
			users.GET("/search", handler.SearchUsers)
			users.GET("/:id", handler.GetUser)
			users.PUT("/:id", handler.UpdateUser)
			
			// Profile operations
			users.PUT("/:id/profile", handler.UpdateProfile)
			users.PUT("/:id/availability", handler.UpdateAvailability)
			users.PUT("/:id/settings", handler.UpdateSettings)
			
			// Verification operations
			users.POST("/:id/verify-email", handler.VerifyEmail)
			users.POST("/:id/verify-phone", handler.VerifyPhone)
			users.POST("/:id/verify-identity", handler.VerifyIdentity)
			
			// Activity tracking
			users.POST("/:id/login", handler.RecordLogin)
			users.PUT("/:id/last-seen", handler.UpdateLastSeen)
			users.PUT("/:id/online-status", handler.SetOnlineStatus)
		}
	}
	
	// Admin routes (require admin role)
	admin := router.Group("/api/v1/admin")
	admin.Use(authMiddleware.RequireAuth())
	admin.Use(rbacMiddleware.RequireRole("admin"))
	{
		adminUsers := admin.Group("/users")
		{
			// Statistics
			adminUsers.GET("/statistics", handler.GetUserStatistics)
			adminUsers.GET("/statistics/growth", handler.GetUserGrowthStats)
			
			// Moderation operations
			adminUsers.POST("/:id/suspend", handler.SuspendUser)
			adminUsers.POST("/:id/ban", handler.BanUser)
			adminUsers.POST("/:id/restore", handler.RestoreUser)
			adminUsers.POST("/:id/warn", handler.AddWarning)
			adminUsers.DELETE("/:id", handler.DeleteUser)
			
			// Badge operations
			adminUsers.POST("/:id/badges", handler.AssignBadge)
			adminUsers.DELETE("/:id/badges/:badge_type", handler.RemoveBadge)
			adminUsers.PUT("/:id/featured", handler.SetFeatured)
		}
	}
}