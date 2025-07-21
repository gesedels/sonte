# Sonte

**Sonte** (*Stephen's Obsessive Note-Taking Engine*) is a command-line plaintext note management suite, written in [Go 1.24][go] by [Stephen Malone][sm].

- See [`changes.md`][ch] for the full changelog.
- See [`license.md`][li] for the open-source license (BSD-3).

## Installation

You can install Sonte using your Go tools...

```
go install github.com/gesedels/sonte@latest
```

or download the [latest release][lr] for your platform.

## Configuration

Sonte stores all data in a single database file, placed in one of three locations depending on which environment variables you have set:

Variable          | Location
----------------- | --------
`SONTE_DIR`       | `$SONTE_DIR/sonte.db`
`XDG_CONFIG_HOME` | `$XDG_CONFIG_HOME/sonte/sonte.db`
`$HOME`           | `$HOME/.sonte`

If you're using `SONTE_DIR`, set it to the absolute path of an existing directory.

## Database Structure

Sonte's data is stored in a [Bolt][bb] database, where each note is accessed by name and has the following stored fields:

Field  | Description
------ | -----------
`body` | A whitespace-trimmed string of the note's contents.
`hash` | A hexidecimal SHA256 hash of the note's contents.
`tags` | A space-separated list of all the hashtags in the note.
`time` | A unix UTC integer of the note's creation time.

## Commands

### General Syntax

- Note names are always lowercase alphanumeric with dashes. Trying to create `My_Note_123` will result in `my-note-123`.
- Hashtags are also always lowercase alphanumeric with dashes. Hashtags in notes must include the octothorpe (e.g.: `#foo`), but in commands it is optional.

### `find TEXT`

List all notes in the database containing `TEXT` (case-insensitive).

```text
$ sonte find foo
bands-i-like
list-of-placeholders
```

### `list [TEXT]`

List all notes in the database with names containing `TEXT` (case-insensitive). If `TEXT` is blank, list all existing notes.

```text
$ sonte list
bands-i-like
list-of-placeholders
third-things

$ sonte list bands
bands-i-like
```

### `open NOTE`

Open `NOTE` in a temporary file your default terminal editor (according to `EDITOR` or `VISUAL`). When finished, the note is added to the database.

```text
$ sonte open new-note
# opens in $EDITOR
```

### `read NOTE`

Print the contents of `NOTE` to standard output. If the note does not exist, print nothing.

### `tags NAME`

List all notes containing hashtag `TAG`,

## Contributions

Please send all bug reports and feature requests to the [issue tracker][it], thank you.

[bb]: https://github.com/etcd-io/bbolt
[ch]: https://github.com/gesedels/sonte/blob/main/changes.md
[li]: https://github.com/gesedels/sonte/blob/main/license.md
[go]: https://go.dev/doc/go1.24
[it]: https://github.com/gesedels/sonte/issues
[lr]: https://github.com/gesedels/sonte/releases/latest
[sm]: https://github.com/gesedels
