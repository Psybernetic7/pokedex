// go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	cfg := newConfig()

	commands := map[string]cliCommand{}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func(args []string, cfg *Config) error {
			return commandHelp(args, cfg, commands)
		},
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "List location areas (20 per page)",
		callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "List previous 20 location areas",
		callback:    commandMapBack,
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explore each area",
		callback:    commandExplore,
	}

	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Catch a pokemon",
		callback:    commandCatch,
	}

	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Details of your caught pokemon",
		callback:    commandInspect,
	}

	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "List of all your pokemons",
		callback:    commandPokedex,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(strings.ToLower(line))
		cmdName := fields[0]
		args := []string{}
		if len(fields) > 1 {
			args = fields[1:]
		}

		cmd, ok := commands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(args, cfg); err != nil {
			fmt.Println(err)
		}
	}
}
