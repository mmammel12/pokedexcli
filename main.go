// Package main
//
// Pokedex CLI for boot.dev
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range getCommands() {
		fmt.Println(fmt.Sprintf("%v: %v", value.name, value.description))
	}

	return nil
}

func commandExit() error {
	os.Exit(0)

	return errors.New("Failed to exit")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		input := scanner.Text()

		if command, ok := commands[input]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("")
	}
}
