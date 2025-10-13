// internal/interfaces/http/handlers/user_handler.go
package handlers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "skillsier.dev/platform-shared/httpx"
    
    "users-be/internal/application/user"
    userDomain "users-be/internal/domain/user"
)

type UserHandler struct {
    service *user.Service
}

func NewUserHandler(service *user.Service) *UserHandler {
    return &UserHandler{service: service}
}

// ============================================================================
// CREATE OPERATIONS
// ============================================================================

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.CreateUserDTO true "User creation data"
// @Success 201 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 409 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req user.CreateUserDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    dto, err := h.service.CreateUser(c.Request.Context(), req)
    if err != nil {
        if err == userDomain.ErrEmailTaken {
            httpx.Error(c, http.StatusConflict, "Email already taken", err)
            return
        }
        if err == userDomain.ErrUsernameTaken {
            httpx.Error(c, http.StatusConflict, "Username already taken", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to create user", err)
        return
    }
    
    httpx.Success(c, http.StatusCreated, dto)
}

// CreateBulkUsers godoc
// @Summary Create multiple users
// @Description Bulk create user accounts
// @Tags users
// @Accept json
// @Produce json
// @Param users body []user.CreateUserDTO true "Array of user creation data"
// @Success 201 {object} []user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/bulk [post]
// @Security BearerAuth
func (h *UserHandler) CreateBulkUsers(c *gin.Context) {
    var req []user.CreateUserDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    dtos, err := h.service.CreateBulkUsers(c.Request.Context(), req)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to create users", err)
        return
    }
    
    httpx.Success(c, http.StatusCreated, dtos)
}

// ============================================================================
// READ OPERATIONS - Single User
// ============================================================================

// GetUser godoc
// @Summary Get user by ID
// @Description Get detailed user information by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/{id} [get]
// @Security BearerAuth
func (h *UserHandler) GetUser(c *gin.Context) {
    id := c.Param("id")
    
    dto, err := h.service.GetUserByID(c.Request.Context(), id)
    if err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to get user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

// GetCurrentUser godoc
// @Summary Get current authenticated user
// @Description Get the currently authenticated user's information
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} user.UserDTO
// @Failure 401 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/me [get]
// @Security BearerAuth
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        httpx.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
        return
    }
    
    dto, err := h.service.GetUserByID(c.Request.Context(), userID.(string))
    if err != nil {
        httpx.Error(c, http.StatusNotFound, "User not found", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

// GetUserByUsername godoc
// @Summary Get user by username
// @Description Get user information by username
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} user.PublicUserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/username/{username} [get]
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
    username := c.Param("username")
    
    dto, err := h.service.GetUserByUsername(c.Request.Context(), username)
    if err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to get user", err)
        return
    }
    
    // Return public info only
    publicDTO := user.SanitizeUserDTOForPublic(dto)
    httpx.Success(c, http.StatusOK, publicDTO)
}

// GetPublicUserProfile godoc
// @Summary Get public user profile
// @Description Get public user profile information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.PublicUserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id}/public [get]
func (h *UserHandler) GetPublicUserProfile(c *gin.Context) {
    id := c.Param("id")
    
    dto, err := h.service.GetUserByID(c.Request.Context(), id)
    if err != nil {
        httpx.Error(c, http.StatusNotFound, "User not found", err)
        return
    }
    
    publicDTO := user.SanitizeUserDTOForPublic(dto)
    httpx.Success(c, http.StatusOK, publicDTO)
}

// ============================================================================
// READ OPERATIONS - Lists & Search
// ============================================================================

// ListUsers godoc
// @Summary List users with filters
// @Description Get paginated list of users with optional filters
// @Tags users
// @Accept json
// @Produce json
// @Param filter query user.UserFilterDTO false "Filter parameters"
// @Success 200 {object} user.UserListResponseDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users [get]
// @Security BearerAuth
func (h *UserHandler) ListUsers(c *gin.Context) {
    var filterDTO user.UserFilterDTO
    if err := c.ShouldBindQuery(&filterDTO); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid filter parameters", err)
        return
    }
    
    // Set defaults
    if filterDTO.Page == 0 {
        filterDTO.Page = 1
    }
    if filterDTO.PageSize == 0 {
        filterDTO.PageSize = 20
    }
    
    filter := user.ToListFilter(filterDTO)
    
    response, err := h.service.ListUsers(c.Request.Context(), filter)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to list users", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, response)
}

// SearchUsers godoc
// @Summary Search users
// @Description Full-text search across users
// @Tags users
// @Accept json
// @Produce json
// @Param query query user.UserSearchQueryDTO true "Search parameters"
// @Success 200 {object} user.UserListResponseDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/search [get]
func (h *UserHandler) SearchUsers(c *gin.Context) {
    var searchDTO user.UserSearchQueryDTO
    if err := c.ShouldBindQuery(&searchDTO); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid search parameters", err)
        return
    }
    
    // Set defaults
    if searchDTO.Page == 0 {
        searchDTO.Page = 1
    }
    if searchDTO.PageSize == 0 {
        searchDTO.PageSize = 20
    }
    
    filter := user.ToSearchFilter(searchDTO)
    
    response, err := h.service.SearchUsers(c.Request.Context(), searchDTO.Query, filter)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to search users", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, response)
}

// GetTopRatedFreelancers godoc
// @Summary Get top rated freelancers
// @Description Get list of top rated freelancers
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} []user.PublicUserDTO
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/top-rated [get]
func (h *UserHandler) GetTopRatedFreelancers(c *gin.Context) {
    limitStr := c.DefaultQuery("limit", "10")
    limit, _ := strconv.Atoi(limitStr)
    
    if limit <= 0 || limit > 100 {
        limit = 10
    }
    
    users, err := h.service.GetTopRatedFreelancers(c.Request.Context(), limit)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get top rated freelancers", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, users)
}

// GetFeaturedUsers godoc
// @Summary Get featured users
// @Description Get list of featured users
// @Tags users
// @Accept json
// @Produce json
// @Param user_type query string false "User type filter (freelancer/client)"
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} []user.PublicUserDTO
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/featured [get]
func (h *UserHandler) GetFeaturedUsers(c *gin.Context) {
    userType := c.Query("user_type")
    limitStr := c.DefaultQuery("limit", "10")
    limit, _ := strconv.Atoi(limitStr)
    
    if limit <= 0 || limit > 100 {
        limit = 10
    }
    
    users, err := h.service.GetFeaturedUsers(c.Request.Context(), userDomain.UserType(userType), limit)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get featured users", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, users)
}

// GetOnlineUsers godoc
// @Summary Get online users
// @Description Get list of currently online users
// @Tags users
// @Accept json
// @Produce json
// @Param user_type query string false "User type filter"
// @Success 200 {object} []user.UserProfileSummaryDTO
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/online [get]
// @Security BearerAuth
func (h *UserHandler) GetOnlineUsers(c *gin.Context) {
    userType := c.Query("user_type")
    
    users, err := h.service.GetOnlineUsers(c.Request.Context(), userDomain.UserType(userType))
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get online users", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, users)
}

// ============================================================================
// UPDATE OPERATIONS
// ============================================================================

// UpdateUser godoc
// @Summary Update user
// @Description Update user information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body user.UpdateUserDTO true "User update data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 403 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/{id} [put]
// @Security BearerAuth
func (h *UserHandler) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    currentUserID, _ := c.Get("user_id")
    
    // Check if user can update (must be owner or admin)
    if currentUserID != id {
        // TODO: Check if user is admin
        httpx.Error(c, http.StatusForbidden, "Cannot update other user's information", nil)
        return
    }
    
    var req user.UpdateUserDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    dto, err := h.service.UpdateUser(c.Request.Context(), id, req)
    if err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        if err == userDomain.ErrUserBanned {
            httpx.Error(c, http.StatusForbidden, "User is banned", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to update user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

// UpdateCurrentUser godoc
// @Summary Update current user
// @Description Update the currently authenticated user's information
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.UpdateUserDTO true "User update data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 401 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/me [put]
// @Security BearerAuth
func (h *UserHandler) UpdateCurrentUser(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        httpx.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
        return
    }
    
    var req user.UpdateUserDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    dto, err := h.service.UpdateUser(c.Request.Context(), userID.(string), req)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to update user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

// UpdateUserStats godoc
// @Summary Update user statistics
// @Description Update user statistics (internal use)
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param stats body user.UpdateUserStatsDTO true "Stats update data"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id}/stats [put]
// @Security BearerAuth
func (h *UserHandler) UpdateUserStats(c *gin.Context) {
    id := c.Param("id")
    
    var req user.UpdateUserStatsDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    if err := h.service.UpdateUserStats(c.Request.Context(), id, req); err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to update stats", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Stats updated successfully"})
}

// UpdateOnlineStatus godoc
// @Summary Update online status
// @Description Update user's online/offline status
// @Tags users
// @Accept json
// @Produce json
// @Param status body map[string]bool true "Online status {\"is_online\": true}"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 401 {object} httpx.ErrorResponse
// @Router /users/me/online-status [put]
// @Security BearerAuth
func (h *UserHandler) UpdateOnlineStatus(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        httpx.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
        return
    }
    
    var req struct {
        IsOnline bool `json:"is_online"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    if err := h.service.UpdateOnlineStatus(c.Request.Context(), userID.(string), req.IsOnline); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to update online status", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Online status updated"})
}

// ============================================================================
// VERIFICATION OPERATIONS
// ============================================================================

// VerifyEmail godoc
// @Summary Verify email
// @Description Verify user's email address
// @Tags users
// @Accept json
// @Produce json
// @Param verification body user.VerifyEmailDTO true "Email verification data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/verify-email [post]
func (h *UserHandler) VerifyEmail(c *gin.Context) {
    var req user.VerifyEmailDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    // TODO: Validate token
    
    dto, err := h.service.VerifyEmail(c.Request.Context(), req.UserID)
    if err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to verify email", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

// VerifyPhone godoc
// @Summary Verify phone
// @Description Verify user's phone number
// @Tags users
// @Accept json
// @Produce json
// @Param verification body user.VerifyPhoneDTO true "Phone verification data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/verify-phone [post]
// @Security BearerAuth
func (h *UserHandler) VerifyPhone(c *gin.Context) {
    var req user.VerifyPhoneDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    // TODO: Validate code
    
    dto, err := h.service.VerifyPhone(c.Request.Context(), req.UserID)
    if err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to verify phone", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, dto)
}

// ============================================================================
// DELETE OPERATIONS
// ============================================================================

// DeleteUser godoc
// @Summary Delete user (soft delete)
// @Description Soft delete a user account
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 403 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/{id} [delete]
// @Security BearerAuth
func (h *UserHandler) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    currentUserID, _ := c.Get("user_id")
    
    // Check if user can delete (must be owner or admin)
    if currentUserID != id {
        // TODO: Check if user is admin
        httpx.Error(c, http.StatusForbidden, "Cannot delete other user's account", nil)
        return
    }
    
    if err := h.service.DeleteUser(c.Request.Context(), id); err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to delete user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// DeleteCurrentUser godoc
// @Summary Delete current user
// @Description Delete the currently authenticated user's account
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} httpx.SuccessResponse
// @Failure 401 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /users/me [delete]
// @Security BearerAuth
func (h *UserHandler) DeleteCurrentUser(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        httpx.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
        return
    }
    
    if err := h.service.DeleteUser(c.Request.Context(), userID.(string)); err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to delete user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Your account has been deleted"})
}

// ============================================================================
// ADMIN OPERATIONS
// ============================================================================

// SuspendUser godoc
// @Summary Suspend user (Admin only)
// @Description Suspend a user account
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param suspension body user.SuspendUserDTO true "Suspension data"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 403 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/suspend [post]
// @Security BearerAuth
func (h *UserHandler) SuspendUser(c *gin.Context) {
    id := c.Param("id")
    adminID, _ := c.Get("user_id")
    
    var req user.SuspendUserDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    if err := h.service.SuspendUser(c.Request.Context(), id, req.Reason, adminID.(string)); err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to suspend user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "User suspended successfully"})
}

// UnsuspendUser godoc
// @Summary Unsuspend user (Admin only)
// @Description Remove suspension from a user account
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/suspend [delete]
// @Security BearerAuth
func (h *UserHandler) UnsuspendUser(c *gin.Context) {
    id := c.Param("id")
    
    if err := h.service.UnsuspendUser(c.Request.Context(), id); err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to unsuspend user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "User unsuspended successfully"})
}

// BanUser godoc
// @Summary Ban user (Admin only)
// @Description Ban a user account
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param ban body user.BanUserDTO true "Ban data"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/ban [post]
// @Security BearerAuth
func (h *UserHandler) BanUser(c *gin.Context) {
    id := c.Param("id")
    adminID, _ := c.Get("user_id")
    
    var req user.BanUserDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    if err := h.service.BanUser(c.Request.Context(), id, req.Reason, adminID.(string)); err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to ban user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "User banned successfully"})
}

// UnbanUser godoc
// @Summary Unban user (Admin only)
// @Description Remove ban from a user account
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/ban [delete]
// @Security BearerAuth
func (h *UserHandler) UnbanUser(c *gin.Context) {
    id := c.Param("id")
    
    if err := h.service.UnbanUser(c.Request.Context(), id); err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to unban user", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "User unbanned successfully"})
}

// SetFeatured godoc
// @Summary Set user as featured (Admin only)
// @Description Mark/unmark user as featured
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param featured body map[string]bool true "Featured status {\"featured\": true}"
// @Success 200 {object} httpx.SuccessResponse
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/featured [put]
// @Security BearerAuth
func (h *UserHandler) SetFeatured(c *gin.Context) {
    id := c.Param("id")
    
    var req struct {
        Featured bool `json:"featured"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }
    
    if err := h.service.MarkAsFeatured(c.Request.Context(), id, req.Featured); err != nil {
        if err == userDomain.ErrUserNotFound {
            httpx.Error(c, http.StatusNotFound, "User not found", err)
            return
        }
        httpx.Error(c, http.StatusInternalServerError, "Failed to update featured status", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{"message": "Featured status updated"})
}

// ============================================================================
// ANALYTICS & STATISTICS
// ============================================================================

// GetUserStatistics godoc
// @Summary Get platform user statistics (Admin only)
// @Description Get aggregated user statistics
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} user.UserStatisticsDTO
// @Failure 500 {object} httpx.ErrorResponse
// @Router /admin/users/statistics [get]
// @Security BearerAuth
func (h *UserHandler) GetUserStatistics(c *gin.Context) {
    stats, err := h.service.GetUserStatistics(c.Request.Context())
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to get statistics", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, stats)
}

// GetUserTrustScore godoc
// @Summary Get user trust score
// @Description Get calculated trust score for a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserTrustScoreDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id}/trust-score [get]
// @Security BearerAuth
func (h *UserHandler) GetUserTrustScore(c *gin.Context) {
    id := c.Param("id")
    
    trustLevel, err := h.service.CheckUserTrustLevel(c.Request.Context(), id)
    if err != nil {
        httpx.Error(c, http.StatusNotFound, "User not found", err)
        return
    }
    
    httpx.Success(c, http.StatusOK, gin.H{
        "user_id": id,
        "trust_level": trustLevel,
    })
}

// ============================================================================
// VALIDATION HELPERS
// ============================================================================

// CheckUsernameAvailability godoc
// @Summary Check username availability
// @Description Check if a username is available
// @Tags users
// @Accept json
// @Produce json
// @Param username query string true "Username to check"
// @Success 200 {object} user.UsernameAvailabilityResponseDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users/check-username [get]
func (h *UserHandler) CheckUsernameAvailability(c *gin.Context) {
    username := c.Query("username")
    if username == "" {
        httpx.Error(c, http.StatusBadRequest, "Username is required", nil)
        return
    }
    
    exists, err := h.service.CheckUsernameExists(c.Request.Context(), username)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to check username", err)
        return
    }
    
    response := user.UsernameAvailabilityResponseDTO{
        Available: !exists,
    }
    
    if exists {
        response.Message = "Username is already taken"
    } else {
        response.Message = "Username is available"
    }
    
    httpx.Success(c, http.StatusOK, response)
}

// CheckEmailAvailability godoc
// @Summary Check email availability
// @Description Check if an email is available
// @Tags users
// @Accept json
// @Produce json
// @Param email query string true "Email to check"
// @Success 200 {object} user.EmailAvailabilityResponseDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users/check-email [get]
func (h *UserHandler) CheckEmailAvailability(c *gin.Context) {
    email := c.Query("email")
    if email == "" {
        httpx.Error(c, http.StatusBadRequest, "Email is required", nil)
        return
    }
    
    exists, err := h.service.CheckEmailExists(c.Request.Context(), email)
    if err != nil {
        httpx.Error(c, http.StatusInternalServerError, "Failed to check email", err)
        return
    }
    
    response := user.EmailAvailabilityResponseDTO{
        Available: !exists,
    }
    
    if exists {
        response.Message = "Email is already registered"
    } else {
        response.Message = "Email is available"
    }
    
    httpx.Success(c, http.StatusOK, response)
}