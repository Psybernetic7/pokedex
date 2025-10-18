package pokeapi

import (
	"net/http"
	"time"

	"github.com/Psybernetic7/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(5 * time.Minute),
	}
}
