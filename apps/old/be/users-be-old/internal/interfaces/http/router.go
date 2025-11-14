package http

import (
	"github.com/gin-gonic/gin"

	"users-be/internal/interfaces/http/handlers"
	"users-be/internal/interfaces/http/middleware"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	skillHandler *handlers.SkillHandler,
	experienceHandler *handlers.ExperienceHandler,
	educationHandler *handlers.EducationHandler,
	certificationHandler *handlers.CertificationHandler,
	portfolioHandler *handlers.PortfolioHandler,
	freelancerHandler *handlers.FreelancerHandler,
	clientHandler *handlers.ClientHandler,
) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			users.GET("/profile", userHandler.GetProfile)
			users.PATCH("/profile", userHandler.UpdateProfile)

			profile := users.Group("/profile")
			{
				// Skills
				profile.GET("/skills", skillHandler.GetAll)
				profile.POST("/skills", skillHandler.Create)
				profile.PATCH("/skills/:id", skillHandler.Update)
				profile.DELETE("/skills/:id", skillHandler.Delete)

				// Experience
				profile.GET("/experience", experienceHandler.GetAll)
				profile.POST("/experience", experienceHandler.Create)
				profile.PATCH("/experience/:id", experienceHandler.Update)
				profile.DELETE("/experience/:id", experienceHandler.Delete)

				// Education
				profile.GET("/education", educationHandler.GetAll)
				profile.POST("/education", educationHandler.Create)
				profile.PATCH("/education/:id", educationHandler.Update)
				profile.DELETE("/education/:id", educationHandler.Delete)

				// Certifications
				profile.GET("/certifications", certificationHandler.GetAll)
				profile.POST("/certifications", certificationHandler.Create)
				profile.PATCH("/certifications/:id", certificationHandler.Update)
				profile.DELETE("/certifications/:id", certificationHandler.Delete)

				// Portfolio
				profile.GET("/portfolio", portfolioHandler.GetAll)
				profile.POST("/portfolio", portfolioHandler.Create)
				profile.PATCH("/portfolio/:id", portfolioHandler.Update)
				profile.DELETE("/portfolio/:id", portfolioHandler.Delete)
				profile.POST("/portfolio/:id/images", portfolioHandler.UploadImage)
			}

			// Freelancer profile
			freelancer := users.Group("/freelancer")
			{
				freelancer.GET("/profile", freelancerHandler.GetProfile)
				freelancer.PATCH("/profile", freelancerHandler.UpdateProfile)
			}

			// Client profile
			client := users.Group("/client")
			{
				client.GET("/profile", clientHandler.GetProfile)
				client.PATCH("/profile", clientHandler.UpdateProfile)
			}
		}
	}

	return router
}