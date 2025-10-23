package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	pokecache "github.com/skylarhoughtongithub/gopokedex/internal"
)

type config struct {
	nextURL *string
	prevURL *string
}

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaDetailResponse struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandMap(cfg *config, cache *pokecache.Cache, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}

	var locationAreas LocationAreasResponse

	// Check cache first
	cachedData, found := cache.Get(url)
	if found {
		fmt.Println("Using cached data")
		if err := json.Unmarshal(cachedData, &locationAreas); err != nil {
			return fmt.Errorf("error unmarshaling cached data: %v", err)
		}
	} else {
		// If not in cache, make HTTP request
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
			return err
		}

		// Cache the response
		responseBody, err := json.Marshal(locationAreas)
		if err != nil {
			return err
		}
		cache.Add(url, responseBody)
	}

	// Common processing for both cached and fresh data
	cfg.nextURL = &locationAreas.Next
	cfg.prevURL = &locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, cache *pokecache.Cache, args ...string) error {
	if cfg.prevURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	if *cfg.prevURL == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	var locationAreas LocationAreasResponse

	cachedData, found := cache.Get(*cfg.prevURL)
	if found {
		fmt.Println("Using cached data")
		if err := json.Unmarshal(cachedData, &locationAreas); err != nil {
			return fmt.Errorf("error unmarshaling cached data: %v", err)
		}
	} else {
		resp, err := http.Get(*cfg.prevURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
			return err
		}

		responseBody, err := json.Marshal(locationAreas)
		if err != nil {
			return err
		}
		cache.Add(*cfg.prevURL, responseBody)
	}

	cfg.nextURL = &locationAreas.Next
	cfg.prevURL = &locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cache *pokecache.Cache, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a location area to explore")
	}

	locationArea := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationArea)

	fmt.Printf("Exploring %s...\n", locationArea)

	cachedData, found := cache.Get(url)
	var locationDetails LocationAreaDetailResponse

	if found {
		fmt.Println("Using cached data")
		if err := json.Unmarshal(cachedData, &locationDetails); err != nil {
			return fmt.Errorf("error unmarshaling cached data: %v", err)
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&locationDetails); err != nil {
			return err
		}

		responseBody, err := json.Marshal(locationDetails)
		if err != nil {
			return err
		}
		cache.Add(url, responseBody)
	}

	if len(locationDetails.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		uniquePokemon := make(map[string]bool)
		for _, encounter := range locationDetails.PokemonEncounters {
			pokeName := strings.ToLower(encounter.Pokemon.Name)
			if !uniquePokemon[pokeName] {
				fmt.Printf(" - %s\n", pokeName)
				uniquePokemon[pokeName] = true
			}
		}
	} else {
		fmt.Println("No Pokemon found in this location area.")
	}

	return nil
}
