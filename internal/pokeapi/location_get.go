package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	url := baseURL + "location-area/" + locationAreaName + "/"

	cachedData, ok := c.cache.Get(url)
	if ok {
		// Cache hit! Use the cached data
		var out LocationArea
		if err := json.Unmarshal(cachedData, &out); err != nil {
			return LocationArea{}, err
		}
		return out, nil
	}

	// Cache miss - make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// Add the response data to the cache
	c.cache.Add(url, data)

	var out LocationArea
	if err := json.Unmarshal(data, &out); err != nil {
		return LocationArea{}, err
	}
	return out, nil
}
