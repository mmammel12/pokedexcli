// Package main
//
// Pokedex CLI for boot.dev
package main

import (
	"time"

	"github.com/mmammel12/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 30*time.Second)
	c := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(c)
}
