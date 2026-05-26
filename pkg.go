package log

import (
	"io"
	"log"
	"sync"
	"sync/atomic"

	"github.com/charmbracelet/colorprofile"
)

var (
	// defaultLogger is the default global logger instance.
	defaultLogger     atomic.Pointer[Logger]
	defaultLoggerOnce sync.Once
)

// Default returns the default logger. The default logger comes with timestamp enabled.
func Default() *Logger { _ = "STUB: not implemented"; return nil }

// SetDefault sets the default global logger.
func SetDefault(logger *Logger) { _ = "STUB: not implemented"; return }

// New returns a new logger with the default options.
func New(w io.Writer) *Logger { _ = "STUB: not implemented"; return nil }

// NewWithOptions returns a new logger using the provided options.
func NewWithOptions(w io.Writer, o Options) *Logger { _ = "STUB: not implemented"; return nil }

// Detect color profile from the writer and environment.

// SetReportTimestamp sets whether to report timestamp for the default logger.
func SetReportTimestamp(report bool) { _ = "STUB: not implemented"; return }

// SetReportCaller sets whether to report caller location for the default logger.
func SetReportCaller(report bool) { _ = "STUB: not implemented"; return }

// SetLevel sets the level for the default logger.
func SetLevel(level Level) { _ = "STUB: not implemented"; return }

// GetLevel returns the level for the default logger.
func GetLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

// SetTimeFormat sets the time format for the default logger.
func SetTimeFormat(format string) { _ = "STUB: not implemented"; return }

// SetTimeFunction sets the time function for the default logger.
func SetTimeFunction(f TimeFunction) { _ = "STUB: not implemented"; return }

// SetOutput sets the output for the default logger.
func SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

// SetFormatter sets the formatter for the default logger.
func SetFormatter(f Formatter) { _ = "STUB: not implemented"; return }

// SetCallerFormatter sets the caller formatter for the default logger.
func SetCallerFormatter(f CallerFormatter) { _ = "STUB: not implemented"; return }

// SetCallerOffset sets the caller offset for the default logger.
func SetCallerOffset(offset int) { _ = "STUB: not implemented"; return }

// SetPrefix sets the prefix for the default logger.
func SetPrefix(prefix string) { _ = "STUB: not implemented"; return }

// SetColorProfile force sets the underlying color profile for the
// TextFormatter.
func SetColorProfile(profile colorprofile.Profile) { _ = "STUB: not implemented"; return }

// SetStyles sets the logger styles for the TextFormatter.
func SetStyles(s *Styles) { _ = "STUB: not implemented"; return }

// GetPrefix returns the prefix for the default logger.
func GetPrefix() string { _ = "STUB: not implemented"; return "" }

// With returns a new logger with the given keyvals.
func With(keyvals ...any) *Logger { _ = "STUB: not implemented"; return nil }

// WithPrefix returns a new logger with the given prefix.
func WithPrefix(prefix string) *Logger { _ = "STUB: not implemented"; return nil }

// Helper marks the calling function as a helper
// and skips it for source location information.
// It's the equivalent of testing.TB.Helper().
func Helper() { _ = "STUB: not implemented"; return }

// Log logs a message with the given level.
func Log(level Level, msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Debug logs a debug message.
func Debug(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Info logs an info message.
func Info(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Warn logs a warning message.
func Warn(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Error logs an error message.
func Error(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Fatal logs a fatal message and exit.
func Fatal(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Print logs a message with no level.
func Print(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Logf logs a message with formatting and level.
func Logf(level Level, format string, args ...any) { _ = "STUB: not implemented"; return }

// Debugf logs a debug message with formatting.
func Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Infof logs an info message with formatting.
func Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

// Warnf logs a warning message with formatting.
func Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Errorf logs an error message with formatting.
func Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Fatalf logs a fatal message with formatting and exit.
func Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Printf logs a message with formatting and no level.
func Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

// StandardLog returns a standard logger from the default logger.
func StandardLog(opts ...StandardLogOptions) *log.Logger { _ = "STUB: not implemented"; return nil }
