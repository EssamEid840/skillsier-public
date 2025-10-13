package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"users-be/internal/application/education"
)

type EducationHandler struct {
	educationService *education.Service
}

func NewEducationHandler(educationService *education.Service) *EducationHandler {
	return &EducationHandler{educationService: educationService}
}

func (h *EducationHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	var dto education.CreateEducationDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.educationService.CreateEducation(c.Request.Context(), uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *EducationHandler) GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	result, err := h.educationService.GetAllEducations(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *EducationHandler) Update(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	id, _ := uuid.Parse(c.Param("id"))
	var dto education.UpdateEducationDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.educationService.UpdateEducation(c.Request.Context(), id, uid, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *EducationHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := uuid.Parse(userID.(string))
	id, _ := uuid.Parse(c.Param("id"))
	if err := h.educationService.DeleteEducation(c.Request.Context(), id, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}