package comms

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/gesedels/sonte/sonte/items/book"
	"github.com/stretchr/testify/assert"
)

func mockCommand(w io.Writer, _ *book.Book, elems []string) error {
	fmt.Fprintf(w, "%v\n", elems)
	return nil
}

func TestRun(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	Commands["test"] = mockCommand

	// success
	err := Run(w, nil, []string{"test", "alpha"})
	assert.Equal(t, "[alpha]\n", w.String())
	assert.NoError(t, err)

	// error - no command provided
	err = Run(nil, nil, nil)
	assert.EqualError(t, err, `cannot run command - none provided`)

	// error - command does not exist
	err = Run(nil, nil, []string{"nope"})
	assert.EqualError(t, err, `cannot run command "nope" - does not exist`)
}
