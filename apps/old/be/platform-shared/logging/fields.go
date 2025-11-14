package logging

// Common field names for consistent logging across services
const (
	// FieldRequestID is the request ID field name
	FieldRequestID = "request_id"
	
	// FieldUserID is the user ID field name
	FieldUserID = "user_id"
	
	// FieldTraceID is the distributed trace ID field name
	FieldTraceID = "trace_id"
	
	// FieldSpanID is the span ID within a trace
	FieldSpanID = "span_id"
	
	// FieldService is the service name field
	FieldService = "service"
	
	// FieldMethod is the HTTP method field
	FieldMethod = "method"
	
	// FieldPath is the HTTP path field
	FieldPath = "path"
	
	// FieldStatusCode is the HTTP status code field
	FieldStatusCode = "status_code"
	
	// FieldDuration is the request duration field
	FieldDuration = "duration_ms"
	
	// FieldError is the error field name
	FieldError = "error"
	
	// FieldErrorType is the error type field
	FieldErrorType = "error_type"
	
	// FieldErrorStack is the error stack trace field
	FieldErrorStack = "error_stack"
	
	// FieldEventType is the event type field (for event-driven systems)
	FieldEventType = "event_type"
	
	// FieldEventID is the event ID field
	FieldEventID = "event_id"
	
	// FieldAggregateID is the aggregate ID field (DDD)
	FieldAggregateID = "aggregate_id"
	
	// FieldClientIP is the client IP address field
	FieldClientIP = "client_ip"
	
	// FieldUserAgent is the user agent field
	FieldUserAgent = "user_agent"
)