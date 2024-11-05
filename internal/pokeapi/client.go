// Package pokeapi - api for interacting with pokeapi.co
package pokeapi

import (
	"net/http"
	"time"

	"github.com/mmammel12/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Pokecache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(cacheInterval),
	}
}
