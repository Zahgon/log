package log

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"

	"github.com/charmbracelet/colorprofile"
)

// ErrMissingValue is returned when a key is missing a value.
var ErrMissingValue = fmt.Errorf("missing value")

// LoggerOption is an option for a logger.
type LoggerOption = func(*Logger)

// Logger is a Logger that implements Logger.
type Logger struct {
	w  colorprofile.Writer
	b  bytes.Buffer
	mu *sync.RWMutex

	isDiscard uint32

	level           int64
	prefix          string
	timeFunc        TimeFunction
	timeFormat      string
	callerOffset    int
	callerFormatter CallerFormatter
	formatter       Formatter

	reportCaller    bool
	reportTimestamp bool

	fields []any

	helpers *sync.Map
	styles  *Styles
}

// Logf logs a message with formatting.
func (l *Logger) Logf(level Level, format string, args ...any) { _ = "STUB: not implemented"; return }

// Log logs the given message with the given keyvals for the given level.
func (l *Logger) Log(level Level, msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// check if the level is allowed

// Skip log.log, the caller, and any offset added.

// Found a frame that wasn't a helper function.
// Or we ran out of frames to check.

func (l *Logger) handle(level Level, ts time.Time, frames []runtime.Frame, msg any, keyvals ...any) {
	_ = "STUB: not implemented"
	return
}

// append logger fields

// append the rest

// WriteTo will reset the buffer

// Reset the buffer even if the lengths don't match up. If we're
// using colorprofile's Writer, it will strip the ansi sequences based on
// the color profile which can cause this error.

// Helper marks the calling function as a helper
// and skips it for source location information.
// It's the equivalent of testing.TB.Helper().
func (l *Logger) Helper() { _ = "STUB: not implemented"; return }

func (l *Logger) helper(skip int) {
	_ = "STUB: not implemented"

	// Skip runtime.Callers, and l.helper
	return
}

// frames returns the runtime.Frames for the caller.
func (l *Logger) frames(skip int) *runtime.Frames {
	_ = "STUB: not implemented"
	// Copied from testing.T
	return nil
}

// Skip runtime.Callers, and l.frame

func (l *Logger) location(frames []runtime.Frame) (file string, line int, fn string) {
	_ = "STUB: not implemented"
	return "", 0, ""
}

// Cleanup a path by returning the last n segments of the path only.
func trimCallerPath(path string, n int) string {
	_ = "STUB: not implemented"
	// lovely borrowed from zap
	// nb. To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	return ""
}

// Return the full path if n is 0.

// Find the last separator.

// Find the penultimate separator.

// SetReportTimestamp sets whether the timestamp should be reported.
func (l *Logger) SetReportTimestamp(report bool) { _ = "STUB: not implemented"; return }

// SetReportCaller sets whether the caller location should be reported.
func (l *Logger) SetReportCaller(report bool) { _ = "STUB: not implemented"; return }

// GetLevel returns the current level.
func (l *Logger) GetLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

// SetLevel sets the current level.
func (l *Logger) SetLevel(level Level) { _ = "STUB: not implemented"; return }

// GetPrefix returns the current prefix.
func (l *Logger) GetPrefix() string { _ = "STUB: not implemented"; return "" }

// SetPrefix sets the current prefix.
func (l *Logger) SetPrefix(prefix string) { _ = "STUB: not implemented"; return }

// SetTimeFormat sets the time format.
func (l *Logger) SetTimeFormat(format string) { _ = "STUB: not implemented"; return }

// SetTimeFunction sets the time function.
func (l *Logger) SetTimeFunction(f TimeFunction) { _ = "STUB: not implemented"; return }

// SetOutput sets the output destination.
func (l *Logger) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

// SetColorProfile force sets the underlying color profile for the
// TextFormatter.
func (l *Logger) SetColorProfile(profile colorprofile.Profile) { _ = "STUB: not implemented"; return }

// SetFormatter sets the formatter.
func (l *Logger) SetFormatter(f Formatter) { _ = "STUB: not implemented"; return }

// SetCallerFormatter sets the caller formatter.
func (l *Logger) SetCallerFormatter(f CallerFormatter) { _ = "STUB: not implemented"; return }

// SetCallerOffset sets the caller offset.
func (l *Logger) SetCallerOffset(offset int) { _ = "STUB: not implemented"; return }

// SetStyles sets the logger styles for the TextFormatter.
func (l *Logger) SetStyles(s *Styles) { _ = "STUB: not implemented"; return }

// With returns a new logger with the given keyvals added.
func (l *Logger) With(keyvals ...any) *Logger { _ = "STUB: not implemented"; return nil }

// WithPrefix returns a new logger with the given prefix.
func (l *Logger) WithPrefix(prefix string) *Logger { _ = "STUB: not implemented"; return nil }

// Debug prints a debug message.
func (l *Logger) Debug(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Info prints an info message.
func (l *Logger) Info(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Warn prints a warning message.
func (l *Logger) Warn(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Error prints an error message.
func (l *Logger) Error(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Fatal prints a fatal message and exits.
func (l *Logger) Fatal(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Print prints a message with no level.
func (l *Logger) Print(msg any, keyvals ...any) { _ = "STUB: not implemented"; return }

// Debugf prints a debug message with formatting.
func (l *Logger) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Infof prints an info message with formatting.
func (l *Logger) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

// Warnf prints a warning message with formatting.
func (l *Logger) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Errorf prints an error message with formatting.
func (l *Logger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Fatalf prints a fatal message with formatting and exits.
func (l *Logger) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Printf prints a message with no level and formatting.
func (l *Logger) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }
