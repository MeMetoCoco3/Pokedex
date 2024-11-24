package types

import "github.com/MeMetoCoco3/Pokedex/internal"

type CliCommand struct {
	Name            string
	Description     string
	Function        func(*Config, string) error
	AcceptsArgument bool
}

type Config struct {
	PointLocation int
	Cache         *pokecache.Cache
}

type Respose struct {
	Area []struct {
		Name string `json:"name"`
	} `json:"results"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}
