// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Env returns the value of an existing environment variable.
func Env(name string) (string, error) {
	data, ok := os.LookupEnv(name)
	data = strings.TrimSpace(data)

	switch {
	case !ok:
		return "", fmt.Errorf("environment variable %q is not set", name)
	case data == "":
		return "", fmt.Errorf("environment variable %q is empty", name)
	default:
		return data, nil
	}
}

// Exec executes a shell command with arguments.
func Exec(prog string, elems ...string) error {
	path, err := exec.LookPath(prog)
	if err != nil {
		return fmt.Errorf("cannot find path to program %q", prog)
	}

	comm := exec.Command(path, elems...)
	comm.Stdin = os.Stdin
	comm.Stdout = os.Stdout
	comm.Stderr = os.Stderr

	if err := comm.Run(); err != nil {
		return fmt.Errorf("failed to run program %q - %w", prog, err)
	}

	return nil
}
