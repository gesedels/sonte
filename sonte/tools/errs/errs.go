// Package errs implements error handling functions.
package errs

import (
	"fmt"
	"path/filepath"
)

// EvarDoesNotExist returns a custom error for non-existent environment variables.
func EvarDoesNotExist(name string) error {
	return fmt.Errorf("environment variable %q does not exist", name)
}

// FileAlreadyExists returns a custom error for pre-existing files.
func FileAlreadyExists(orig string) error {
	base := filepath.Base(orig)
	return fmt.Errorf("file %q already exists", base)
}

// FileDoesNotExist returns a custom error for non-existent files.
func FileDoesNotExist(orig string) error {
	base := filepath.Base(orig)
	return fmt.Errorf("file %q does not exist", base)
}

// FileSystemError returns a custom error for file system errors.
func FileSystemError(orig string, err error) error {
	base := filepath.Base(orig)
	return fmt.Errorf("file %q had system error: %w", base, err)
}

// ProgDoesNotExist  returns a custom error for non-existent programs.
func ProgDoesNotExist(name string) error {
	return fmt.Errorf("program %q does not exist", name)
}

// ProgSystemError  returns a custom error for program execution errors.
func ProgSystemError(name string, err error) error {
	return fmt.Errorf("program %q had system error: %w", name, err)
}
