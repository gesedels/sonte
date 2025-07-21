// Package neat implements string sanitisation and conversion functions.
package neat

import (
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Body returns a whitespace-trimmed note body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Hash returns a base64 SHA256 hash of a string.
func Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return base64.RawURLEncoding.EncodeToString(hash[:])
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

// Time returns a Time object from a Unix UTC integer string.
func Time(unix string) time.Time {
	unix = strings.TrimSpace(unix)
	uint, _ := strconv.ParseInt(unix, 10, 64)
	return time.Unix(uint, 0)
}

// Unix returns a Unix UTC integer string from a Time object.
func Unix(tobj time.Time) string {
	uint := tobj.Unix()
	return strconv.FormatInt(uint, 10)
}
