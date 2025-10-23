package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/skylarhoughtongithub/gopokedex/internal"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	trim := strings.TrimSpace(lower)
	fields := strings.Fields(trim)

	return fields
}

func main() {
	cfg := &config{}

	cache := pokecache.NewCache(5 * time.Minute)

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    func(args ...string) error { return commandHelp(args...) },
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays next 20 location areas",
		callback:    func(args ...string) error { return commandMap(cfg, cache, args...) },
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays previous 20 location areas",
		callback:    func(args ...string) error { return commandMapB(cfg, cache, args...) },
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a specific location area and list its Pokemon",
		callback:    func(args ...string) error { return commandExplore(cache, args...) },
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
