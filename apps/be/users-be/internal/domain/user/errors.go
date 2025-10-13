// internal/domain/user/errors.go
package user

import "errors"

var (
    // Not found errors
    ErrUserNotFound = errors.New("user not found")
    
    // Duplicate errors
    ErrEmailTaken    = errors.New("email already taken")
    ErrUsernameTaken = errors.New("username already taken")
    ErrKeycloakIDExists = errors.New("keycloak ID already exists")
    
    // Validation errors
    ErrInvalidEmail       = errors.New("invalid email format")
    ErrInvalidPhone       = errors.New("invalid phone format")
    ErrInvalidUserType    = errors.New("invalid user type")
    ErrInvalidStatus      = errors.New("invalid account status")
    ErrInvalidCountryCode = errors.New("invalid country code")
    ErrInvalidTimezone    = errors.New("invalid timezone")
    ErrInvalidCurrency    = errors.New("invalid currency code")
    
    // State errors
    ErrUserSuspended      = errors.New("user account is suspended")
    ErrUserBanned         = errors.New("user account is banned")
    ErrUserDeleted        = errors.New("user account is deleted")
    ErrUserInactive       = errors.New("user account is inactive")
    ErrEmailNotVerified   = errors.New("email not verified")
    ErrPhoneNotVerified   = errors.New("phone not verified")
    ErrIdentityNotVerified = errors.New("identity not verified")
    ErrPaymentNotVerified = errors.New("payment method not verified")
    
    // Operation errors
    ErrCannotChangeType   = errors.New("cannot change user type")
    ErrCannotDeleteAdmin  = errors.New("cannot delete admin user")
    ErrAccountLocked      = errors.New("account locked due to failed login attempts")
    ErrProfileIncomplete  = errors.New("profile is incomplete")
    ErrCannotReceivePayments = errors.New("user cannot receive payments")
    ErrCannotHire         = errors.New("user cannot hire")
    
    // Business rule errors
    ErrInsufficientReputation = errors.New("insufficient reputation score")
    ErrHasActiveWarnings  = errors.New("user has active warnings")
    ErrMaxReferralsReached = errors.New("maximum referrals reached")
)