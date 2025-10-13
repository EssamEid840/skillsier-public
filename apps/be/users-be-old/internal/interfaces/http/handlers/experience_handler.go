package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"users-be/internal/application/experience"
)

type ExperienceHandler struct {
	experienceService *experience.Service
}

func NewExperienceHandler(experienceService *experience.Service) *ExperienceHandler {
	return &ExperienceHandler{experienceService: experienceService}
}

func (h *ExperienceHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	var dto experience.CreateWorkExperienceDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.experienceService.CreateExperience(c.Request.Context(), uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *ExperienceHandler) GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	result, err := h.experienceService.GetAllExperiences(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ExperienceHandler) Update(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experience ID"})
		return
	}

	var dto experience.UpdateWorkExperienceDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.experienceService.UpdateExperience(c.Request.Context(), id, uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ExperienceHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experience ID"})
		return
	}

	if err := h.experienceService.DeleteExperience(c.Request.Context(), id, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}