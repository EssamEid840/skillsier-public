package http

import (
	"jobs-be/internal/interfaces/http/handlers"
	"jobs-be/internal/interfaces/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	jobHandler *handlers.JobHandler,
	healthHandler *handlers.HealthHandler,
) *gin.Engine {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Health endpoints (no auth required)
	router.GET("/health", healthHandler.Health)
	router.GET("/live", healthHandler.Liveness)
	router.GET("/ready", healthHandler.Readiness)

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Jobs routes (with authentication)
		jobs := v1.Group("/jobs")
		jobs.Use(middleware.AuthMiddleware())
		{
			jobs.GET("", jobHandler.ListJobs)           // GET /api/v1/jobs
			jobs.POST("", jobHandler.CreateJob)         // POST /api/v1/jobs
			jobs.GET("/my-jobs", jobHandler.GetMyJobs)  // GET /api/v1/jobs/my-jobs
			jobs.GET("/:id", jobHandler.GetJob)         // GET /api/v1/jobs/:id
			jobs.PATCH("/:id", jobHandler.UpdateJob)    // PATCH /api/v1/jobs/:id
			jobs.DELETE("/:id", jobHandler.DeleteJob)   // DELETE /api/v1/jobs/:id
		}
	}

	return router
}
