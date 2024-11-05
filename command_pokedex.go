package main

import "fmt"

func commandPokedex(c *config, _ []string) error {
	fmt.Println()
	if len(c.caughtPokemon) > 0 {
		fmt.Println("Your Pokedex:")
		for k := range c.caughtPokemon {
			fmt.Printf("\t- %v\n", k)
		}
	} else {
		fmt.Println("Your Pokedex is currently empty")
		fmt.Println("Catch some pokemon to add entries")
	}
	fmt.Println()

	return nil
}
