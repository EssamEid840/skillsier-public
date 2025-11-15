package initial_entity

import (
	"errors"
	"fmt"
)

// Domain errors for InitialEntity
var (
	// ErrNotFound is returned when an entity is not found
	ErrNotFound = errors.New("initial entity not found")
	
	// ErrAlreadyExists is returned when an entity already exists
	ErrAlreadyExists = errors.New("initial entity already exists")
	
	// ErrNameRequired is returned when name is empty
	ErrNameRequired = errors.New("name is required")
	
	// ErrNameTooShort is returned when name is too short
	ErrNameTooShort = errors.New("name must be at least 3 characters")
	
	// ErrNameTooLong is returned when name is too long
	ErrNameTooLong = errors.New("name must not exceed 255 characters")
	
	// ErrOwnerIDRequired is returned when owner ID is missing
	ErrOwnerIDRequired = errors.New("owner ID is required")
	
	// ErrInvalidStatus is returned when status is invalid
	ErrInvalidStatus = errors.New("invalid status")
	
	// ErrInvalidStatusTransition is returned when status transition is not allowed
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	
	// ErrAlreadyDeleted is returned when trying to delete an already deleted entity
	ErrAlreadyDeleted = errors.New("entity is already deleted")
	
	// ErrCannotModifyDeleted is returned when trying to modify a deleted entity
	ErrCannotModifyDeleted = errors.New("cannot modify deleted entity")
	
	// ErrCannotModifyArchived is returned when trying to modify an archived entity
	ErrCannotModifyArchived = errors.New("cannot modify archived entity")
	
	// ErrUnauthorized is returned when user is not authorized to perform action
	ErrUnauthorized = errors.New("unauthorized to perform this action")
)

// ValidationError represents a validation error with field information
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// NewValidationError creates a new ValidationError
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// IsValidationError checks if an error is a ValidationError
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// DomainError represents a domain-specific error with context
type DomainError struct {
	Code    string
	Message string
	Cause   error
}

// Error implements the error interface
func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap implements the errors.Unwrap interface
func (e *DomainError) Unwrap() error {
	return e.Cause
}

// NewDomainError creates a new DomainError
func NewDomainError(code, message string, cause error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// Error codes for structured error handling
const (
	ErrCodeNotFound           = "ENTITY_NOT_FOUND"
	ErrCodeAlreadyExists      = "ENTITY_ALREADY_EXISTS"
	ErrCodeValidation         = "VALIDATION_ERROR"
	ErrCodeInvalidTransition  = "INVALID_STATUS_TRANSITION"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeDeleted            = "ENTITY_DELETED"
	ErrCodeInternal           = "INTERNAL_ERROR"
)