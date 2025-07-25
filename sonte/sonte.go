package main

import (
	"fmt"
	"os"

	"github.com/gesedels/sonte/sonte/comms"
	"github.com/gesedels/sonte/sonte/items/book"
)

func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	book, err := book.NewEnv("SONTE_DIR", "SONTE_EXT", 0666)
	try(err)
	try(comms.Run(os.Stdout, book, os.Args[1:]))
}
