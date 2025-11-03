# Pokedex CLI

An interactive command-line Pokédex built in Go. The REPL-style interface lets you browse the [PokeAPI](https://pokeapi.co/), explore location areas, and build up a local Pokédex by catching and inspecting Pokémon. Responses from the API are cached in-memory to keep the experience snappy as you page back and forth through data.

## Features
- Interactive REPL with contextual help and graceful exit handling
- Browse location areas with forward/back pagination driven by the PokeAPI
- Explore a location area to see which Pokémon can spawn there
- Attempt to catch Pokémon and build a personal Pokédex
- Inspect caught Pokémon to review base stats, height, weight, and types
- Transparent HTTP response caching to reduce redundant network calls

## Requirements
- Go 1.25 or newer (see `go.mod`)
- Network access to https://pokeapi.co/

## Getting Started
Clone the repository and install dependencies (the project relies only on the Go standard library).

```bash
git clone https://github.com/Psybernetic7/pokedex.git
cd pokedex
go run .
```

You should see the interactive prompt:

```
Pokedex >
```

Type `help` at any time to list the supported commands.

## Command Reference

| Command | Usage | Description |
|---------|-------|-------------|
| `help` | `help` | Print available commands and their descriptions. |
| `map` | `map` | Fetch the next page of 20 location areas from the PokeAPI. |
| `mapb` | `mapb` | Go back to the previous page of location areas. |
| `explore` | `explore <location-area>` | Show Pokémon encounters for the specified location area. |
| `catch` | `catch <pokemon>` | Attempt to catch a Pokémon; successful catches are added to your Pokédex. |
| `inspect` | `inspect <pokemon>` | Display stats, height, weight, and types for a caught Pokémon. |
| `pokedex` | `pokedex` | List every Pokémon you have caught in this session. |
| `exit` | `exit` | Close the REPL. |

Notes:
- Location area names and Pokémon names are case-insensitive inside the REPL.
- Successful catches are stored in-memory; restarting the program clears the Pokédex.

## Development

Key packages:
- `internal/pokeapi`: Thin client responsible for talking to the public PokeAPI with automatic caching.
- `internal/pokecache`: Goroutine-backed in-memory cache with TTL-based eviction.
- `command_*.go`: Individual command implementations that manipulate shared REPL state.

## Troubleshooting
- If `map` or `explore` hangs, verify that the PokeAPI is reachable from your network.
- Cached responses expire after five minutes. Use `map`/`explore` again to refresh stale entries.

