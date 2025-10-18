package main

import (
	"fmt"
)

func commandPokedex(args []string, cfg *Config) error {
	fmt.Print("Your Pokedex:\n")
	for _, p := range cfg.Pokedex {
		fmt.Printf(" - %s\n", p.Name)

	}
	return nil
}
