// Package comms implements user-facing command functions.
package comms

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/gesedels/sonte/sonte/items/book"
)

// Usage is the default program help message.
const Usage = `
Sonte: Stephen's Obsessive Note-Taking Engine.

Command-line flags:
%s
See github.com/gesedels/sonte for more information.
`

// Run parsed and executes command-line arguments.
func Run(w io.Writer, book *book.Book, elems []string) error {
	fset := flag.NewFlagSet("sonte", flag.ExitOnError)
	help := fset.Bool("h", false, "show help")
	fset.Parse(elems)

	switch {
	case *help:
		WriteUsage(w, fset)
		return nil

	case fset.NArg() != 0:
		return CommandOpen(w, book, flag.Arg(0))

	default:
		WriteUsage(w, fset)
		return nil
	}
}

// WriteUsage writes a formatted usage string to a Writer.
func WriteUsage(w io.Writer, fset *flag.FlagSet) {
	b := bytes.NewBuffer(nil)
	fset.SetOutput(b)
	fset.PrintDefaults()
	fmt.Fprintf(w, strings.TrimLeft(Usage, "\n"), b.String())
}
