package log

import (
	"errors"
	"math"
)

// Level is a logging level.
type Level int

const (
	// DebugLevel is the debug level.
	DebugLevel Level = -4
	// InfoLevel is the info level.
	InfoLevel Level = 0
	// WarnLevel is the warn level.
	WarnLevel Level = 4
	// ErrorLevel is the error level.
	ErrorLevel Level = 8
	// FatalLevel is the fatal level.
	FatalLevel Level = 12
	// noLevel is used with log.Print.
	noLevel Level = math.MaxInt
)

// String returns the string representation of the level.
func (l Level) String() string { _ = "STUB: not implemented"; return "" }

// ErrInvalidLevel is an error returned when parsing an invalid level string.
var ErrInvalidLevel = errors.New("invalid level")

// ParseLevel converts level in string to Level type. Default level is InfoLevel.
func ParseLevel(level string) (Level, error) { _ = "STUB: not implemented"; return *new(Level), nil }
