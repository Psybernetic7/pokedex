// go
package main

import "fmt"

func commandHelp(args []string, cfg *Config, commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
