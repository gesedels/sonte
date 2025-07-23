// Package comms implements user-facing command functions.
package comms

import (
	"flag"
	"fmt"
	"io"

	"github.com/gesedels/sonte/sonte/items/book"
)

// usage prints a FlagSet's default usage and returns a nil error.
func usage(w io.Writer, fset *flag.FlagSet) error {
	fmt.Fprintf(w, "Usage of %s:\n", fset.Name())
	fset.SetOutput(w)
	fset.PrintDefaults()
	return nil
}

// Run parsed and executes command-line arguments.
func Run(w io.Writer, book *book.Book, elems []string) error {
	fset := flag.NewFlagSet("sonte", flag.ExitOnError)
	help := fset.Bool("h", false, "show help")
	fset.Parse(elems)

	switch {
	case *help:
		return usage(w, fset)

	case fset.NArg() != 0:
		return CommandOpen(w, book, flag.Arg(0))

	default:
		return usage(w, fset)
	}
}
