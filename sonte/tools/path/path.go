// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"slices"
	"strings"
)

// Base returns a file path's name with the extension.
func Base(orig string) string {
	return filepath.Base(orig)
}

// Dire returns a file path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a file path's extension with a leading dot.
func Extn(orig string) string {
	base := filepath.Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[clip:]
	}

	return ""
}

// Glob returns all file paths in a directory matching an extension in alphabetical order.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	origs, _ := filepath.Glob(glob)
	slices.Sort(origs)
	return origs
}

// Join returns a joined file path from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Match returns true if a file path's name contains a substring.
func Match(orig, text string) bool {
	name := Name(orig)
	name = strings.ToLower(name)
	text = strings.ToLower(text)
	return strings.Contains(name, text)
}

// Name returns a file path's name without the extension.
func Name(orig string) string {
	base := filepath.Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[:clip]
	}

	return base
}
