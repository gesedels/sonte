// Package clui implements command-line user interface functions.
package clui

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gesedels/sonte/sonte/tools/errs"
)

// Env returns the value of an existing environment variable.
func Env(name string) (string, error) {
	data, ok := os.LookupEnv(name)
	data = strings.TrimSpace(data)
	if !ok || data == "" {
		return "", errs.EvarDoesNotExist(name)
	}

	return data, nil
}

// Exec executes a shell command with arguments.
func Exec(prog string, elems ...string) error {
	path, err := exec.LookPath(prog)
	if err != nil {
		return errs.ProgDoesNotExist(prog)
	}

	comm := exec.Command(path, elems...)
	comm.Stdin = os.Stdin
	comm.Stdout = os.Stdout
	comm.Stderr = os.Stderr

	if err := comm.Run(); err != nil {
		return errs.ProgSystemError(prog, err)
	}

	return nil
}
