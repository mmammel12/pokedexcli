package main

import "fmt"

func commandHelp(*config, []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, value := range getCommands() {
		fmt.Printf("\t- %v: %v\n", value.name, value.description)
	}
	fmt.Println()

	return nil
}
