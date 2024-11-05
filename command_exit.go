// internal commands and functions
package main

import (
	"errors"
	"os"
)

func commandExit(*config, []string) error {
	os.Exit(0)

	return errors.New("Failed to exit")
}
