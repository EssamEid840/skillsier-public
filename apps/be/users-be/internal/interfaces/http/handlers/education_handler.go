// internal/interfaces/http/handlers/education_handler.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "skillsier.dev/platform-shared/httpx"
    "users-be/internal/application/education"
)

type EducationHandler struct {
    service *education.Service
}

func NewEducationHandler(service *education.Service) *EducationHandler {
    return &EducationHandler{service: service}
}

func (h *EducationHandler) AddEducation(c *gin.Context) {
    userID := c.Param("user_id")
    
    var req education.AddEducationDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request", err)
        return
    }
    
    req.UserID = userID
    
    dto, err := h.service.AddEducation(c.Request.Context(), req)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to add education", err)
        return
    }
    
    httpx.Success(c, http.StatusCreated, dto)
}

func (h *EducationHandler) GetEducations(c *gin.Context) {
    userID := c.Param("user_id")
    
    dtos, err := h.service.GetUserEducations(c.Request.Context(), userID)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get educations", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dtos)
}

func (h *EducationHandler) GetHighestDegree(c *gin.Context) {
    userID := c.Param("user_id")
    
    dto, err := h.service.GetHighestDegree(c.Request.Context(), userID)
    if err != nil {
        httpx.Error(c, http.StatusNotFound, "No education found", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

func (h *EducationHandler) UpdateEducation(c *gin.Context) {
    id := c.Param("id")
    
    var req education.UpdateEducationDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request", err)
        return
    }
    
    dto, err := h.service.UpdateEducation(c.Request.Context(), id, req)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to update education", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

func (h *EducationHandler) DeleteEducation(c *gin.Context) {
    id := c.Param("id")
    
    if err := h.service.DeleteEducation(c.Request.Context(), id); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to delete education", err)
        return
    }
    
    httpx.Success(c, http.StatusNoContent, nil)
}

func (h *EducationHandler) ReorderEducations(c *gin.Context) {
    userID := c.Param("user_id")
    
    var req education.ReorderEducationsDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request", err)
        return
    }
    
    if err := h.service.ReorderEducations(c.Request.Context(), userID, req); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to reorder educations", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Educations reordered successfully"})
}

func (h *EducationHandler) VerifyEducation(c *gin.Context) {
    id := c.Param("id")
    adminID, _ := c.Get("user_id")
    
    if err := h.service.VerifyEducation(c.Request.Context(), id, adminID.(string)); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to verify education", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Education verified successfully"})
}