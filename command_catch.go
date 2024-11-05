package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon provided")
	}

	if _, exists := c.caughtPokemon[args[0]]; exists {
		fmt.Println()
		fmt.Printf("You already caught %v\n", args[0])
		fmt.Println()

		return nil
	}

	species, err := c.pokeapiClient.GetPokemonSpecies(args[0])
	if err != nil {
		return err
	}
	pokemon, err := c.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	caught := rand.Intn(255) <= species.CaptureRate

	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	time.Sleep(1 * time.Second)
	fmt.Println(".")
	time.Sleep(1 * time.Second)
	fmt.Println(".")
	time.Sleep(1 * time.Second)
	fmt.Println(".")
	time.Sleep(1 * time.Second)
	if caught {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		c.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	fmt.Println()

	return nil
}
