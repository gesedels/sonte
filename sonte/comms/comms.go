// Package comms implements the Command type and functions.
package comms

import (
	"fmt"
	"io"

	"github.com/gesedels/sonte/sonte/items/book"
	"github.com/gesedels/sonte/sonte/tools/neat"
)

// Command is a user-facing command-line operation.
type Command func(io.Writer, *book.Book, []string) error

// Commands is a map of all existing Commands.
var Commands = map[string]Command{
	// "open": CommandOpen,
}

// Run discovers and executes a Command with arguments.
func Run(w io.Writer, book *book.Book, elems []string) error {
	if len(elems) == 0 {
		return fmt.Errorf("cannot run command - none provided")
	}

	name := neat.Name(elems[0])
	cfun, ok := Commands[name]
	if !ok {
		return fmt.Errorf("cannot run command %q - does not exist", name)
	}

	return cfun(w, book, elems[1:])
}
