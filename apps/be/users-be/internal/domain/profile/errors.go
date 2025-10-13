// internal/domain/profile/errors.go
package profile

import "errors"

var (
    ErrProfileNotFound        = errors.New("profile not found")
    ErrProfileAlreadyExists   = errors.New("profile already exists for this user")
    ErrInvalidBio             = errors.New("invalid bio format")
    ErrInvalidLocation        = errors.New("invalid location")
    ErrInvalidRate            = errors.New("invalid hourly rate")
    ErrInvalidURL             = errors.New("invalid URL format")
    ErrProfileIncomplete      = errors.New("profile is incomplete")
    ErrQualityScoreTooLow     = errors.New("profile quality score is too low")
    ErrInvalidAvailability    = errors.New("invalid availability status")
)
