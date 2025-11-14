package httpx

import (
	"context"
	"net/http"
)

// SuccessResponse represents a standardized success response
type SuccessResponse struct {
	// Data contains the response payload
	Data interface{} `json:"data"`
	
	// Meta contains metadata about the response (pagination, etc.)
	Meta interface{} `json:"meta,omitempty"`
	
	// RequestID for tracing
	RequestID string `json:"request_id,omitempty"`
}

// WriteSuccess writes a successful JSON response
func WriteSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	requestID := GetRequestID(r.Context())
	
	resp := SuccessResponse{
		Data:      data,
		RequestID: requestID,
	}
	
	WriteJSON(w, http.StatusOK, resp)
}

// WriteSuccessWithMeta writes a successful JSON response with metadata
func WriteSuccessWithMeta(w http.ResponseWriter, r *http.Request, data interface{}, meta interface{}) {
	requestID := GetRequestID(r.Context())
	
	resp := SuccessResponse{
		Data:      data,
		Meta:      meta,
		RequestID: requestID,
	}
	
	WriteJSON(w, http.StatusOK, resp)
}

// WriteCreated writes a 201 Created response
func WriteCreated(w http.ResponseWriter, r *http.Request, data interface{}) {
	requestID := GetRequestID(r.Context())
	
	resp := SuccessResponse{
		Data:      data,
		RequestID: requestID,
	}
	
	WriteJSON(w, http.StatusCreated, resp)
}

// WriteNoContent writes a 204 No Content response
func WriteNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// WriteAccepted writes a 202 Accepted response
func WriteAccepted(w http.ResponseWriter, r *http.Request, data interface{}) {
	requestID := GetRequestID(r.Context())
	
	resp := SuccessResponse{
		Data:      data,
		RequestID: requestID,
	}
	
	WriteJSON(w, http.StatusAccepted, resp)
}

// contextKey for request ID
type contextKey string

const requestIDKey contextKey = "http:request_id"

// WithRequestID adds a request ID to the context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestID retrieves the request ID from context
func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return ""
}