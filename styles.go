package log

import (
	"charm.land/lipgloss/v2"
)

// Styles defines the styles for the text logger.
type Styles struct {
	// Timestamp is the style for timestamps.
	Timestamp lipgloss.Style

	// Caller is the style for source caller.
	Caller lipgloss.Style

	// Prefix is the style for prefix.
	Prefix lipgloss.Style

	// Message is the style for messages.
	Message lipgloss.Style

	// Key is the style for keys.
	Key lipgloss.Style

	// Value is the style for values.
	Value lipgloss.Style

	// Separator is the style for separators.
	Separator lipgloss.Style

	// Levels are the styles for each level.
	Levels map[Level]lipgloss.Style

	// Keys overrides styles for specific keys.
	Keys map[string]lipgloss.Style

	// Values overrides value styles for specific keys.
	Values map[string]lipgloss.Style
}

// DefaultStyles returns the default styles.
func DefaultStyles() *Styles {
	_ = "STUB: not implemented"
	// TODO handle this based on light/dark colors
	return nil
}
