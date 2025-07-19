// Package errs implements error handling functions.
package errs

import "fmt"

// Error is a custom error with a manually wrapped sub-error.
type Error struct {
	Text string
	Wrap error
}

// New returns a new formatted Error.
func New(text string, elems ...any) *Error {
	form := fmt.Sprintf(text, elems...)
	return &Error{form, nil}
}

// Wrap returns a new formatted Error with a wrapped sub-error.
func Wrap(err error, text string, elems ...any) *Error {
	form := fmt.Sprintf(text, elems...)
	return &Error{form, err}
}

// Error returns the Error's message string.
func (e *Error) Error() string {
	return e.Text
}

// Unwrap returns the Error's wrapped sub-error.
func (e *Error) Unwrap() error {
	return e.Wrap
}
