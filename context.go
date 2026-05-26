package log

import "context"

// WithContext wraps the given logger in context.
func WithContext(ctx context.Context, logger *Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext returns the logger from the given context.
// This will return the default package logger if no logger
// found in context.
func FromContext(ctx context.Context) *Logger { _ = "STUB: not implemented"; return nil }

type contextKey struct{ string }

// ContextKey is the key used to store the logger in context.
var ContextKey = contextKey{"log"}
