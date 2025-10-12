package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// ValidationError represents a single field validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors is a collection of validation errors
type ValidationErrors []ValidationError

// Error implements the error interface
func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return "validation failed"
	}
	
	messages := make([]string, len(v))
	for i, err := range v {
		messages[i] = fmt.Sprintf("%s: %s", err.Field, err.Message)
	}
	return strings.Join(messages, "; ")
}

// ParseJSONBody parses JSON request body into the provided struct
func ParseJSONBody(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return ErrBadRequest.WithDetails("request body is empty")
	}
	
	// Limit request body size to 10MB
	r.Body = http.MaxBytesReader(nil, r.Body, 10<<20)
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	
	if err := decoder.Decode(v); err != nil {
		if err == io.EOF {
			return ErrBadRequest.WithDetails("request body is empty")
		}
		return ErrBadRequest.WithDetails("invalid JSON: " + err.Error())
	}
	
	return nil
}

// Common validation helpers

// IsValidEmail checks if the string is a valid email address
func IsValidEmail(email string) bool {
	// Simple email regex
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsValidUsername checks if the string is a valid username
func IsValidUsername(username string) bool {
	// Username: 3-30 characters, alphanumeric and underscores
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`)
	return re.MatchString(username)
}

// IsValidURL checks if the string is a valid URL
func IsValidURL(urlStr string) bool {
	re := regexp.MustCompile(`^https?://[a-zA-Z0-9\-._~:/?#\[\]@!$&'()*+,;=]+$`)
	return re.MatchString(urlStr)
}

// IsValidUUID checks if the string is a valid UUID
func IsValidUUID(uuid string) bool {
	re := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	return re.MatchString(uuid)
}

// Required checks if a string is not empty
func Required(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MinLength checks if a string meets minimum length
func MinLength(value string, min int) bool {
	return len(strings.TrimSpace(value)) >= min
}

// MaxLength checks if a string doesn't exceed maximum length
func MaxLength(value string, max int) bool {
	return len(strings.TrimSpace(value)) <= max
}

// InRange checks if an integer is within range
func InRange(value, min, max int) bool {
	return value >= min && value <= max
}

// OneOf checks if a string is one of the allowed values
func OneOf(value string, allowed []string) bool {
	for _, a := range allowed {
		if value == a {
			return true
		}
	}
	return false
}

// Validator provides a fluent interface for validation
type Validator struct {
	errors ValidationErrors
}

// NewValidator creates a new Validator
func NewValidator() *Validator {
	return &Validator{
		errors: make(ValidationErrors, 0),
	}
}

// AddError adds a validation error
func (v *Validator) AddError(field, message string) {
	v.errors = append(v.errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// Required validates that a field is not empty
func (v *Validator) Required(field, value string) *Validator {
	if !Required(value) {
		v.AddError(field, "is required")
	}
	return v
}

// Email validates that a field is a valid email
func (v *Validator) Email(field, value string) *Validator {
	if value != "" && !IsValidEmail(value) {
		v.AddError(field, "must be a valid email address")
	}
	return v
}

// Username validates that a field is a valid username
func (v *Validator) Username(field, value string) *Validator {
	if value != "" && !IsValidUsername(value) {
		v.AddError(field, "must be 3-30 characters, alphanumeric and underscores only")
	}
	return v
}

// MinLen validates minimum length
func (v *Validator) MinLen(field, value string, min int) *Validator {
	if value != "" && !MinLength(value, min) {
		v.AddError(field, fmt.Sprintf("must be at least %d characters", min))
	}
	return v
}

// MaxLen validates maximum length
func (v *Validator) MaxLen(field, value string, max int) *Validator {
	if value != "" && !MaxLength(value, max) {
		v.AddError(field, fmt.Sprintf("must not exceed %d characters", max))
	}
	return v
}

// Range validates integer range
func (v *Validator) Range(field string, value, min, max int) *Validator {
	if !InRange(value, min, max) {
		v.AddError(field, fmt.Sprintf("must be between %d and %d", min, max))
	}
	return v
}

// OneOf validates that value is one of allowed values
func (v *Validator) OneOf(field, value string, allowed []string) *Validator {
	if value != "" && !OneOf(value, allowed) {
		v.AddError(field, fmt.Sprintf("must be one of: %s", strings.Join(allowed, ", ")))
	}
	return v
}

// IsValid returns true if there are no validation errors
func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

// Errors returns all validation errors
func (v *Validator) Errors() ValidationErrors {
	return v.errors
}

// Error returns validation errors as an HTTPError
func (v *Validator) Error() *HTTPError {
	if v.IsValid() {
		return nil
	}
	return ErrValidation.WithDetails(v.errors)
}