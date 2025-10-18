package main

import "strings"

type cliCommand struct {
	name        string
	description string
	callback    func(args []string, cfg *Config) error
}

func cleanInput(text string) []string {
	output := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	return output
}
