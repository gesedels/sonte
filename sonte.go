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

/// note type and funcs ///

type Note struct {
	Time time.Time `json:"time"`
	Body string    `json:"body"`
}

func saveNote(dire string, note *Note) error {
	base := fmt.Sprintf("%d.json", note.Time.Unix())
	path := filepath.Join(dire, base)
	bytes, err := json.MarshalIndent(note, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, 0666)
}

func loadNote(orig string) (*Note, error) {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		return nil, err
	}

	note := new(Note)
	if err := json.Unmarshal(bytes, note); err != nil {
		return nil, err
	}

	return note, nil
}

/// main funcs ///

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dire := "."
	temp := filepath.Join(dire, "sonte.txt")

	if len(os.Args) <= 1 {
		edit := os.Getenv("EDITOR")
		look, err := exec.LookPath(edit)

		comm := exec.Command(look, temp)
		comm.Stdin = os.Stdin
		comm.Stdout = os.Stdout
		comm.Stderr = os.Stderr
		try(err)
		try(comm.Run())

		bytes, err := os.ReadFile(temp)
		try(err)
		if len(bytes) != 0 {
			note := &Note{time.Now(), string(bytes)}
			saveNote(dire, note)
		}

	} else {
		var tags []string
		for _, arg := range os.Args[1:] {
			tags = append(tags, "#"+strings.ToLower(arg))
		}

		paths, err := filepath.Glob(filepath.Join(dire, "*.json"))
		try(err)

	checknote:
		for _, path := range paths {
			note, err := loadNote(path)
			try(err)

			for _, tag := range tags {
				if strings.Contains(note.Body, tag) {
					fmt.Println(note.Time)
					fmt.Println(note.Body)
					fmt.Println("-----")
					continue checknote
				}
			}
		}
	}
}
