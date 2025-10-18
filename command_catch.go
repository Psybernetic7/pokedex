package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Psybernetic7/pokedex/internal/pokeapi"
)

func commandCatch(args []string, cfg *Config) error {

	_ = pokeapi.Pokemon{}

	if len(args) != 1 {
		return errors.New("invalid pokemon argument")
	}

	res, err := cfg.Client.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)

	roll := rand.Intn(100) // 0..99
	chance := 100 - res.BaseExperience/2
	if chance < 10 {
		chance = 10
	}
	if chance > 90 {
		chance = 90
	}

	caught := roll < chance

	if !caught {
		fmt.Printf("%s escaped!\n", res.Name)
	} else {
		fmt.Printf("%s was caught!\n", res.Name)
		name := strings.ToLower(res.Name)
		cfg.Pokedex[name] = res
	}
	return nil

}
