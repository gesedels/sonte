///////////////////////////////////////////////////////////////////////////////////////
//                   sonte · stephen's obsessive note-taking engine                  //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

///////////////////////////////////////////////////////////////////////////////////////
//                       part one · value conversion functions                       //
///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////
//                           part two · json file functions                          //
///////////////////////////////////////////////////////////////////////////////////////

// ReadJSON returns an unmarshalled JSON value from a file.
func ReadJSON(orig string, jval any) error {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, jval)
}

// WriteJSON writes a marshalled JSON value to a file.
func WriteJSON(dest string, jval any) error {
	bytes, err := json.Marshal(jval)
	if err != nil {
		return err
	}

	return os.WriteFile(dest, bytes, 0666)
}

///////////////////////////////////////////////////////////////////////////////////////
//                   part three · command-line interface functions                   //
///////////////////////////////////////////////////////////////////////////////////////

// EditFile opens a file in $EDITOR, returning after exit.
func EditFile(orig string) error {
	name, ok := os.LookupEnv("EDITOR")
	if !ok {
		return fmt.Errorf("environment variable $EDITOR not set")
	}

	path, err := exec.LookPath(name)
	if err != nil {
		return err
	}

	comm := exec.Command(path, orig)
	comm.Stdin = os.Stdin
	comm.Stdout = os.Stdout
	comm.Stderr = os.Stderr
	return comm.Run()
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part four · type definitions                           //
///////////////////////////////////////////////////////////////////////////////////////

// Entry is a single JSON file containing a system entry.
type Entry struct {
	Time time.Time `json:"time"`
	Body string    `json:"body"`
}

// NewEntry returns a new Entry.
func NewEntry(body string) *Entry {
	return &Entry{time.Now(), strings.TrimSpace(body) + "\n"}
}

// OpenEntry returns an existing Entry from a file.
func OpenEntry(orig string) (*Entry, error) {
	var eobj = new(Entry)
	if err := ReadJSON(orig, eobj); err != nil {
		return nil, err
	}

	return eobj, nil
}

// Dest returns the Entry's destination filepath in a directory.
func (e *Entry) Dest(dire string) string {
	base := fmt.Sprintf("%d.json", e.Time.Unix())
	return filepath.Join(dire, base)
}

// Match returns true if the Entry's body contains a hashtag.
func (e *Entry) Match(tags []string) bool {
	for _, tag := range tags {
		if strings.Contains(e.Body, tag) {
			return true
		}
	}

	return false
}

// Write writes the Entry to a file in a directory.
func (e *Entry) Write(dire string) error {
	dest := e.Dest(dire)
	return WriteJSON(dest, e)
}

///////////////////////////////////////////////////////////////////////////////////////
//                                      old code                                     //
///////////////////////////////////////////////////////////////////////////////////////

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dire := "."
	temp := filepath.Join(dire, "sonte.txt")
	try(os.WriteFile(temp, nil, 0666))

	if len(os.Args) <= 1 {
		try(EditFile(temp))
		bytes, err := os.ReadFile(temp)
		try(err)

		if len(bytes) != 0 {
			try(NewEntry(string(bytes)).Write(dire))
		}

	} else {
		var tags []string
		for _, arg := range os.Args[1:] {
			tags = append(tags, "#"+strings.ToLower(arg))
		}

		origs, err := filepath.Glob(filepath.Join(dire, "*.json"))
		try(err)

		for _, orig := range origs {
			entry, err := OpenEntry(orig)
			try(err)

			if entry.Match(tags) {
				fmt.Println(entry.Time)
				fmt.Println(entry.Body)
				fmt.Println("-----")
			}
		}
	}
}
