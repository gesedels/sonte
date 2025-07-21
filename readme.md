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

*TODO: Bolt DB, show fields table (time, body, hash, tags).*

## Commands

*TODO: find [text], list [text=""], open [name] (or sonte [name] shortcut), read [name], tags [tags...].*

## Contributions

Please send all bug reports and feature requests to the [issue tracker][it], thank you.

[ch]: https://github.com/gesedels/sonte/blob/main/changes.md
[li]: https://github.com/gesedels/sonte/blob/main/license.md
[go]: https://go.dev/doc/go1.24
[it]: https://github.com/gesedels/sonte/issues
[lr]: https://github.com/gesedels/sonte/releases/latest
[sm]: https://github.com/gesedels
