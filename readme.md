# Sonte

**Sonte** (*Stephen's Obsessive Note-Taking Engine*) is a command-line note management system, written in [Go 1.24][go] by [Stephen Malone][sm]. If you have a directory full of plaintext note files you need to handle, Sonte can help you.

- See [`changes.md`][ch] for the full changelog.
- See [`license.md`][li] for the open-source license (BSD-3).

## Installation

You can install Sonte using your Go tools...

```
go install github.com/gesedels/sonte@latest
```

...or download the [latest release][lr] for your platform.

## Configuration

Sonte uses two environment variables for configuration:

- `SONTE_DIR` is the directory your note files are in.
- `SONTE_EXT` is the extension your note files use (including the dot).

```fish
$ export SONTE_DIR = "$HOME/Notes"
$ export SONTE_EXT = ".txt"
```

That's it! That's all you need to do.

## Commands

### Basic Syntax

- Note names are always lowercase alphanumeric with dashes, so trying to create `My_Note_123` will result in `my-note-123`.

### Open a note

The `open` command will open a new or existing note in your default editor (according to `$EDITOR` or `$VISUAL`):

```fish
$ sonte open foo
# Opens "$HOME/Notes/foo.txt" in $EDITOR.
```

- If the note does not exist, it is created and left empty.

## Contributions

Please send all bug reports and feature requests to the [issue tracker][it], thank you.

[ch]: https://github.com/gesedels/sonte/blob/main/changes.md
[li]: https://github.com/gesedels/sonte/blob/main/license.md
[go]: https://go.dev/doc/go1.24
[it]: https://github.com/gesedels/sonte/issues
[lr]: https://github.com/gesedels/sonte/releases/latest
[sm]: https://github.com/gesedels
