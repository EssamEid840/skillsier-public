// internal/domain/user/errors.go
package user

import (
	"errors"
	"fmt"
)

// ============================================================================
// USER ENTITY ERRORS
// ============================================================================

var (
	// Not Found Errors
	ErrUserNotFound        = errors.New("user not found")
	ErrUserNotFoundByID    = errors.New("user not found by ID")
	ErrUserNotFoundByEmail = errors.New("user not found by email")
	ErrUserNotFoundByUsername = errors.New("user not found by username")
	ErrUserNotFoundByKeycloakID = errors.New("user not found by Keycloak ID")
	
	// Duplicate/Conflict Errors
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrEmailAlreadyTaken     = errors.New("email address is already taken")
	ErrUsernameAlreadyTaken  = errors.New("username is already taken")
	ErrKeycloakIDAlreadyUsed = errors.New("Keycloak ID is already associated with another user")
	ErrReferralCodeAlreadyUsed = errors.New("referral code is already in use")
	
	// Validation Errors - General
	ErrInvalidUserID       = errors.New("invalid user ID")
	ErrInvalidKeycloakID   = errors.New("invalid Keycloak ID")
	ErrInvalidUserType     = errors.New("invalid user type")
	ErrInvalidAccountStatus = errors.New("invalid account status")
	ErrInvalidVerificationStatus = errors.New("invalid verification status")
	
	// Validation Errors - Required Fields
	ErrEmailRequired     = errors.New("email is required")
	ErrUsernameRequired  = errors.New("username is required")
	ErrFirstNameRequired = errors.New("first name is required")
	ErrLastNameRequired  = errors.New("last name is required")
	ErrUserTypeRequired  = errors.New("user type is required")
	
	// Validation Errors - Format
	ErrInvalidEmailFormat    = errors.New("invalid email format")
	ErrEmailTooLong          = errors.New("email too long (max 255 characters)")
	ErrInvalidUsernameFormat = errors.New("invalid username format")
	ErrUsernameTooShort      = errors.New("username too short (min 3 characters)")
	ErrUsernameTooLong       = errors.New("username too long (max 30 characters)")
	ErrInvalidPhoneFormat    = errors.New("invalid phone number format")
	ErrInvalidBioLength      = errors.New("bio exceeds maximum length (5000 characters)")
	ErrInvalidTaglineLength  = errors.New("tagline exceeds maximum length (100 characters)")
	
	// Business Logic Errors
	ErrAccountNotActive      = errors.New("account is not active")
	ErrAccountSuspended      = errors.New("account is suspended")
	ErrAccountBanned         = errors.New("account is banned")
	ErrAccountDeleted        = errors.New("account is deleted")
	ErrAccountRestricted     = errors.New("account is restricted")
	ErrEmailNotVerified      = errors.New("email is not verified")
	ErrPhoneNotVerified      = errors.New("phone number is not verified")
	ErrIdentityNotVerified   = errors.New("identity is not verified")
	ErrVerificationPending   = errors.New("verification is pending")
	ErrVerificationRejected  = errors.New("verification was rejected")
	
	// Permission Errors
	ErrUnauthorized          = errors.New("unauthorized action")
	ErrForbidden             = errors.New("forbidden: insufficient permissions")
	ErrCannotModifyOwnStatus = errors.New("cannot modify your own account status")
	ErrCannotDeleteSelf      = errors.New("cannot delete your own account")
	ErrCannotBanSelf         = errors.New("cannot ban your own account")
	ErrCannotSuspendSelf     = errors.New("cannot suspend your own account")
	
	// State Transition Errors
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrCannotActivateBannedUser = errors.New("cannot activate a banned user")
	ErrCannotReactivateDeletedUser = errors.New("cannot reactivate a deleted user")
	ErrAlreadyActive         = errors.New("user is already active")
	ErrAlreadySuspended      = errors.New("user is already suspended")
	ErrAlreadyBanned         = errors.New("user is already banned")
	ErrAlreadyDeleted        = errors.New("user is already deleted")
	
	// Rating/Stats Errors
	ErrInvalidRating         = errors.New("rating must be between 0 and 5")
	ErrInvalidCompletionRate = errors.New("completion rate must be between 0 and 100")
	ErrInvalidResponseTime   = errors.New("response time cannot be negative")
	
	// Profile Completeness Errors
	ErrProfileIncomplete     = errors.New("profile is incomplete")
	ErrMissingRequiredFields = errors.New("missing required profile fields")
	
	// Referral Errors
	ErrInvalidReferralCode   = errors.New("invalid referral code")
	ErrSelfReferralNotAllowed = errors.New("self-referral is not allowed")
	ErrReferralExpired       = errors.New("referral code has expired")
	
	// Availability Errors
	ErrInvalidAvailability   = errors.New("invalid availability status")
	ErrHoursPerWeekInvalid   = errors.New("hours per week must be between 0 and 168")
	
	// Badge Errors
	ErrInvalidBadgeType      = errors.New("invalid badge type")
	ErrBadgeAlreadyAssigned  = errors.New("badge already assigned to user")
	ErrBadgeNotFound         = errors.New("badge not found")
	
	// Search/Filter Errors
	ErrInvalidSearchQuery    = errors.New("invalid search query")
	ErrInvalidFilterCriteria = errors.New("invalid filter criteria")
	ErrInvalidSortField      = errors.New("invalid sort field")
	
	// Batch Operation Errors
	ErrEmptyBatchOperation   = errors.New("batch operation contains no items")
	ErrBatchSizeTooLarge     = errors.New("batch size exceeds maximum allowed")
	ErrPartialBatchFailure   = errors.New("some items in batch operation failed")
)

// ============================================================================
// CUSTOM ERROR TYPES WITH CONTEXT
// ============================================================================

// UserError wraps errors with additional context
type UserError struct {
	Code    string
	Message string
	Field   string
	Err     error
}

func (e *UserError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s (field: %s)", e.Code, e.Message, e.Field)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *UserError) Unwrap() error {
	return e.Err
}

// NewUserError creates a new UserError
func NewUserError(code, message, field string, err error) *UserError {
	return &UserError{
		Code:    code,
		Message: message,
		Field:   field,
		Err:     err,
	}
}

// ValidationError represents validation failure
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s", e.Field, e.Message)
}

// NewValidationError creates a new ValidationError
func NewValidationError(field, message string, value interface{}) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
		Value:   value,
	}
}

// ValidationErrors represents multiple validation errors
type ValidationErrors struct {
	Errors []*ValidationError
}

func (ve *ValidationErrors) Error() string {
	if len(ve.Errors) == 0 {
		return "validation failed"
	}
	if len(ve.Errors) == 1 {
		return ve.Errors[0].Error()
	}
	return fmt.Sprintf("validation failed: %d errors", len(ve.Errors))
}

func (ve *ValidationErrors) Add(field, message string, value interface{}) {
	ve.Errors = append(ve.Errors, NewValidationError(field, message, value))
}

func (ve *ValidationErrors) HasErrors() bool {
	return len(ve.Errors) > 0
}

// NewValidationErrors creates a new ValidationErrors
func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		Errors: make([]*ValidationError, 0),
	}
}

// ============================================================================
// ERROR CHECKING HELPERS
// ============================================================================

// IsNotFoundError checks if error is a not found error
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrUserNotFound) ||
		errors.Is(err, ErrUserNotFoundByID) ||
		errors.Is(err, ErrUserNotFoundByEmail) ||
		errors.Is(err, ErrUserNotFoundByUsername) ||
		errors.Is(err, ErrUserNotFoundByKeycloakID)
}

// IsConflictError checks if error is a conflict/duplicate error
func IsConflictError(err error) bool {
	return errors.Is(err, ErrUserAlreadyExists) ||
		errors.Is(err, ErrEmailAlreadyTaken) ||
		errors.Is(err, ErrUsernameAlreadyTaken) ||
		errors.Is(err, ErrKeycloakIDAlreadyUsed) ||
		errors.Is(err, ErrReferralCodeAlreadyUsed)
}

// IsValidationError checks if error is a validation error
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	
	// Check for our custom validation error types
	var valErr *ValidationError
	var valErrs *ValidationErrors
	if errors.As(err, &valErr) || errors.As(err, &valErrs) {
		return true
	}
	
	// Check for standard validation errors
	return errors.Is(err, ErrInvalidEmailFormat) ||
		errors.Is(err, ErrEmailRequired) ||
		errors.Is(err, ErrUsernameRequired) ||
		errors.Is(err, ErrInvalidUsernameFormat) ||
		errors.Is(err, ErrInvalidUserType) ||
		errors.Is(err, ErrInvalidAccountStatus)
}

// IsAuthorizationError checks if error is an authorization error
func IsAuthorizationError(err error) bool {
	return errors.Is(err, ErrUnauthorized) ||
		errors.Is(err, ErrForbidden) ||
		errors.Is(err, ErrAccountNotActive) ||
		errors.Is(err, ErrAccountSuspended) ||
		errors.Is(err, ErrAccountBanned) ||
		errors.Is(err, ErrAccountDeleted)
}

// IsAccountBlockedError checks if error is due to blocked account
func IsAccountBlockedError(err error) bool {
	return errors.Is(err, ErrAccountSuspended) ||
		errors.Is(err, ErrAccountBanned) ||
		errors.Is(err, ErrAccountDeleted) ||
		errors.Is(err, ErrAccountRestricted)
}

// ============================================================================
// ERROR CODE CONSTANTS (for API responses)
// ============================================================================

const (
	ErrorCodeUserNotFound        = "USER_NOT_FOUND"
	ErrorCodeEmailTaken          = "EMAIL_TAKEN"
	ErrorCodeUsernameTaken       = "USERNAME_TAKEN"
	ErrorCodeInvalidEmail        = "INVALID_EMAIL"
	ErrorCodeInvalidUsername     = "INVALID_USERNAME"
	ErrorCodeAccountSuspended    = "ACCOUNT_SUSPENDED"
	ErrorCodeAccountBanned       = "ACCOUNT_BANNED"
	ErrorCodeAccountDeleted      = "ACCOUNT_DELETED"
	ErrorCodeUnauthorized        = "UNAUTHORIZED"
	ErrorCodeForbidden           = "FORBIDDEN"
	ErrorCodeValidationFailed    = "VALIDATION_FAILED"
	ErrorCodeInvalidInput        = "INVALID_INPUT"
	ErrorCodeProfileIncomplete   = "PROFILE_INCOMPLETE"
	ErrorCodeEmailNotVerified    = "EMAIL_NOT_VERIFIED"
	ErrorCodeVerificationPending = "VERIFICATION_PENDING"
	ErrorCodeInvalidStatusTransition = "INVALID_STATUS_TRANSITION"
)

// GetErrorCode returns the error code for an error
func GetErrorCode(err error) string {
	switch {
	case errors.Is(err, ErrUserNotFound):
		return ErrorCodeUserNotFound
	case errors.Is(err, ErrEmailAlreadyTaken):
		return ErrorCodeEmailTaken
	case errors.Is(err, ErrUsernameAlreadyTaken):
		return ErrorCodeUsernameTaken
	case errors.Is(err, ErrInvalidEmailFormat):
		return ErrorCodeInvalidEmail
	case errors.Is(err, ErrInvalidUsernameFormat):
		return ErrorCodeInvalidUsername
	case errors.Is(err, ErrAccountSuspended):
		return ErrorCodeAccountSuspended
	case errors.Is(err, ErrAccountBanned):
		return ErrorCodeAccountBanned
	case errors.Is(err, ErrAccountDeleted):
		return ErrorCodeAccountDeleted
	case errors.Is(err, ErrUnauthorized):
		return ErrorCodeUnauthorized
	case errors.Is(err, ErrForbidden):
		return ErrorCodeForbidden
	case errors.Is(err, ErrProfileIncomplete):
		return ErrorCodeProfileIncomplete
	case errors.Is(err, ErrEmailNotVerified):
		return ErrorCodeEmailNotVerified
	case errors.Is(err, ErrVerificationPending):
		return ErrorCodeVerificationPending
	case errors.Is(err, ErrInvalidStatusTransition):
		return ErrorCodeInvalidStatusTransition
	case IsValidationError(err):
		return ErrorCodeValidationFailed
	default:
		return "INTERNAL_ERROR"
	}
}