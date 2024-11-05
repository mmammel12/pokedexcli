// internal commands and functions
package main

import (
	"fmt"
	"os"
)

func commandExit(*config, []string) error {
	os.Exit(0)

	return fmt.Errorf("Failed to exit")
}
