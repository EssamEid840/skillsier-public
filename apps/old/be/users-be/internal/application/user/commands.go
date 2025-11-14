// internal/application/user/commands.go
package user

import (
	"context"
	userDomain "users-be/internal/domain/user"
)

// ============================================================================
// COMMAND DEFINITIONS (CQRS Pattern)
// ============================================================================

// CreateUserCommand represents a command to create a new user
type CreateUserCommand struct {
	KeycloakID       string
	Username         string
	Email            string
	FirstName        string
	LastName         string
	UserType         string
	PhoneNumber      string
	PhoneCountryCode string
	City             string
	Country          string
	CountryCode      string
	Timezone         string
	Bio              string
	Tagline          string
}

// UpdateUserCommand represents a command to update user information
type UpdateUserCommand struct {
	UserID           string
	FirstName        *string
	LastName         *string
	DisplayName      *string
	Bio              *string
	Tagline          *string
	Title            *string
	Website          *string
	ProfilePictureURL *string
	CoverImageURL    *string
	PhoneNumber      *string
	PhoneCountryCode *string
	City             *string
	Country          *string
	CountryCode      *string
	Timezone         *string
}

// VerifyEmailCommand represents a command to verify email
type VerifyEmailCommand struct {
	UserID string
}

// VerifyPhoneCommand represents a command to verify phone
type VerifyPhoneCommand struct {
	UserID string
}

// VerifyIdentityCommand represents a command to verify identity
type VerifyIdentityCommand struct {
	UserID string
}

// SuspendUserCommand represents a command to suspend a user
type SuspendUserCommand struct {
	UserID      string
	Reason      string
	SuspendedBy string
}

// BanUserCommand represents a command to ban a user
type BanUserCommand struct {
	UserID   string
	Reason   string
	BannedBy string
}

// RestoreUserCommand represents a command to restore a user
type RestoreUserCommand struct {
	UserID     string
	RestoredBy string
}

// AddWarningCommand represents a command to add a warning
type AddWarningCommand struct {
	UserID   string
	Reason   string
	IssuedBy string
}

// AssignBadgeCommand represents a command to assign a badge
type AssignBadgeCommand struct {
	UserID    string
	BadgeType userDomain.BadgeType
}

// RemoveBadgeCommand represents a command to remove a badge
type RemoveBadgeCommand struct {
	UserID    string
	BadgeType userDomain.BadgeType
}

// UpdateAvailabilityCommand represents a command to update availability
type UpdateAvailabilityCommand struct {
	UserID       string
	Status       string
	HoursPerWeek int
}

// UpdateSettingsCommand represents a command to update settings
type UpdateSettingsCommand struct {
	UserID            string
	ProfileVisibility *string
	ShowEmail         *bool
	ShowPhone         *bool
	ShowLocation      *bool
	SearchableProfile *bool
	AcceptingWork     *bool
}

// RecordLoginCommand represents a command to record login
type RecordLoginCommand struct {
	UserID    string
	IPAddress string
	UserAgent string
}

// SetFeaturedCommand represents a command to set featured flag
type SetFeaturedCommand struct {
	UserID   string
	Featured bool
}

// DeleteUserCommand represents a command to delete a user
type DeleteUserCommand struct {
	UserID    string
	DeletedBy string
}

// UpdateProfileCommand represents a command to update profile
type UpdateProfileCommand struct {
	UserID  string
	Bio     string
	Tagline string
	Title   string
}

// UpdateRatingCommand represents a command to update user rating
type UpdateRatingCommand struct {
	UserID       string
	Rating       float64
	TotalReviews int
}

// UpdateStatsCommand represents a command to update user stats
type UpdateStatsCommand struct {
	UserID        string
	CompletedJobs int
	TotalJobs     int
	SuccessRate   float64
}

// UpdateEarningsCommand represents a command to update earnings
type UpdateEarningsCommand struct {
	UserID  string
	Amount  float64
}

// UpdateSpendingCommand represents a command to update spending
type UpdateSpendingCommand struct {
	UserID string
	Amount float64
}

// ============================================================================
// COMMAND HANDLERS
// ============================================================================

// CommandHandler handles user commands
type CommandHandler struct {
	service *Service
}

// NewCommandHandler creates a new command handler
func NewCommandHandler(service *Service) *CommandHandler {
	return &CommandHandler{service: service}
}

// HandleCreateUser handles create user command
func (h *CommandHandler) HandleCreateUser(ctx context.Context, cmd CreateUserCommand) (*UserDTO, error) {
	dto := &CreateUserDTO{
		KeycloakID:       cmd.KeycloakID,
		Username:         cmd.Username,
		Email:            cmd.Email,
		FirstName:        cmd.FirstName,
		LastName:         cmd.LastName,
		UserType:         cmd.UserType,
		PhoneNumber:      cmd.PhoneNumber,
		PhoneCountryCode: cmd.PhoneCountryCode,
		City:             cmd.City,
		Country:          cmd.Country,
		CountryCode:      cmd.CountryCode,
		Timezone:         cmd.Timezone,
		Bio:              cmd.Bio,
		Tagline:          cmd.Tagline,
	}
	
	return h.service.CreateUser(ctx, dto)
}

// HandleUpdateUser handles update user command
func (h *CommandHandler) HandleUpdateUser(ctx context.Context, cmd UpdateUserCommand) (*UserDTO, error) {
	dto := &UpdateUserDTO{
		FirstName:         cmd.FirstName,
		LastName:          cmd.LastName,
		DisplayName:       cmd.DisplayName,
		Bio:               cmd.Bio,
		Tagline:           cmd.Tagline,
		Title:             cmd.Title,
		Website:           cmd.Website,
		ProfilePictureURL: cmd.ProfilePictureURL,
		CoverImageURL:     cmd.CoverImageURL,
		PhoneNumber:       cmd.PhoneNumber,
		PhoneCountryCode:  cmd.PhoneCountryCode,
		City:              cmd.City,
		Country:           cmd.Country,
		CountryCode:       cmd.CountryCode,
		Timezone:          cmd.Timezone,
	}
	
	return h.service.UpdateUser(ctx, cmd.UserID, dto)
}

// HandleVerifyEmail handles verify email command
func (h *CommandHandler) HandleVerifyEmail(ctx context.Context, cmd VerifyEmailCommand) (*UserDTO, error) {
	return h.service.VerifyEmail(ctx, cmd.UserID)
}

// HandleVerifyPhone handles verify phone command
func (h *CommandHandler) HandleVerifyPhone(ctx context.Context, cmd VerifyPhoneCommand) (*UserDTO, error) {
	return h.service.VerifyPhone(ctx, cmd.UserID)
}

// HandleVerifyIdentity handles verify identity command
func (h *CommandHandler) HandleVerifyIdentity(ctx context.Context, cmd VerifyIdentityCommand) (*UserDTO, error) {
	return h.service.VerifyIdentity(ctx, cmd.UserID)
}

// HandleSuspendUser handles suspend user command
func (h *CommandHandler) HandleSuspendUser(ctx context.Context, cmd SuspendUserCommand) (*UserDTO, error) {
	return h.service.SuspendUser(ctx, cmd.UserID, cmd.Reason, cmd.SuspendedBy)
}

// HandleBanUser handles ban user command
func (h *CommandHandler) HandleBanUser(ctx context.Context, cmd BanUserCommand) (*UserDTO, error) {
	return h.service.BanUser(ctx, cmd.UserID, cmd.Reason, cmd.BannedBy)
}

// HandleRestoreUser handles restore user command
func (h *CommandHandler) HandleRestoreUser(ctx context.Context, cmd RestoreUserCommand) (*UserDTO, error) {
	return h.service.RestoreUser(ctx, cmd.UserID, cmd.RestoredBy)
}

// HandleAddWarning handles add warning command
func (h *CommandHandler) HandleAddWarning(ctx context.Context, cmd AddWarningCommand) (*UserDTO, error) {
	return h.service.AddWarning(ctx, cmd.UserID, cmd.Reason, cmd.IssuedBy)
}

// HandleAssignBadge handles assign badge command
func (h *CommandHandler) HandleAssignBadge(ctx context.Context, cmd AssignBadgeCommand) (*UserDTO, error) {
	return h.service.AssignBadge(ctx, cmd.UserID, cmd.BadgeType)
}

// HandleRemoveBadge handles remove badge command
func (h *CommandHandler) HandleRemoveBadge(ctx context.Context, cmd RemoveBadgeCommand) (*UserDTO, error) {
	return h.service.RemoveBadge(ctx, cmd.UserID, cmd.BadgeType)
}

// HandleUpdateAvailability handles update availability command
func (h *CommandHandler) HandleUpdateAvailability(ctx context.Context, cmd UpdateAvailabilityCommand) (*UserDTO, error) {
	dto := &UpdateAvailabilityDTO{
		Status:       cmd.Status,
		HoursPerWeek: cmd.HoursPerWeek,
	}
	
	return h.service.UpdateAvailability(ctx, cmd.UserID, dto)
}

// HandleUpdateSettings handles update settings command
func (h *CommandHandler) HandleUpdateSettings(ctx context.Context, cmd UpdateSettingsCommand) (*UserDTO, error) {
	dto := &UpdateSettingsDTO{
		ProfileVisibility: cmd.ProfileVisibility,
		ShowEmail:         cmd.ShowEmail,
		ShowPhone:         cmd.ShowPhone,
		ShowLocation:      cmd.ShowLocation,
		SearchableProfile: cmd.SearchableProfile,
		AcceptingWork:     cmd.AcceptingWork,
	}
	
	return h.service.UpdateSettings(ctx, cmd.UserID, dto)
}

// HandleRecordLogin handles record login command
func (h *CommandHandler) HandleRecordLogin(ctx context.Context, cmd RecordLoginCommand) error {
	return h.service.RecordLogin(ctx, cmd.UserID, cmd.IPAddress, cmd.UserAgent)
}

// HandleSetFeatured handles set featured command
func (h *CommandHandler) HandleSetFeatured(ctx context.Context, cmd SetFeaturedCommand) (*UserDTO, error) {
	return h.service.SetFeatured(ctx, cmd.UserID, cmd.Featured)
}

// HandleDeleteUser handles delete user command
func (h *CommandHandler) HandleDeleteUser(ctx context.Context, cmd DeleteUserCommand) error {
	return h.service.DeleteUser(ctx, cmd.UserID, cmd.DeletedBy)
}

// HandleUpdateProfile handles update profile command
func (h *CommandHandler) HandleUpdateProfile(ctx context.Context, cmd UpdateProfileCommand) (*UserDTO, error) {
	dto := &UpdateProfileDTO{
		Bio:     cmd.Bio,
		Tagline: cmd.Tagline,
		Title:   cmd.Title,
	}
	
	return h.service.UpdateProfile(ctx, cmd.UserID, dto)
}

// ============================================================================
// BATCH COMMAND DEFINITIONS
// ============================================================================

// UpdateStatusBatchCommand represents a batch status update command
type UpdateStatusBatchCommand struct {
	UserIDs   []string
	Status    userDomain.AccountStatus
	UpdatedBy string
}

// VerifyEmailBatchCommand represents a batch email verification command
type VerifyEmailBatchCommand struct {
	UserIDs []string
}

// DeleteUserBatchCommand represents a batch delete command
type DeleteUserBatchCommand struct {
	UserIDs   []string
	DeletedBy string
}

// AssignBadgeBatchCommand represents a batch badge assignment command
type AssignBadgeBatchCommand struct {
	UserIDs   []string
	BadgeType userDomain.BadgeType
}

// ============================================================================
// COMMAND VALIDATION
// ============================================================================

// Validate validates CreateUserCommand
func (cmd CreateUserCommand) Validate() error {
	if cmd.KeycloakID == "" {
		return userDomain.ErrInvalidKeycloakID
	}
	if cmd.Username == "" {
		return userDomain.ErrUsernameRequired
	}
	if cmd.Email == "" {
		return userDomain.ErrEmailRequired
	}
	if cmd.FirstName == "" {
		return userDomain.ErrFirstNameRequired
	}
	if cmd.LastName == "" {
		return userDomain.ErrLastNameRequired
	}
	if cmd.UserType == "" {
		return userDomain.ErrUserTypeRequired
	}
	
	userType := userDomain.UserType(cmd.UserType)
	if !userType.Valid() {
		return userDomain.ErrInvalidUserType
	}
	
	return nil
}

// Validate validates UpdateUserCommand
func (cmd UpdateUserCommand) Validate() error {
	if cmd.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	
	// At least one field must be provided
	if cmd.FirstName == nil && cmd.LastName == nil && cmd.DisplayName == nil &&
		cmd.Bio == nil && cmd.Tagline == nil && cmd.Title == nil &&
		cmd.Website == nil && cmd.ProfilePictureURL == nil && cmd.CoverImageURL == nil &&
		cmd.PhoneNumber == nil && cmd.City == nil && cmd.Country == nil {
		return userDomain.ErrMissingRequiredFields
	}
	
	return nil
}

// Validate validates SuspendUserCommand
func (cmd SuspendUserCommand) Validate() error {
	if cmd.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	if cmd.Reason == "" {
		return userDomain.ErrMissingRequiredFields
	}
	if cmd.SuspendedBy == "" {
		return userDomain.ErrMissingRequiredFields
	}
	return nil
}

// Validate validates BanUserCommand
func (cmd BanUserCommand) Validate() error {
	if cmd.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	if cmd.Reason == "" {
		return userDomain.ErrMissingRequiredFields
	}
	if cmd.BannedBy == "" {
		return userDomain.ErrMissingRequiredFields
	}
	return nil
}

// Validate validates AddWarningCommand
func (cmd AddWarningCommand) Validate() error {
	if cmd.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	if cmd.Reason == "" {
		return userDomain.ErrMissingRequiredFields
	}
	if cmd.IssuedBy == "" {
		return userDomain.ErrMissingRequiredFields
	}
	return nil
}

// Validate validates AssignBadgeCommand
func (cmd AssignBadgeCommand) Validate() error {
	if cmd.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	if !cmd.BadgeType.Valid() {
		return userDomain.ErrInvalidBadgeType
	}
	return nil
}

// Validate validates UpdateAvailabilityCommand
func (cmd UpdateAvailabilityCommand) Validate() error {
	if cmd.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	
	status := userDomain.AvailabilityStatus(cmd.Status)
	if !status.Valid() {
		return userDomain.ErrInvalidAvailability
	}
	
	if cmd.HoursPerWeek < 0 || cmd.HoursPerWeek > 168 {
		return userDomain.ErrHoursPerWeekInvalid
	}
	
	return nil
}

// Validate validates DeleteUserCommand
func (cmd DeleteUserCommand) Validate() error {
	if cmd.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	if cmd.DeletedBy == "" {
		return userDomain.ErrMissingRequiredFields
	}
	return nil
}