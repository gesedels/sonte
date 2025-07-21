// Package neat implements string sanitisation and conversion functions.
package neat

import (
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// Body returns a whitespace-trimmed note body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Date returns a Time object from a YYYY-MM-DD date string.
func Date(date string) time.Time {
	date = strings.TrimSpace(date)
	tobj, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Unix(0, 0)
	}

	return tobj
}

// Name returns a lowercase alphanumeric-with-dashes note name string.
func Name(name string) string {
	var chars []rune
	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(char) || unicode.IsNumber(char):
			chars = append(chars, char)
		case unicode.IsSpace(char) || char == '_':
			chars = append(chars, '-')
		}
	}

	return strings.Trim(string(chars), "-")
}

// Path returns a cleaned file path.
func Path(path string) string {
	return filepath.Clean(path)
}
