package comms

import (
	"bytes"
	"flag"
	"testing"

	"github.com/gesedels/sonte/sonte/items/book"
	"github.com/gesedels/sonte/sonte/tools/test"
	"github.com/stretchr/testify/assert"
)

func mockWriterBook(t *testing.T) (*bytes.Buffer, *book.Book) {
	dire := test.MockDire(t)
	book := book.New(dire, ".extn", 0666)
	return bytes.NewBuffer(nil), book
}

func TestUsage(t *testing.T) {
	// setup
	w := bytes.NewBuffer(nil)
	fset := flag.NewFlagSet("test", flag.ExitOnError)
	fset.Bool("t", false, "test flag")

	// success
	err := usage(w, fset)
	assert.Equal(t, "Usage of test:\n  -t\ttest flag\n", w.String())
	assert.NoError(t, err)
}
