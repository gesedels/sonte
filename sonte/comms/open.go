package comms

import (
	"io"

	"github.com/gesedels/sonte/sonte/items/book"
	"github.com/gesedels/sonte/sonte/tools/clui"
)

// CommandOpen opens a new or existing note in $EDITOR.
func CommandOpen(w io.Writer, book *book.Book, name string) error {
	prog, err := clui.Env("EDITOR")
	if err != nil {
		return err
	}

	note, err := book.GetOrCreate(name, "")
	if err != nil {
		return err
	}

	return clui.Exec(prog, note.Orig)
}
