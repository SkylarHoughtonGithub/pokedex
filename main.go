package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays a map of Pokedex locations",
		callback:    commandHelp,
	}

	commands["bmap"] = cliCommand{
		name:        "map",
		description: "Displays a previous page of map of Pokedex locations",
		callback:    commandHelp,
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
