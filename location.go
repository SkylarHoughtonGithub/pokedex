package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var locationAreas LocationAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
		return err
	}

	cfg.nextURL = &locationAreas.Next
	cfg.prevURL = &locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	if *cfg.prevURL == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := http.Get(*cfg.prevURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var locationAreas LocationAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
		return err
	}

	cfg.nextURL = &locationAreas.Next
	cfg.prevURL = &locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
