package main

import (
	"errors"
	"fmt"
)

func commandInspect(args []string, cfg *Config) error {

	if len(args) != 1 {
		return errors.New("invalid argument")
	}

	p, ok := cfg.Pokedex[args[0]]

	if !ok {
		fmt.Print("you have not caught that pokemon\n")
		return nil
	}

	// p = strings.ToLower(p.Name)

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Print("Stats:\n")

	for _, s := range p.Stats {
		fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Print("Types:\n")

	for _, t := range p.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
