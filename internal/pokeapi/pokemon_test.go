package pokeapi

import (
	"testing"
)

func TestCalculateCatchProbability(t *testing.T) {
	testCases := []struct {
		baseExperience int
		minProb        float64
		maxProb        float64
	}{
		{50, 0.25, 0.5},   // Lower base experience
		{100, 0.25, 0.5},  // Medium base experience
		{500, 0.05, 0.25}, // High base experience
		{1000, 0.01, 0.1}, // Very high base experience
	}

	for _, tc := range testCases {
		prob := calculateCatchProbability(tc.baseExperience)

		if prob < tc.minProb || prob > tc.maxProb {
			t.Errorf("Catch probability for base experience %d is %f, expected between %f and %f",
				tc.baseExperience, prob, tc.minProb, tc.maxProb)
		}
	}
}

func TestCommandInspect(t *testing.T) {
	// Reset the global pokedex for testing
	pokedex = make(map[string]Pokemon)

	// Add a test Pokemon to the Pokedex
	testPokemon := Pokemon{
		Name:           "testmon",
		Height:         10,
		Weight:         50,
		BaseExperience: 100,
		Stats: []PokemonStat{
			{Name: "hp", Value: 45},
			{Name: "attack", Value: 49},
		},
		Types: []string{"normal"},
	}
	pokedex["testmon"] = testPokemon

	// Test successful inspection
	err := CommandInspect("testmon")
	if err != nil {
		t.Errorf("Unexpected error inspecting caught Pokemon: %v", err)
	}

	// Test inspecting uncaught Pokemon
	err = CommandInspect("uncaught")
	if err != nil {
		t.Errorf("Unexpected error handling uncaught Pokemon: %v", err)
	}
}

func TestCommandPokedex(t *testing.T) {
	// Reset the global pokedex for testing
	pokedex = make(map[string]Pokemon)

	// Test empty Pokedex
	err := CommandPokedex()
	if err != nil {
		t.Errorf("Unexpected error with empty Pokedex: %v", err)
	}

	// Add some test Pokemon
	pokedex["pokemon1"] = Pokemon{Name: "pokemon1"}
	pokedex["pokemon2"] = Pokemon{Name: "pokemon2"}

	// Test Pokedex with Pokemon
	err = CommandPokedex()
	if err != nil {
		t.Errorf("Unexpected error with populated Pokedex: %v", err)
	}
}
