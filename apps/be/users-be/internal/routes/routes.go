package routes

import (
	"github.com/gin-gonic/gin"
	"users.go/m/internal/controller/users"
	users2 "users.go/m/internal/usecases/users"
)

func NewUsersRoutes(
	router *gin.Engine,
	createUserUseCase *users2.CreateUserUseCase,
) {
	userController := users.New(createUserUseCase)

	usersRoutes := router.Group("/users")
	{
		usersRoutes.POST("/", userController.Create)
	}
}
