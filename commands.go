package main

import "github.com/mmammel12/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	caughtPokemon map[string]pokeapi.Pokemon
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area to see the pokemon available - Example: explore pastoria-city-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon - Example: catch pikachu",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon you have already caught - Example: inspect pikachu",
			callback:    commandInspect,
		},
	}
}
