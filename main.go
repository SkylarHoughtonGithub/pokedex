package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokecache "github.com/skylarhoughtongithub/gopokedex/internal"
)

func main() {
	cfg := &config{}

	cache := pokecache.NewCache(5 * time.Minute)

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    func() error { return commandHelp() },
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays next 20 location areas",
		callback:    func() error { return commandMap(cfg, cache) },
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays previous 20 location areas",
		callback:    func() error { return commandMapB(cfg, cache) },
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
			err := cmd.callback()
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
