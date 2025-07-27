// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a whitespace-trimmed file body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Dire returns a clean directory path.
func Dire(dire string) string {
	dire = strings.TrimSpace(dire)
	return filepath.Clean(dire)
}

// Extn returns a lowercase file extension string with a leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	extn = strings.TrimSpace(extn)
	return "." + strings.TrimLeft(extn, ".")
}

// Name returns an lowercase alphanumeric-with-dashes file name string.
func Name(name string) string {
	var chars []rune

	for _, char := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(char) || unicode.IsNumber(char):
			chars = append(chars, char)
		case unicode.IsSpace(char) || char == '_' || char == '-':
			chars = append(chars, '-')
		}
	}

	return strings.Trim(string(chars), "-")
}
