//go:build !go1.21

package log

import (
	"context"

	"golang.org/x/exp/slog"
)

// type alises for slog.
type (
	slogAttr      = slog.Attr
	slogValue     = slog.Value
	slogLogValuer = slog.LogValuer
)

var slogAnyValue = slog.AnyValue

const slogKindGroup = slog.KindGroup

// Enabled reports whether the logger is enabled for the given level.
//
// Implements slog.Handler.
func (l *Logger) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

// Handle handles the Record. It will only be called if Enabled returns true.
//
// Implements slog.Handler.
func (l *Logger) Handle(_ context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the caller frame using the record's PC.

// WithAttrs returns a new Handler with the given attributes added.
//
// Implements slog.Handler.
func (l *Logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// WithGroup returns a new Handler with the given group name prepended to the
// current group name or prefix.
//
// Implements slog.Handler.
func (l *Logger) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

var _ slog.Handler = (*Logger)(nil)
