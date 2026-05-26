package log

import (
	"io"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

const (
	separator       = "="
	indentSeparator = "  │ "
)

func (l *Logger) writeIndent(w io.Writer, str string, indent string, newline bool, key string) {
	_ = "STUB: not implemented"

	// kindly borrowed from hclog
	return
}

//nolint:nestif

func needsEscaping(str string) bool { _ = "STUB: not implemented"; return false }

const (
	lowerhex = "0123456789abcdef"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func escapeStringForOutput(str string, escapeQuotes bool) string {
	_ = "STUB: not implemented"
	// kindly borrowed from hclog
	return ""
}

func needsQuoting(s string) bool { _ = "STUB: not implemented"; return false }

var needsQuotingSet = [utf8.RuneSelf]bool{
	'"': true,
	'=': true,
}

func init() {
	for i := 0; i < utf8.RuneSelf; i++ {
		r := rune(i)
		if unicode.IsSpace(r) || !unicode.IsPrint(r) {
			needsQuotingSet[i] = true
		}
	}
}

func writeSpace(w io.Writer, first bool) { _ = "STUB: not implemented"; return }

func (l *Logger) textFormatter(keyvals ...any) { _ = "STUB: not implemented"; return }

// Values may contain multiple lines, and that format
// is preserved, with each line prefixed with a "  | "
// to show it's part of a collection of lines.
//
// Values may also need quoting, if not all the runes
// in the value string are "normal", like if they
// contain ANSI escape sequences.

// Add a newline to the end of the log message.
