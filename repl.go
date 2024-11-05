package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		if command, exists := commands[commandName]; exists {
			err := command.callback(c, args)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println()
			fmt.Printf("Unknown command '%v' - type 'help' for help", commandName)
			fmt.Println()
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
