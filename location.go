package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func commandMap(cfg *config, cache *pokecache.Cache) error {
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

func commandMapB(cfg *config, cache *pokecache.Cache) error {
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
