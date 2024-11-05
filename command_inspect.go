package main

import "fmt"

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon provided")
	}

	if pokemon, exists := c.caughtPokemon[args[0]]; exists {
		fmt.Println()
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf("\t- %v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("\t- %v\n", t.Type.Name)
		}
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Printf("You have not caught %v\n", args[0])
		fmt.Println()
	}

	return nil
}
