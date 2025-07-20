///////////////////////////////////////////////////////////////////////////////////////
//                   sonte · stephen's obsessive note-taking engine                  //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

///////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                         //
///////////////////////////////////////////////////////////////////////////////////////

// 1.1: database schema constants
//////////////////////////////////

// Pragma is the default always-on database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = on;
`

// Schema is the default first-run database schema.
const Schema = `
	create table Notes (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		body text    not null
	);

	create table Tags (
		id   integer primary key asc,
		name text    not null,
		unique(name)
	);

	create table NoteTags (
		id   integer primary key asc,
		note integer not null references Notes(id),
		tag  integer not null references Tags(id),
		unique(note, tag)
	);
`

// 1.2: regular expressions
////////////////////////////

// Hashtags is a regular expression for finding hashtags.
var Hashtags = regexp.MustCompile(`(#\w+)`)

///////////////////////////////////////////////////////////////////////////////////////
//                            part ??? · command functions                           //
///////////////////////////////////////////////////////////////////////////////////////

// RunCreate opens a temporary file in $EDITOR and adds it to the database.
func RunCreate(db *sqlx.DB) error {
	// Create temporary file.
	file, err := os.CreateTemp("", "sonte-*.txt")
	if err != nil {
		return err
	}

	// Close and defer removal of temporary file.
	file.Close()
	defer os.Remove(file.Name())

	// Get $EDITOR variable.
	prog, ok := os.LookupEnv("EDITOR")
	if !ok {
		return fmt.Errorf("envvar $EDITOR not set")
	}

	// Look up $EDITOR path.
	path, err := exec.LookPath(prog)
	if err != nil {
		return err
	}

	// Build $EDITOR command.
	comm := exec.Command(path, file.Name())
	comm.Stdin = os.Stdin
	comm.Stdout = os.Stdout
	comm.Stderr = os.Stderr

	// Run $EDITOR against temporary file.
	if err := comm.Run(); err != nil {
		return err
	}

	// Read temporary file bytes and exit if blank.
	bytes, err := os.ReadFile(file.Name())
	switch {
	case err != nil:
		return err
	case len(bytes) == 0:
		return nil
	}

	// Insert written bytes into database.
	rslt, err := db.Exec("insert into Notes (body) values (?)", string(bytes))
	if err != nil {
		return err
	}

	// Remember note insert ID.
	n_id, err := rslt.LastInsertId()
	if err != nil {
		return err
	}

	// Find hashtags in written bytes.
	tags := Hashtags.FindAllString(string(bytes), -1)

	// Process hashtags in written bytes.
	for _, tag := range tags {
		// Sanitise hashtag.
		tag = strings.ToLower(tag)
		tag = strings.TrimLeft(tag, "#")

		// Upsert hashtag into database.
		_, err := db.Exec("insert into Tags (name) values (?) on conflict(name) do nothing", tag)
		if err != nil {
			return err
		}

		// Get tag row ID.
		var t_id int
		err = db.Get(&t_id, "select id from Tags where name=?", tag)
		if err != nil {
			return err
		}

		// Upsert hashtag link into database.
		_, err = db.Exec("insert into NoteTags (note, tag) values (?, ?) on conflict(note, tag) do nothing", n_id, t_id)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Created note with tags %#v.\n", tags)
	return nil
}

// RunList lists all notes containing a hashtag.
func RunList(db *sqlx.DB, tag string) error {
	// Sanitise hashtag.
	tag = strings.ToLower(tag)
	tag = strings.TrimLeft(tag, "#")

	// Get tag ID.
	var t_id int
	err := db.Get(&t_id, "select id from Tags where name=?", tag)
	if err != nil {
		return err
	}

	// Get note IDs.
	var n_ids []int
	err = db.Select(&n_ids, "select note from NoteTags where tag=?", t_id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// Iterate through note IDs.
	for _, n_id := range n_ids {
		var init int64
		var body string

		// Get note init time.
		err := db.Get(&init, "select init from Notes where id=?", n_id)
		if err != nil {
			return err
		}

		// Get note body string.
		err = db.Get(&body, "select body from Notes where id=?", n_id)
		if err != nil {
			return err
		}

		// Format init time.
		tnow := time.Unix(init, 0)
		tfmt := tnow.Format(time.DateTime)

		// Print note.
		fmt.Printf("# %s\n", tfmt)
		fmt.Println(strings.TrimSpace(body))
		fmt.Println("- - - - -")
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////
//                             part ??? · main functions                             //
///////////////////////////////////////////////////////////////////////////////////////

// try prints a non-nil error and exits.
func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

// main runs the main Sonte program.
func main() {
	// Connect to database.
	db, err := sqlx.Connect("sqlite3", "sonte.db")
	try(err)

	// Execute pragma on database.
	_, err = db.Exec(Pragma)
	try(err)

	// Check for table definitions in database.
	var size int
	try(db.Get(&size, "select count(*) from SQLITE_SCHEMA"))

	// Execute schema on database if no tables are defined.
	if size == 0 {
		_, err = db.Exec(Schema)
		try(err)
	}

	// Execute command.
	if len(os.Args) <= 1 {
		try(RunCreate(db))
	} else {
		try(RunList(db, os.Args[1]))
	}
}
