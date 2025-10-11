package http

import (
	"github.com/gin-gonic/gin"

	"contracts-be/internal/interfaces/http/handlers"
	"contracts-be/internal/interfaces/http/middleware"
)

func SetupRouter(contractHandler *handlers.ContractHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/contracts", contractHandler.GetMyContracts)
		api.POST("/contracts/:id/milestones/:milestone_id/submit", contractHandler.SubmitMilestone)
		api.POST("/contracts/:id/milestones/:milestone_id/approve", contractHandler.ApproveMilestone)
		api.POST("/contracts/:id/complete", contractHandler.CompleteContract)
	}

	return router
}