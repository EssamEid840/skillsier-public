package handlers

import (
	"net/http"
	"users-be/internal/application/freelancer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FreelancerHandler struct {
	service *freelancer.Service
}

func NewFreelancerHandler(service *freelancer.Service) *FreelancerHandler {
	return &FreelancerHandler{service: service}
}

func (h *FreelancerHandler) GetProfile(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	profile, err := h.service.GetProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *FreelancerHandler) UpdateProfile(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	var dto freelancer.UpdateFreelancerProfileDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.UpdateProfile(c.Request.Context(), userID, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}