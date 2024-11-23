package main

import (
	"encoding/json"
	"fmt"
	"github.com/MeMetoCoco3/Pokedex/callAPI"
	"github.com/MeMetoCoco3/Pokedex/internal"
	"github.com/MeMetoCoco3/Pokedex/types"
	"os"
)

func callExit(c *types.Config) error {
	fmt.Println("\nSee you Pokemaniaco!")
	os.Exit(0)
	return nil
}

func callHelp(c *types.Config) error {
	commands := getCommands()
	fmt.Printf("Usage:\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func callMap(c *types.Config) error {
	// Ojo con esta maravilla, offset limit parameter.
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", c.PointLocation)

	reader, ok := c.Cache.Get(fullURL)
	if !ok {
		// Truco del almendruco
		var err error
		reader, err = callapi.GetPokeInfo(fullURL)

		if err != nil {
			return err
		}

		c.Cache.Add(fullURL, reader)
	}

	var res types.Respose
	err := json.Unmarshal(reader, &res)
	if err != nil {
		return err
	}

	for i, area := range res.Area {
		fmt.Printf("%d. %s\n", i+c.PointLocation+1, area.Name)
	}
	c.PointLocation = c.PointLocation + 20
	return nil
}

func callMapb(c *types.Config) error {

	c.PointLocation = c.PointLocation - 20
	if c.PointLocation < 0 {
		c.PointLocation = 0
	}
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", c.PointLocation)

	reader, ok := c.Cache.Get(fullURL)
	if !ok {
		// Truco del almendruco
		var err error
		reader, err = callapi.GetPokeInfo(fullURL)

		if err != nil {
			return err
		}

		c.Cache.Add(fullURL, reader)
	}

	var res types.Respose
	err := json.Unmarshal(reader, &res)
	if err != nil {
		return err
	}

	for i, area := range res.Area {
		fmt.Printf("%d. %s\n", i+c.PointLocation+1, area.Name)
	}
	return nil

}

func getCommands() map[string]types.CliCommand {
	commands := map[string]types.CliCommand{
		"help": {
			Name:        "help",
			Description: "Gives instructions on Pokedex usage",
			Function:    callHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Closes Pokedex",
			Function:    callExit,
		},
		"map": {
			Name:        "map",
			Description: "Shows 20 locations from the pokemon world.\n\tEach subsequent call shows the next 20.",
			Function:    callMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows 20 previous locations from the pokemon world\n\tEach subsequent call shows the next 20.",
			Function:    callMapb,
		},
	}
	return commands
}

func getDefaultConfig() *types.Config {
	c := &types.Config{
		PointLocation: 0,
		Cache:         pokecache.NewCache(5000),
	}
	return c
}

func main() {
	fmt.Println("Pokego!")

	commands := getCommands()
	config := getDefaultConfig()
	for {
		fmt.Printf("pokedex> ")
		var input string
		fmt.Scanln(&input)
		if command, ok := commands[input]; !ok {
			fmt.Printf("Command '%s' not found\n", input)
		} else {
			command.Function(config)
		}
		fmt.Println(config.PointLocation)
	}
}
