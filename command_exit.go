// internal commands and functions
package main

import (
	"errors"
	"os"
)

func commandExit(*config) error {
	os.Exit(0)

	return errors.New("Failed to exit")
}
