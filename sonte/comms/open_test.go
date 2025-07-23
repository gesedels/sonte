package comms

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandOpen(t *testing.T) {
	// setup
	w, book := mockWriterBook(t)
	os.Setenv("EDITOR", "touch")

	// success
	err := CommandOpen(w, book, "alpha")
	assert.NoError(t, err)
}
