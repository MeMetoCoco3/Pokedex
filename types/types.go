package types

import (
	"fmt"
	"github.com/MeMetoCoco3/Pokedex/internal"
)

type CliCommand struct {
	Name            string
	Description     string
	Function        func(*Config, string) error
	AcceptsArgument bool
}

type Config struct {
	PointLocation int
	Cache         *pokecache.Cache
	Pokedex       map[string]Pokemon
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
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`

	Types []struct {
		PokemonType struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}

func PrintStats(pokemon Pokemon) {
	fmt.Printf("Name: %s \n\tHeight: %d Weight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, pType := range pokemon.Types {
		fmt.Printf("\t- %s", pType.PokemonType.Name)
	}
	fmt.Println("\t\t--Stats--")
	fmt.Printf("\t%s: %d\n", "Base Experience", pokemon.BaseExperience)
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

}
