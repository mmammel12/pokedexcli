package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	var data []byte
	if cacheEntry, exists := c.pokeCache.Get(url); exists {
		data = cacheEntry
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}

		if res.StatusCode == 404 {
			return Pokemon{}, fmt.Errorf("Could not find pokemon '%v'", name)
		}
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("Error occurred while fetching pokemon")
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.pokeCache.Add(url, data)

	return pokemon, nil
}

// GetPokemonSpecies -
func (c *Client) GetPokemonSpecies(name string) (Species, error) {
	url := baseURL + "/pokemon-species/" + name

	var data []byte
	if cacheEntry, exists := c.pokeCache.Get(url); exists {
		data = cacheEntry
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Species{}, err
		}

		if res.StatusCode == 404 {
			return Species{}, fmt.Errorf("Could not find pokemon '%v'", name)
		}
		if res.StatusCode > 299 {
			return Species{}, fmt.Errorf("Error occurred while fetching pokemon")
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Species{}, err
		}
	}

	species := Species{}
	err := json.Unmarshal(data, &species)
	if err != nil {
		return Species{}, err
	}

	c.pokeCache.Add(url, data)

	return species, nil
}
