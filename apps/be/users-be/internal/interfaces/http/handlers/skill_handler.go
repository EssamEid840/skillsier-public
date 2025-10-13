// internal/interfaces/http/handlers/skill_handler.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "skillsier.dev/platform-shared/httpx"
    "users-be/internal/application/skill"
)

type SkillHandler struct {
    service *skill.Service
}

func NewSkillHandler(service *skill.Service) *SkillHandler {
    return &SkillHandler{service: service}
}

func (h *SkillHandler) AddSkill(c *gin.Context) {
    userID := c.Param("user_id")
    
    var req skill.AddSkillDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request", err)
        return
    }
    
    req.UserID = userID
    
    dto, err := h.service.AddSkill(c.Request.Context(), req)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to add skill", err)
        return
    }
    
    httpx.Success(c, http.StatusCreated, dto)
}

func (h *SkillHandler) GetSkills(c *gin.Context) {
    userID := c.Param("user_id")
    
    dtos, err := h.service.GetUserSkills(c.Request.Context(), userID)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get skills", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dtos)
}

func (h *SkillHandler) GetPrimarySkills(c *gin.Context) {
    userID := c.Param("user_id")
    
    dtos, err := h.service.GetPrimarySkills(c.Request.Context(), userID)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get primary skills", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dtos)
}

func (h *SkillHandler) UpdateSkill(c *gin.Context) {
    id := c.Param("id")
    
    var req skill.UpdateSkillDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request", err)
        return
    }
    
    dto, err := h.service.UpdateSkill(c.Request.Context(), id, req)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to update skill", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

func (h *SkillHandler) DeleteSkill(c *gin.Context) {
    id := c.Param("id")
    
    if err := h.service.DeleteSkill(c.Request.Context(), id); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to delete skill", err)
        return
    }
    
    httpx.Success(c, http.StatusNoContent, nil)
}

func (h *SkillHandler) ReorderSkills(c *gin.Context) {
    userID := c.Param("user_id")
    
    var req skill.ReorderSkillsDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request", err)
        return
    }
    
    if err := h.service.ReorderSkills(c.Request.Context(), userID, req); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to reorder skills", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Skills reordered successfully"})
}

func (h *SkillHandler) EndorseSkill(c *gin.Context) {
    id := c.Param("id")
    
    if err := h.service.EndorseSkill(c.Request.Context(), id); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to endorse skill", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Skill endorsed successfully"})
}