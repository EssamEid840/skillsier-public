package handlers

import (
	"net/http"
	"users-be/internal/application/freelancer"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FreelancerHandler struct {
	freelancerService *freelancer.Service
}

func NewFreelancerHandler(freelancerService *freelancer.Service) *FreelancerHandler {
	return &FreelancerHandler{freelancerService: freelancerService}
}

func (h *FreelancerHandler) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	result, err := h.freelancerService.GetProfile(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *FreelancerHandler) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	var dto freelancer.UpdateFreelancerProfileDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.freelancerService.UpdateProfile(c.Request.Context(), uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
