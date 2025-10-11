package handlers

import (
	"net/http"
	"users-be/internal/application/skill"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SkillHandler struct {
	skillService *skill.Service
}

func NewSkillHandler(skillService *skill.Service) *SkillHandler {
	return &SkillHandler{skillService: skillService}
}

// CreateSkill handles POST /users/profile/skills
func (h *SkillHandler) CreateSkill(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	var dto skill.CreateSkillDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.skillService.CreateSkill(c.Request.Context(), userID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == skill.ErrMaxSkillsExceeded {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetSkills handles GET /users/profile/skills
func (h *SkillHandler) GetSkills(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	skills, err := h.skillService.GetSkills(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"skills": skills})
}

// UpdateSkill handles PATCH /users/profile/skills/:id
func (h *SkillHandler) UpdateSkill(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	skillID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid skill ID"})
		return
	}

	var dto skill.UpdateSkillDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.skillService.UpdateSkill(c.Request.Context(), skillID, userID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == skill.ErrSkillNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteSkill handles DELETE /users/profile/skills/:id
func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	skillID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid skill ID"})
		return
	}

	if err := h.skillService.DeleteSkill(c.Request.Context(), skillID, userID); err != nil {
		statusCode := http.StatusInternalServerError
		if err == skill.ErrSkillNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "skill deleted successfully"})
}