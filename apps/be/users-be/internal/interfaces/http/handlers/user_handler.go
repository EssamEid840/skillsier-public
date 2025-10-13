package handlers
// internal/interfaces/http/handlers/user_handler.go
package handlers

import (
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	
	"users-be/internal/application/user"
	userDomain "users-be/internal/domain/user"
	"skillsier.dev/platform-shared/httpx"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	service        *user.Service
	commandHandler *user.CommandHandler
	queryHandler   *user.QueryHandler
}

// NewUserHandler creates a new user handler
func NewUserHandler(service *user.Service) *UserHandler {
	return &UserHandler{
		service:        service,
		commandHandler: user.NewCommandHandler(service),
		queryHandler:   user.NewQueryHandler(service),
	}
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
// @Param user body user.CreateUserDTO true "User data"
// @Success 201 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 409 {object} httpx.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var dto user.CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	// Validate DTO
	if err := user.ValidateCreateUserDTO(&dto); err != nil {
		httpx.ValidationError(c, err)
		return
	}
	
	// Create command
	cmd := user.CreateUserCommand{
		KeycloakID:       dto.KeycloakID,
		Username:         dto.Username,
		Email:            dto.Email,
		FirstName:        dto.FirstName,
		LastName:         dto.LastName,
		UserType:         dto.UserType,
		PhoneNumber:      dto.PhoneNumber,
		PhoneCountryCode: dto.PhoneCountryCode,
		City:             dto.City,
		Country:          dto.Country,
		CountryCode:      dto.CountryCode,
		Timezone:         dto.Timezone,
		Bio:              dto.Bio,
		Tagline:          dto.Tagline,
	}
	
	result, err := h.commandHandler.HandleCreateUser(c.Request.Context(), cmd)
	if err != nil {
		if userDomain.IsConflictError(err) {
			httpx.Error(c, http.StatusConflict, "User already exists", err)
			return
		}
		if userDomain.IsValidationError(err) {
			httpx.ValidationError(c, err)
			return
		}
		httpx.Error(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}
	
	httpx.Success(c, http.StatusCreated, result)
}

// ============================================================================
// READ OPERATIONS
// ============================================================================

// GetUser godoc
// @Summary Get user by ID
// @Description Get user information by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	
	query := user.GetUserByIDQuery{UserID: id}
	result, err := h.queryHandler.HandleGetUserByID(c.Request.Context(), query)
	if err != nil {
		if userDomain.IsNotFoundError(err) {
			httpx.Error(c, http.StatusNotFound, "User not found", err)
			return
		}
		httpx.Error(c, http.StatusInternalServerError, "Failed to get user", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
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
	
	query := user.GetUserByUsernameQuery{Username: username}
	result, err := h.queryHandler.HandleGetUserByUsername(c.Request.Context(), query)
	if err != nil {
		if userDomain.IsNotFoundError(err) {
			httpx.Error(c, http.StatusNotFound, "User not found", err)
			return
		}
		httpx.Error(c, http.StatusInternalServerError, "Failed to get user", err)
		return
	}
	
	// Return public info only
	publicDTO := user.ToPublicUserDTO(result)
	httpx.Success(c, http.StatusOK, publicDTO)
}

// GetPublicProfile godoc
// @Summary Get public user profile
// @Description Get public user profile information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.PublicUserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id}/public [get]
func (h *UserHandler) GetPublicProfile(c *gin.Context) {
	id := c.Param("id")
	
	query := user.GetPublicProfileQuery{UserID: id}
	result, err := h.queryHandler.HandleGetPublicProfile(c.Request.Context(), query)
	if err != nil {
		httpx.Error(c, http.StatusNotFound, "User not found", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// ListUsers godoc
// @Summary List users
// @Description List users with filtering and pagination
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(20)
// @Param user_type query string false "User type filter"
// @Param status query string false "Status filter"
// @Param country query string false "Country filter"
// @Param min_rating query number false "Minimum rating"
// @Param is_featured query bool false "Featured only"
// @Param is_online query bool false "Online only"
// @Param sort_by query string false "Sort field" default(created_at)
// @Param sort_order query string false "Sort order (asc/desc)" default(desc)
// @Success 200 {object} user.UserListResponseDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	filter := h.parseListFilter(c)
	
	if err := filter.Validate(); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid filter parameters", err)
		return
	}
	
	query := user.ListUsersQuery{Filter: filter}
	result, err := h.queryHandler.HandleListUsers(c.Request.Context(), query)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to list users", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// SearchUsers godoc
// @Summary Search users
// @Description Search users by query string
// @Tags users
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(20)
// @Success 200 {object} user.UserListResponseDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users/search [get]
func (h *UserHandler) SearchUsers(c *gin.Context) {
	searchQuery := c.Query("q")
	if searchQuery == "" {
		httpx.Error(c, http.StatusBadRequest, "Search query is required", nil)
		return
	}
	
	filter := h.parseListFilter(c)
	if err := filter.Validate(); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid filter parameters", err)
		return
	}
	
	query := user.SearchUsersQuery{
		SearchQuery: searchQuery,
		Filter:      filter,
	}
	
	result, err := h.queryHandler.HandleSearchUsers(c.Request.Context(), query)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to search users", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
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
// @Param user body user.UpdateUserDTO true "User data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	
	var dto user.UpdateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	// Validate DTO
	if err := user.ValidateUpdateUserDTO(&dto); err != nil {
		httpx.ValidationError(c, err)
		return
	}
	
	cmd := user.UpdateUserCommand{
		UserID:            id,
		FirstName:         dto.FirstName,
		LastName:          dto.LastName,
		DisplayName:       dto.DisplayName,
		Bio:               dto.Bio,
		Tagline:           dto.Tagline,
		Title:             dto.Title,
		Website:           dto.Website,
		ProfilePictureURL: dto.ProfilePictureURL,
		CoverImageURL:     dto.CoverImageURL,
		PhoneNumber:       dto.PhoneNumber,
		PhoneCountryCode:  dto.PhoneCountryCode,
		City:              dto.City,
		Country:           dto.Country,
		CountryCode:       dto.CountryCode,
		Timezone:          dto.Timezone,
	}
	
	result, err := h.commandHandler.HandleUpdateUser(c.Request.Context(), cmd)
	if err != nil {
		if userDomain.IsNotFoundError(err) {
			httpx.Error(c, http.StatusNotFound, "User not found", err)
			return
		}
		if userDomain.IsValidationError(err) {
			httpx.ValidationError(c, err)
			return
		}
		httpx.Error(c, http.StatusInternalServerError, "Failed to update user", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update user profile information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param profile body user.UpdateProfileDTO true "Profile data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users/{id}/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	
	var dto user.UpdateProfileDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	if err := user.ValidateUpdateProfileDTO(&dto); err != nil {
		httpx.ValidationError(c, err)
		return
	}
	
	cmd := user.UpdateProfileCommand{
		UserID:  id,
		Bio:     dto.Bio,
		Tagline: dto.Tagline,
		Title:   dto.Title,
	}
	
	result, err := h.commandHandler.HandleUpdateProfile(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to update profile", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// UpdateAvailability godoc
// @Summary Update user availability
// @Description Update user availability status
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param availability body user.UpdateAvailabilityDTO true "Availability data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users/{id}/availability [put]
func (h *UserHandler) UpdateAvailability(c *gin.Context) {
	id := c.Param("id")
	
	var dto user.UpdateAvailabilityDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	if err := user.ValidateUpdateAvailabilityDTO(&dto); err != nil {
		httpx.ValidationError(c, err)
		return
	}
	
	cmd := user.UpdateAvailabilityCommand{
		UserID:       id,
		Status:       dto.Status,
		HoursPerWeek: dto.HoursPerWeek,
	}
	
	result, err := h.commandHandler.HandleUpdateAvailability(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to update availability", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// UpdateSettings godoc
// @Summary Update user settings
// @Description Update user settings and preferences
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param settings body user.UpdateSettingsDTO true "Settings data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users/{id}/settings [put]
func (h *UserHandler) UpdateSettings(c *gin.Context) {
	id := c.Param("id")
	
	var dto user.UpdateSettingsDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	if err := user.ValidateUpdateSettingsDTO(&dto); err != nil {
		httpx.ValidationError(c, err)
		return
	}
	
	cmd := user.UpdateSettingsCommand{
		UserID:            id,
		ProfileVisibility: dto.ProfileVisibility,
		ShowEmail:         dto.ShowEmail,
		ShowPhone:         dto.ShowPhone,
		ShowLocation:      dto.ShowLocation,
		SearchableProfile: dto.SearchableProfile,
		AcceptingWork:     dto.AcceptingWork,
	}
	
	result, err := h.commandHandler.HandleUpdateSettings(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to update settings", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// ============================================================================
// VERIFICATION OPERATIONS
// ============================================================================

// VerifyEmail godoc
// @Summary Verify email
// @Description Mark user email as verified
// @Tags users,verification
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id}/verify-email [post]
func (h *UserHandler) VerifyEmail(c *gin.Context) {
	id := c.Param("id")
	
	cmd := user.VerifyEmailCommand{UserID: id}
	result, err := h.commandHandler.HandleVerifyEmail(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to verify email", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// VerifyPhone godoc
// @Summary Verify phone
// @Description Mark user phone as verified
// @Tags users,verification
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id}/verify-phone [post]
func (h *UserHandler) VerifyPhone(c *gin.Context) {
	id := c.Param("id")
	
	cmd := user.VerifyPhoneCommand{UserID: id}
	result, err := h.commandHandler.HandleVerifyPhone(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to verify phone", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// VerifyIdentity godoc
// @Summary Verify identity
// @Description Mark user identity as verified
// @Tags users,verification
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserDTO
// @Failure 404 {object} httpx.ErrorResponse
// @Router /users/{id}/verify-identity [post]
func (h *UserHandler) VerifyIdentity(c *gin.Context) {
	id := c.Param("id")
	
	cmd := user.VerifyIdentityCommand{UserID: id}
	result, err := h.commandHandler.HandleVerifyIdentity(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to verify identity", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// ============================================================================
// MODERATION OPERATIONS (Admin only)
// ============================================================================

// SuspendUser godoc
// @Summary Suspend user
// @Description Suspend a user account
// @Tags users,admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{reason string, suspended_by string} true "Suspension data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/suspend [post]
func (h *UserHandler) SuspendUser(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		Reason      string `json:"reason" binding:"required"`
		SuspendedBy string `json:"suspended_by" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	cmd := user.SuspendUserCommand{
		UserID:      id,
		Reason:      req.Reason,
		SuspendedBy: req.SuspendedBy,
	}
	
	result, err := h.commandHandler.HandleSuspendUser(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to suspend user", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// BanUser godoc
// @Summary Ban user
// @Description Ban a user account permanently
// @Tags users,admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{reason string, banned_by string} true "Ban data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/ban [post]
func (h *UserHandler) BanUser(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		Reason   string `json:"reason" binding:"required"`
		BannedBy string `json:"banned_by" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	cmd := user.BanUserCommand{
		UserID:   id,
		Reason:   req.Reason,
		BannedBy: req.BannedBy,
	}
	
	result, err := h.commandHandler.HandleBanUser(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to ban user", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// RestoreUser godoc
// @Summary Restore user
// @Description Restore a suspended or banned user
// @Tags users,admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{restored_by string} true "Restore data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/restore [post]
func (h *UserHandler) RestoreUser(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		RestoredBy string `json:"restored_by" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	cmd := user.RestoreUserCommand{
		UserID:     id,
		RestoredBy: req.RestoredBy,
	}
	
	result, err := h.commandHandler.HandleRestoreUser(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to restore user", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// AddWarning godoc
// @Summary Add warning
// @Description Add a warning to user account
// @Tags users,admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{reason string, issued_by string} true "Warning data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/warn [post]
func (h *UserHandler) AddWarning(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		Reason   string `json:"reason" binding:"required"`
		IssuedBy string `json:"issued_by" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	cmd := user.AddWarningCommand{
		UserID:   id,
		Reason:   req.Reason,
		IssuedBy: req.IssuedBy,
	}
	
	result, err := h.commandHandler.HandleAddWarning(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to add warning", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// ============================================================================
// BADGE OPERATIONS (Admin only)
// ============================================================================

// AssignBadge godoc
// @Summary Assign badge
// @Description Assign a badge to user
// @Tags users,admin,badges
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{badge_type string} true "Badge data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/badges [post]
func (h *UserHandler) AssignBadge(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		BadgeType string `json:"badge_type" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	badgeType := userDomain.BadgeType(req.BadgeType)
	if !badgeType.Valid() {
		httpx.Error(c, http.StatusBadRequest, "Invalid badge type", nil)
		return
	}
	
	cmd := user.AssignBadgeCommand{
		UserID:    id,
		BadgeType: badgeType,
	}
	
	result, err := h.commandHandler.HandleAssignBadge(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to assign badge", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// RemoveBadge godoc
// @Summary Remove badge
// @Description Remove a badge from user
// @Tags users,admin,badges
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param badge_type path string true "Badge type"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/badges/{badge_type} [delete]
func (h *UserHandler) RemoveBadge(c *gin.Context) {
	id := c.Param("id")
	badgeTypeStr := c.Param("badge_type")
	
	badgeType := userDomain.BadgeType(badgeTypeStr)
	if !badgeType.Valid() {
		httpx.Error(c, http.StatusBadRequest, "Invalid badge type", nil)
		return
	}
	
	cmd := user.RemoveBadgeCommand{
		UserID:    id,
		BadgeType: badgeType,
	}
	
	result, err := h.commandHandler.HandleRemoveBadge(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to remove badge", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// SetFeatured godoc
// @Summary Set featured
// @Description Set user as featured or unfeatured
// @Tags users,admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{featured bool} true "Featured data"
// @Success 200 {object} user.UserDTO
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id}/featured [put]
func (h *UserHandler) SetFeatured(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		Featured bool `json:"featured"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	cmd := user.SetFeaturedCommand{
		UserID:   id,
		Featured: req.Featured,
	}
	
	result, err := h.commandHandler.HandleSetFeatured(c.Request.Context(), cmd)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to set featured", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// ============================================================================
// STATISTICS & ANALYTICS
// ============================================================================

// GetUserStatistics godoc
// @Summary Get user statistics
// @Description Get comprehensive user statistics
// @Tags users,statistics
// @Accept json
// @Produce json
// @Success 200 {object} user.UserStatistics
// @Failure 500 {object} httpx.ErrorResponse
// @Router /admin/users/statistics [get]
func (h *UserHandler) GetUserStatistics(c *gin.Context) {
	query := user.GetUserStatisticsQuery{}
	result, err := h.queryHandler.HandleGetUserStatistics(c.Request.Context(), query)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to get statistics", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// GetUserGrowthStats godoc
// @Summary Get growth statistics
// @Description Get user growth statistics for specified period
// @Tags users,statistics
// @Accept json
// @Produce json
// @Param days query int false "Number of days" default(30)
// @Success 200 {object} user.UserGrowthStats
// @Failure 500 {object} httpx.ErrorResponse
// @Router /admin/users/statistics/growth [get]
func (h *UserHandler) GetUserGrowthStats(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))
	
	query := user.GetUserGrowthStatsQuery{Days: days}
	result, err := h.queryHandler.HandleGetUserGrowthStats(c.Request.Context(), query)
	if err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to get growth statistics", err)
		return
	}
	
	httpx.Success(c, http.StatusOK, result)
}

// ============================================================================
// DELETE OPERATIONS
// ============================================================================

// DeleteUser godoc
// @Summary Delete user
// @Description Soft delete a user account
// @Tags users,admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{deleted_by string} true "Delete data"
// @Success 204
// @Failure 400 {object} httpx.ErrorResponse
// @Router /admin/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		DeletedBy string `json:"deleted_by" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	cmd := user.DeleteUserCommand{
		UserID:    id,
		DeletedBy: req.DeletedBy,
	}
	
	if err := h.commandHandler.HandleDeleteUser(c.Request.Context(), cmd); err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}
	
	c.Status(http.StatusNoContent)
}

// ============================================================================
// ACTIVITY TRACKING
// ============================================================================

// RecordLogin godoc
// @Summary Record login
// @Description Record user login activity
// @Tags users,activity
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} httpx.ErrorResponse
// @Router /users/{id}/login [post]
func (h *UserHandler) RecordLogin(c *gin.Context) {
	id := c.Param("id")
	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	
	cmd := user.RecordLoginCommand{
		UserID:    id,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
	
	if err := h.commandHandler.HandleRecordLogin(c.Request.Context(), cmd); err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to record login", err)
		return
	}
	
	c.Status(http.StatusNoContent)
}

// UpdateLastSeen godoc
// @Summary Update last seen
// @Description Update user's last seen timestamp
// @Tags users,activity
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Router /users/{id}/last-seen [put]
func (h *UserHandler) UpdateLastSeen(c *gin.Context) {
	id := c.Param("id")
	
	if err := h.service.UpdateLastSeen(c.Request.Context(), id); err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to update last seen", err)
		return
	}
	
	c.Status(http.StatusNoContent)
}

// SetOnlineStatus godoc
// @Summary Set online status
// @Description Set user online/offline status
// @Tags users,activity
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param data body object{is_online bool} true "Online status"
// @Success 204
// @Router /users/{id}/online-status [put]
func (h *UserHandler) SetOnlineStatus(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		IsOnline bool `json:"is_online"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	if err := h.service.SetOnlineStatus(c.Request.Context(), id, req.IsOnline); err != nil {
		httpx.Error(c, http.StatusInternalServerError, "Failed to set online status", err)
		return
	}
	
	c.Status(http.StatusNoContent)
}

// ============================================================================
// HELPER METHODS
// ============================================================================

// parseListFilter parses query parameters into ListFilter
func (h *UserHandler) parseListFilter(c *gin.Context) userDomain.ListFilter {
	filter := userDomain.NewListFilter()
	
	// Pagination
	if page := c.Query("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			filter.Page = p
		}
	}
	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Limit = l
		}
	}
	
	// Sorting
	if sortBy := c.Query("sort_by"); sortBy != "" {
		filter.SortBy = sortBy
	}
	if sortOrder := c.Query("sort_order"); sortOrder != "" {
		filter.SortOrder = sortOrder
	}
	
	// User type filter
	if userType := c.Query("user_type"); userType != "" {
		ut := userDomain.UserType(userType)
		filter.UserType = &ut
	}
	
	// Status filter
	if status := c.Query("status"); status != "" {
		s := userDomain.AccountStatus(status)
		filter.Status = &s
	}
	
	// Verification status filter
	if verStatus := c.Query("verification_status"); verStatus != "" {
		vs := userDomain.VerificationStatus(verStatus)
		filter.VerificationStatus = &vs
	}
	
	// Location filters
	if country := c.Query("country"); country != "" {
		filter.Country = country
	}
	if city := c.Query("city"); city != "" {
		filter.City = city
	}
	
	// Rating filters
	if minRating := c.Query("min_rating"); minRating != "" {
		if mr, err := strconv.ParseFloat(minRating, 64); err == nil {
			filter.MinRating = mr
		}
	}
	if maxRating := c.Query("max_rating"); maxRating != "" {
		if mr, err := strconv.ParseFloat(maxRating, 64); err == nil {
			filter.MaxRating = mr
		}
	}
	
	// Boolean filters
	if isFeatured := c.Query("is_featured"); isFeatured != "" {
		if f, err := strconv.ParseBool(isFeatured); err == nil {
			filter.IsFeatured = &f
		}
	}
	if isTopRated := c.Query("is_top_rated"); isTopRated != "" {
		if tr, err := strconv.ParseBool(isTopRated); err == nil {
			filter.IsTopRated = &tr
		}
	}
	if isOnline := c.Query("is_online"); isOnline != "" {
		if o, err := strconv.ParseBool(isOnline); err == nil {
			filter.IsOnline = &o
		}
	}
	if emailVerified := c.Query("email_verified"); emailVerified != "" {
		if ev, err := strconv.ParseBool(emailVerified); err == nil {
			filter.EmailVerified = &ev
		}
	}
	if identityVerified := c.Query("identity_verified"); identityVerified != "" {
		if iv, err := strconv.ParseBool(identityVerified); err == nil {
			filter.IdentityVerified = &iv
		}
	}
	
	return filter
}