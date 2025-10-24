package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"

	cache "github.com/skylarhoughtongithub/gopokedex/internal/cache"
)

var pokedex = make(map[string]Pokemon)

func calculateCatchProbability(baseExperience int) float64 {
	baseProbability := 0.5
	experienceFactor := float64(baseExperience) / 100.0
	catchProbability := baseProbability / (1 + experienceFactor)

	return catchProbability
}

func CommandCatch(cfg *Config, cache *cache.Cache, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a pokemon to catch")
	}

	pokemon := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	cachedData, found := cache.Get(url)
	var pokemonResponse PokemonResponse

	if found {
		if err := json.Unmarshal(cachedData, &pokemonResponse); err != nil {
			return fmt.Errorf("error unmarshaling cached data: %v", err)
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&pokemonResponse); err != nil {
			return err
		}

		responseBody, err := json.Marshal(pokemonResponse)
		if err != nil {
			return err
		}
		cache.Add(url, responseBody)
	}

	catchProbability := calculateCatchProbability(pokemonResponse.BaseExperience)

	if rand.Float64() < catchProbability {
		fmt.Printf("%s was caught!\n", pokemon)

		stats := []PokemonStat{}
		for _, s := range pokemonResponse.Stats {
			stats = append(stats, PokemonStat{
				Name:  s.Stat.Name,
				Value: s.BaseStat,
			})
		}

		types := []string{}
		for _, t := range pokemonResponse.Types {
			types = append(types, t.Type.Name)
		}

		pokedex[pokemon] = Pokemon{
			Name:           pokemonResponse.Name,
			Height:         pokemonResponse.Height,
			Weight:         pokemonResponse.Weight,
			BaseExperience: pokemonResponse.BaseExperience,
			Stats:          stats,
			Types:          types,
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}

func CommandInspect(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a pokemon to inspect")
	}

	pokemon := args[0]

	p, exists := pokedex[pokemon]
	if !exists {
		fmt.Printf("You have not caught %s yet\n", pokemon)
		return nil
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Base Experience: %d\n", p.BaseExperience)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  - %s: %d\n", stat.Name, stat.Value)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t)
	}

	return nil
}

func CommandPokedex() error {
	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
	} else {

		fmt.Println("Your Pokedex:")
		for _, p := range pokedex {
			fmt.Printf("  - %s\n", p.Name)
		}
	}

	return nil
}
