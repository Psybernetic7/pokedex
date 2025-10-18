package main

import (
	"fmt"

	"github.com/Psybernetic7/pokedex/internal/pokeapi"
)

func commandMap(args []string, cfg *Config) error {
	var (
		resp pokeapi.LocationAreasResp
		err  error
	)

	if cfg.NextLocationURL == "" {
		resp, err = cfg.Client.ListLocationAreas()
	} else {
		resp, err = cfg.Client.ListLocationAreasByURL(cfg.NextLocationURL)
	}
	if err != nil {
		return err
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	if resp.Next != nil {
		cfg.NextLocationURL = *resp.Next
	} else {
		cfg.NextLocationURL = ""
	}
	if resp.Previous != nil {
		cfg.PrevLocationURL = *resp.Previous
	} else {
		cfg.PrevLocationURL = ""
	}

	return nil
}
