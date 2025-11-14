package http

import (
	"github.com/gin-gonic/gin"
	"proposals-be/internal/interfaces/http/handlers"
	"proposals-be/internal/interfaces/http/middleware"
)

func SetupRouter(proposalHandler *handlers.ProposalHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/proposals", proposalHandler.Create)
		api.GET("/proposals/my-proposals", proposalHandler.GetMyProposals)
		api.PATCH("/proposals/:id/withdraw", proposalHandler.Withdraw)
	}

	return router
}