package main

import "fmt"

func commandHelp(*config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, value := range getCommands() {
		fmt.Println(fmt.Sprintf("%v: %v", value.name, value.description))
	}
	fmt.Println()

	return nil
}
