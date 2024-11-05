package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("No location provided")
	}
	pokemon, err := c.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Exploring " + args[0] + "...")
	fmt.Println("Found Pokemon:")
	for _, name := range pokemon {
		fmt.Println("\t- " + name)
	}
	fmt.Println()

	return nil
}
