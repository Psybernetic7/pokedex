// go
package main

import (
	"fmt"
	"os"
)

func commandExit(args []string, cfg *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
