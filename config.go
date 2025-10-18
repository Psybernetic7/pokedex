package main

import (
	"github.com/Psybernetic7/pokedex/internal/pokeapi"
)

type Config struct {
	NextLocationURL string
	PrevLocationURL string
	Client          *pokeapi.Client
	Pokedex         map[string]pokeapi.Pokemon
}

func newConfig() *Config {
	c := pokeapi.NewClient()
	return &Config{
		Client:  &c,
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
}
