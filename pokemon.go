package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"

	pokecache "github.com/skylarhoughtongithub/gopokedex/internal"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

type PokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

var pokedex = make(map[string]Pokemon)

func calculateCatchProbability(baseExperience int) float64 {
	baseProbability := 0.5
	experienceFactor := float64(baseExperience) / 100.0
	catchProbability := baseProbability / (1 + experienceFactor)

	return catchProbability
}

func commandCatch(cache *pokecache.Cache, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a pokemon to catch")
	}

	pokemon := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	cachedData, found := cache.Get(url)
	var pokemonDetails PokemonResponse

	if found {
		if err := json.Unmarshal(cachedData, &pokemonDetails); err != nil {
			return fmt.Errorf("error unmarshaling cached data: %v", err)
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&pokemonDetails); err != nil {
			return err
		}

		responseBody, err := json.Marshal(pokemonDetails)
		if err != nil {
			return err
		}
		cache.Add(url, responseBody)
	}

	catchProbability := calculateCatchProbability(pokemonDetails.BaseExperience)

	if rand.Float64() < catchProbability {
		fmt.Printf("%s was caught!\n", pokemon)
		pokedex[pokemon] = Pokemon{
			Name:           pokemonDetails.Name,
			BaseExperience: pokemonDetails.BaseExperience,
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
