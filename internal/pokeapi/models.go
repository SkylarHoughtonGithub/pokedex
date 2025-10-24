package pokeapi

// Config holds the navigation state for paginated API responses
type Config struct {
	NextURL *string
	PrevURL *string
}

// LocationAreasResponse represents the API response for location areas
type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// LocationAreaDetailResponse represents detailed information about a location area
type LocationAreaDetailResponse struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// PokemonStat represents a Pokemon's individual stat
type PokemonStat struct {
	Name  string `json:"stat.name"`
	Value int    `json:"base_stat"`
}

// Pokemon represents the full details of a Pokemon
type Pokemon struct {
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	BaseExperience int           `json:"base_experience"`
	Stats          []PokemonStat `json:"stats"`
	Types          []string      `json:"types"`
}

// PokemonResponse represents a simplified Pokemon response
type PokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
