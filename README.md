# GoPokedex

A CLI-based Pokedex application for exploring and catching Pokemon.

## Features

- Browse location areas
- Explore Pokemon encounters
- Catch and collect Pokemon
- Inspect captured Pokemon
- Manage your personal Pokedex

## Installation

```bash
git clone https://github.com/your-username/gopokedex.git
cd gopokedex
go build
```

## Usage

Run the application:
```bash
go run cmd/main.go
```

### Available Commands

- `help`: Show available commands
- `map`: View location areas (next 20)
- `mapb`: View previous location areas
- `explore [area]`: List Pokemon in a specific area
- `catch [pokemon]`: Attempt to catch a Pokemon
- `inspect [pokemon]`: View details of a caught Pokemon
- `pokedex`: List your captured Pokemon
- `exit`: Close the Pokedex

### Example Workflow

1. Use `map` to browse location areas
2. Use `explore [area]` to see Pokemon in a location
3. Use `catch [pokemon]` to try capturing Pokemon
4. Use `pokedex` to view your captured Pokemon
5. Use `inspect [pokemon]` to see details of a caught Pokemon

## Requirements

- Go 1.25.1+

Enjoy exploring the world of Pokemon!