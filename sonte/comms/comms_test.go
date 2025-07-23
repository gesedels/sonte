package comms

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/gesedels/sonte/sonte/items/book"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	Commands["test"] = func(w io.Writer, _ *book.Book, elems []string) error {
		fmt.Fprintf(w, "elems=%v\n", elems)
		return nil
	}

	// success
	err := Run(w, nil, []string{"test", "alpha"})
	assert.Equal(t, "elems=[alpha]\n", w.String())
	assert.NoError(t, err)

	// error - no command provided
	err = Run(nil, nil, nil)
	assert.EqualError(t, err, `cannot run command - none provided`)

	// error - command does not exist
	err = Run(nil, nil, []string{"nope"})
	assert.EqualError(t, err, `cannot run command "nope" - does not exist`)
}
