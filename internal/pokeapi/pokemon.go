package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListPokemon -
func (c *Client) ListPokemon(locationArea string) ([]string, error) {
	url := baseURL + "/location-area/" + locationArea

	var data []byte
	if cacheEntry, exists := c.pokeCache.Get(url); exists {
		data = cacheEntry
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode == 404 {
			return nil, fmt.Errorf("Could not find area '%v'", locationArea)
		}
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("Error occurred while fetching pokemon")
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
	}

	locationsResp := LocationFullResponse{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		fmt.Println("unmarshal fail")
		return nil, err
	}

	c.pokeCache.Add(url, data)

	pokemon := make([]string, len(locationsResp.PokemonEncounters))
	for i, p := range locationsResp.PokemonEncounters {
		pokemon[i] = p.Pokemon.Name
	}

	return pokemon, nil
}
