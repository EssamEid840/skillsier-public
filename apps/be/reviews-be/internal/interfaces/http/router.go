package http

import (
	"reviews-be/internal/interfaces/http/handlers"
	"reviews-be/internal/interfaces/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	reviewHandler *handlers.ReviewHandler,
	healthHandler *handlers.HealthHandler,
) *gin.Engine {
	router := gin.Default()

	// CORS
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

	// Health endpoints
	router.GET("/health", healthHandler.Health)
	router.GET("/live", healthHandler.Liveness)
	router.GET("/ready", healthHandler.Readiness)

	// API routes
	v1 := router.Group("/api/v1")
	{
		reviews := v1.Group("/reviews")
		reviews.Use(middleware.AuthMiddleware())
		{
			reviews.POST("", reviewHandler.CreateReview)                     // POST /api/v1/reviews
			reviews.GET("/:id", reviewHandler.GetReview)                     // GET /api/v1/reviews/:id
			reviews.GET("/received", reviewHandler.GetReceivedReviews)       // GET /api/v1/reviews/received
			reviews.GET("/given", reviewHandler.GetGivenReviews)             // GET /api/v1/reviews/given
			reviews.GET("/user/:userId/rating", reviewHandler.GetUserRating) // GET /api/v1/reviews/user/:userId/rating
		}
	}

	return router
}