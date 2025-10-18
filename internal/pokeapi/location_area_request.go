// go
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	url := baseURL + "location-area"
	return c.listLocationAreasByURL(url)
}

func (c *Client) ListLocationAreasByURL(url string) (LocationAreasResp, error) {
	return c.listLocationAreasByURL(url)
}

func (c *Client) listLocationAreasByURL(url string) (LocationAreasResp, error) {
	// Check if data is in cache first
	cachedData, ok := c.cache.Get(url)
	if ok {
		// Cache hit! Use the cached data
		var out LocationAreasResp
		if err := json.Unmarshal(cachedData, &out); err != nil {
			return LocationAreasResp{}, err
		}
		return out, nil
	}

	// Cache miss - make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// Add the response data to the cache
	c.cache.Add(url, data)

	var out LocationAreasResp
	if err := json.Unmarshal(data, &out); err != nil {
		return LocationAreasResp{}, err
	}
	return out, nil
}
