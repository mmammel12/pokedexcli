package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations - fetch 20 locations from pokeapi.co
func (c *Client) ListLocations(pageURL *string) (LocationsShallowResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte
	if cacheEntry, exists := c.pokeCache.Get(url); exists {
		data = cacheEntry
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return LocationsShallowResponse{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationsShallowResponse{}, err
		}

		c.pokeCache.Add(url, data)
	}

	locationsResp := LocationsShallowResponse{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationsShallowResponse{}, err
	}

	return locationsResp, nil
}
