package http

import (
	"github.com/gin-gonic/gin"

	"jobs-be/internal/interfaces/http/handlers"
	"jobs-be/internal/interfaces/http/middleware"
)

func SetupRouter(jobHandler *handlers.JobHandler) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	api := router.Group("/api/v1")
	{
		jobs := api.Group("/jobs")
		{
			jobs.GET("", jobHandler.GetAll)
			jobs.GET("/:id", jobHandler.Get)
		}

		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			authenticated.POST("/jobs", jobHandler.Create)
			authenticated.PATCH("/jobs/:id", jobHandler.Update)
			authenticated.DELETE("/jobs/:id", jobHandler.Delete)
			authenticated.GET("/jobs/my-jobs", jobHandler.GetMyJobs)
		}
	}

	return router
}
