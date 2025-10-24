package pokedex

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	cache "github.com/skylarhoughtongithub/gopokedex/internal/cache"
	pokeapi "github.com/skylarhoughtongithub/gopokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args ...string) error
}

var (
	cfg           *pokeapi.Config
	cacheInstance *cache.Cache
	commands      = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
)

// RunCLI starts the Pokedex CLI interface
func RunCLI() {
	// Initialize configuration and cache
	cfg = &pokeapi.Config{}
	cacheInstance = cache.NewCache(5 * time.Minute)

	// Set up commands after initializing cfg and cacheInstance
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    func(args ...string) error { return commandHelp() },
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays next 20 location areas",
		callback:    func(args ...string) error { return pokeapi.CommandMap(cfg, cacheInstance, args...) },
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays previous 20 location areas",
		callback:    func(args ...string) error { return pokeapi.CommandMapB(cfg, cacheInstance, args...) },
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a specific location area and list its Pokemon",
		callback:    func(args ...string) error { return pokeapi.CommandExplore(cfg, cacheInstance, args...) },
	}

	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Attempt to catch a Pokemon",
		callback:    func(args ...string) error { return pokeapi.CommandCatch(cfg, cacheInstance, args...) },
	}

	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect a pokemons attributes",
		callback:    func(args ...string) error { return pokeapi.CommandInspect(args...) },
	}

	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "List all Pokedex content",
		callback:    func(args ...string) error { return pokeapi.CommandPokedex() },
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		cleanedLine := cleanInput(line)
		if len(cleanedLine) == 0 {
			continue
		}

		cmd, exists := commands[cleanedLine[0]]
		if exists {
			err := cmd.callback(cleanedLine[1:]...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("read error:", err)
	}
}

// cleanInput processes and normalizes user input
func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	trim := strings.TrimSpace(lower)
	fields := strings.Fields(trim)

	return fields
}

// helpCommand displays available commands
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, cmd := range commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

// exitCommand closes the Pokedex application
func commandExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
