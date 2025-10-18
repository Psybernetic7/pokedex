package main

import (
	"errors"
	"fmt"

	"github.com/Psybernetic7/pokedex/internal/pokeapi"
)

func commandExplore(args []string, cfg *Config) error {
	_ = pokeapi.LocationArea{}

	if len(args) != 1 {
		return errors.New("no location provided")
	}
	res, err := cfg.Client.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + args[0] + "...")
	fmt.Println("Found Pokemon:")

	for _, pokemon := range res.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
