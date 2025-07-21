// Package neat implements string sanitisation and conversion functions.
package neat

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a whitespace-trimmed note body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
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
