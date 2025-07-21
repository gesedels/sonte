# Sonte

**Sonte** (*Stephen's Obsessive Note-Taking Engine*) is a command-line plaintext note management suite, written in [Go 1.24][go] by [Stephen Malone][sm].

- See [`changes.md`][ch] for the full changelog.
- See [`license.md`][li] for the open-source license (BSD-3).

## Installation

You can install Sonte using your Go tools...

```
go install github.com/gesedels/sonte@latest
```

...or download the [latest release][lr] for your platform.

## Configuration

Sonte stores all data in a single directory in one of three locations, depending on which environment variables you have set:

- If `$SONTE_DIR` is set, the directory is used exactly.
- If `$XDG_CONFIG_HOME` is set, the `sonte` subdirectory is used.
- Otherwise, `$HOME` is used with dotfiles (e.g.: `.sonte.db`).

## Structure & Syntax

Sonte's data is stored in a [Bolt][bb] database, where each note is accessed by name and has the following stored fields:

Field  | Description
------ | -----------
`body` | A whitespace-trimmed string of the note's contents.
`hash` | A hexidecimal SHA256 hash of the note's contents.
`tags` | A space-separated list of every hashtag in the note.
`time` | A unix UTC integer of the note's creation time.

Note names and hashtags are always lowercase alphanumeric with dashes, so trying to create `My_Note_123` will result in `my-note-123`. Hashtags in notes must include the octothorpe (e.g.: `#foo`), but in commands it's optional.

## Commands

### Write a note

The `open` command will create the scratch file `sonte.temp` in the data directory, open it in your default editor (according to `$EDITOR` or `$VISUAL`) and save the contents as the named note:

```text
$ export EDITOR=vim
$ sonte open foo
```

- If the note exists it will be overwritten.
- The scratch file is cleared before opening and left intact after.

## Contributions

Please send all bug reports and feature requests to the [issue tracker][it], thank you.

[bb]: https://github.com/etcd-io/bbolt
[ch]: https://github.com/gesedels/sonte/blob/main/changes.md
[li]: https://github.com/gesedels/sonte/blob/main/license.md
[go]: https://go.dev/doc/go1.24
[it]: https://github.com/gesedels/sonte/issues
[lr]: https://github.com/gesedels/sonte/releases/latest
[sm]: https://github.com/gesedels
