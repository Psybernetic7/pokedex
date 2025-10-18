package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "pokemon/" + pokemonName + "/"

	cachedData, ok := c.cache.Get(url)
	if ok {
		// Cache hit! Use the cached data
		pokemon := Pokemon{}
		if err := json.Unmarshal(cachedData, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// Cache miss - make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Add the response data to the cache
	c.cache.Add(url, data)

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
