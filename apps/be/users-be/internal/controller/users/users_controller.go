package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"users.go/m/internal/usecases/users"
)

type UserController struct {
	createUserUseCase *users.CreateUserUseCase
}

func New(
	createUserUseCase *users.CreateUserUseCase,
) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
	}
}

func (controller *UserController) Create(c *gin.Context) {
	var req users.CreateUserInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.createUserUseCase.Execute(
		c.Request.Context(),
		req,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
	return
}
