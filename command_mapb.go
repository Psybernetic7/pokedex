// go
package main

import (
	"fmt"

	"github.com/Psybernetic7/pokedex/internal/pokeapi"
)

func commandMapBack(args []string, cfg *Config) error {
	if cfg.PrevLocationURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var (
		resp pokeapi.LocationAreasResp
		err  error
	)

	resp, err = cfg.Client.ListLocationAreasByURL(cfg.PrevLocationURL)
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
