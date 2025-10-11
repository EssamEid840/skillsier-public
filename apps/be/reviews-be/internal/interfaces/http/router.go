package http

import (
	"github.com/gin-gonic/gin"

	"reviews-be/internal/interfaces/http/handlers"
	"reviews-be/internal/interfaces/http/middleware"
)

func SetupRouter(reviewHandler *handlers.ReviewHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/reviews", reviewHandler.Create)
		api.GET("/reviews/received", reviewHandler.GetReceivedReviews)
	}

	return router
}