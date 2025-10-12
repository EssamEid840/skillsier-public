package httpx

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	// Code is a machine-readable error code
	Code string `json:"code"`
	
	// Message is a human-readable error message
	Message string `json:"message"`
	
	// Details provides additional error context (optional)
	Details interface{} `json:"details,omitempty"`
	
	// RequestID helps with tracing and debugging
	RequestID string `json:"request_id,omitempty"`
	
	// Timestamp of the error
	Timestamp string `json:"timestamp,omitempty"`
}

// HTTPError represents an HTTP error with status code
type HTTPError struct {
	StatusCode int
	Code       string
	Message    string
	Details    interface{}
}

// Error implements the error interface
func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError creates a new HTTP error
func NewHTTPError(statusCode int, code, message string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

// WithDetails adds details to the error
func (e *HTTPError) WithDetails(details interface{}) *HTTPError {
	e.Details = details
	return e
}

// Common HTTP errors
var (
	// 400 Bad Request
	ErrBadRequest = NewHTTPError(http.StatusBadRequest, "bad_request", "The request is invalid")
	ErrValidation = NewHTTPError(http.StatusBadRequest, "validation_error", "Validation failed")
	
	// 401 Unauthorized
	ErrUnauthorized = NewHTTPError(http.StatusUnauthorized, "unauthorized", "Authentication required")
	ErrInvalidCredentials = NewHTTPError(http.StatusUnauthorized, "invalid_credentials", "Invalid username or password")
	ErrInvalidToken = NewHTTPError(http.StatusUnauthorized, "invalid_token", "Token is invalid or expired")
	
	// 403 Forbidden
	ErrForbidden = NewHTTPError(http.StatusForbidden, "forbidden", "You don't have permission to access this resource")
	ErrInsufficientPermissions = NewHTTPError(http.StatusForbidden, "insufficient_permissions", "You lack the required permissions")
	
	// 404 Not Found
	ErrNotFound = NewHTTPError(http.StatusNotFound, "not_found", "Resource not found")
	
	// 409 Conflict
	ErrConflict = NewHTTPError(http.StatusConflict, "conflict", "Resource already exists or conflicts with existing data")
	ErrDuplicateEntry = NewHTTPError(http.StatusConflict, "duplicate_entry", "A resource with this identifier already exists")
	
	// 422 Unprocessable Entity
	ErrUnprocessableEntity = NewHTTPError(http.StatusUnprocessableEntity, "unprocessable_entity", "The request is well-formed but cannot be processed")
	
	// 429 Too Many Requests
	ErrTooManyRequests = NewHTTPError(http.StatusTooManyRequests, "too_many_requests", "Rate limit exceeded")
	
	// 500 Internal Server Error
	ErrInternal = NewHTTPError(http.StatusInternalServerError, "internal_error", "An internal server error occurred")
	ErrDatabaseError = NewHTTPError(http.StatusInternalServerError, "database_error", "A database error occurred")
	
	// 503 Service Unavailable
	ErrServiceUnavailable = NewHTTPError(http.StatusServiceUnavailable, "service_unavailable", "Service temporarily unavailable")
	ErrDependencyFailure = NewHTTPError(http.StatusServiceUnavailable, "dependency_failure", "A dependent service is unavailable")
)

// WriteError writes an error response to the HTTP response writer
func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	// Get request ID from context if available
	requestID := GetRequestID(r.Context())
	
	// Check if it's an HTTPError
	httpErr, ok := err.(*HTTPError)
	if !ok {
		// Default to internal server error
		httpErr = ErrInternal
	}
	
	// Create error response
	errResp := ErrorResponse{
		Code:      httpErr.Code,
		Message:   httpErr.Message,
		Details:   httpErr.Details,
		RequestID: requestID,
	}
	
	// Write JSON response
	WriteJSON(w, httpErr.StatusCode, errResp)
}

// WriteJSON writes a JSON response
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// If encoding fails, write a plain text error
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}