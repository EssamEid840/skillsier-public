package logging

import "context"

type contextKey string

const loggerKey contextKey = "logging:logger"

// WithLogger adds a Logger to the context
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext retrieves the Logger from the context
// Returns a default logger if none is found
func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		// Return default logger
		return New(DefaultConfig("default"))
	}
	return logger
}

// MustFromContext retrieves the Logger from the context
// Panics if no logger is found
func MustFromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		panic("logging: no logger in context")
	}
	return logger
}