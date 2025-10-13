// internal/interfaces/http/handlers/profile_handler.go
package handlers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "skillsier.dev/platform-shared/httpx"
    
    "users-be/internal/application/profile"
    profileDomain "users-be/internal/domain/profile"
)

type ProfileHandler struct {
    service *profile.Service
}

func NewProfileHandler(service *profile.Service) *ProfileHandler {
    return &ProfileHandler{service: service}
}

// CREATE
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
    var req profile.CreateProfileDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    userID, _ := c.Get("user_id")
    req.UserID = userID.(string)
    
    dto, err := h.service.CreateProfile(c.Request.Context(), req)
    if err != nil {
        if err == profileDomain.ErrProfileAlreadyExists {
            httpx.Error(c, http.StatusConflict, "Profile already exists", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to create profile", err)
        return
    }
    
    httpx.Success(c, http.StatusCreated, dto)
}

// READ
func (h *ProfileHandler) GetProfile(c *gin.Context) {
    userID := c.Param("user_id")
    
    dto, err := h.service.GetProfile(c.Request.Context(), userID)
    if err != nil {
        if err == profileDomain.ErrProfileNotFound {
            httpx.Error(c, http.StatusNotFound, "Profile not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to get profile", err)
        return
    }
    
    // Track profile view
    _ = h.service.IncrementProfileViews(c.Request.Context(), userID)
    
    httpx.Success(c, http.StatusOK, dto)
}

func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
    userID, _ := c.Get("user_id")
    
    dto, err := h.service.GetProfile(c.Request.Context(), userID.(string))
    if err != nil {
        httpx.Error(c, http.StatusNotFound, "Profile not found", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

func (h *ProfileHandler) SearchProfiles(c *gin.Context) {
    query := c.Query("q")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
    
    filter := profileDomain.ListFilter{
        Page:     page,
        PageSize: pageSize,
    }
    
    // Parse additional filters
    if country := c.Query("country"); country != "" {
        filter.Country = &country
    }
    if city := c.Query("city"); city != "" {
        filter.City = &city
    }
    if minRate := c.Query("min_rate"); minRate != "" {
        if rate, err := strconv.ParseFloat(minRate, 64); err == nil {
            filter.MinRate = &rate
        }
    }
    if maxRate := c.Query("max_rate"); maxRate != "" {
        if rate, err := strconv.ParseFloat(maxRate, 64); err == nil {
            filter.MaxRate = &rate
        }
    }
    
    dtos, total, err := h.service.SearchProfiles(c.Request.Context(), query, filter)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to search profiles", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{
        "profiles":    dtos,
        "total":       total,
        "page":        page,
        "page_size":   pageSize,
        "total_pages": (int(total) + pageSize - 1) / pageSize,
    })
}

func (h *ProfileHandler) GetAvailableProfiles(c *gin.Context) {
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
    
    dtos, err := h.service.GetAvailableProfiles(c.Request.Context(), limit)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get available profiles", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dtos)
}

// UPDATE
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
    userID := c.Param("user_id")
    currentUserID, _ := c.Get("user_id")
    
    if currentUserID != userID {
        httpx.Error(c, http.StatusForbidden, "Cannot update other user's profile", nil)
        return
    }
    
    var req profile.UpdateProfileDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    dto, err := h.service.UpdateProfile(c.Request.Context(), userID, req)
    if err != nil {
        if err == profileDomain.ErrProfileNotFound {
            httpx.Error(c, http.StatusNotFound, "Profile not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to update profile", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

func (h *ProfileHandler) UpdateAvailability(c *gin.Context) {
    userID, _ := c.Get("user_id")
    
    var req struct {
        Status string `json:"status" binding:"required,oneof=available busy not_available"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    if err := h.service.UpdateAvailabilityStatus(c.Request.Context(), userID.(string), req.Status); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to update availability", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Availability status updated"})
}

// DELETE
func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
    userID := c.Param("user_id")
    currentUserID, _ := c.Get("user_id")
    
    if currentUserID != userID {
        httpx.Error(c, http.StatusForbidden, "Cannot delete other user's profile", nil)
        return
    }
    
    if err := h.service.DeleteProfile(c.Request.Context(), userID); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to delete profile", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}

// ANALYTICS
func (h *ProfileHandler) GetProfileStatistics(c *gin.Context) {
    stats, err := h.service.GetProfileStatistics(c.Request.Context())
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get statistics", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, stats)
}

func (h *ProfileHandler) GetIncompleteProfiles(c *gin.Context) {
    dtos, err := h.service.GetIncompleteProfiles(c.Request.Context())
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get incomplete profiles", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dtos)
}