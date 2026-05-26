package log

import (
	"log"
)

type stdLogWriter struct {
	l   *Logger
	opt *StandardLogOptions
}

func (l *stdLogWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:exhaustive

// StandardLogOptions can be used to configure the standard log adapter.
type StandardLogOptions struct {
	ForceLevel Level
}

// StandardLog returns a standard logger from Logger. The returned logger
// can infer log levels from message prefix. Expected prefixes are DEBUG, INFO,
// WARN, ERROR, and ERR.
func (l *Logger) StandardLog(opts ...StandardLogOptions) *log.Logger {
	_ = "STUB: not implemented"

	// The caller stack is
	// log.Printf() -> l.Output() -> l.out.Write(stdLogger.Write)
	return nil
}
