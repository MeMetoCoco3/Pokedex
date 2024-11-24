package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/MeMetoCoco3/Pokedex/callAPI"
	"github.com/MeMetoCoco3/Pokedex/internal"
	"github.com/MeMetoCoco3/Pokedex/types"
	"os"
	"strings"
)

func callExit(c *types.Config, area string) error {
	fmt.Println("\nSee you Pokemaniaco!")
	os.Exit(0)
	return nil
}

func callHelp(c *types.Config, area string) error {
	commands := getCommands()
	fmt.Printf("Usage:\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func callMap(c *types.Config, area string) error {
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

func callMapb(c *types.Config, area string) error {

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

func callExplore(c *types.Config, area string) error {
	var err error
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	fmt.Println(fullURL)
	reader, ok := c.Cache.Get(fullURL)
	if !ok {
		reader, err = callapi.GetPokeInfo(fullURL)
		if err != nil {
			return err
		}
		c.Cache.Add(fullURL, reader)
	}

	var res types.Respose
	err = json.Unmarshal(reader, &res)
	if err != nil {
		return err
	}
	for _, pokeName := range res.PokemonEncounters {
		fmt.Printf("- %s\n", pokeName.Pokemon.Name)
	}
	return nil
}

func getCommands() map[string]types.CliCommand {
	commands := map[string]types.CliCommand{
		"help": {
			Name:          "help",
			Description:   "Gives instructions on Pokedex usage.",
			Function:      callHelp,
			AcceptsString: false,
		},
		"exit": {
			Name:          "exit",
			Description:   "Closes Pokedex.",
			Function:      callExit,
			AcceptsString: false,
		},
		"map": {
			Name:          "map",
			Description:   "Shows 20 locations from the pokemon world.\n\tEach subsequent call shows the next 20.",
			Function:      callMap,
			AcceptsString: false,
		},
		"mapb": {
			Name:          "mapb",
			Description:   "Shows 20 previous locations from the pokemon world.\n\tEach subsequent call shows the next 20.",
			Function:      callMapb,
			AcceptsString: false,
		},
		"explore": {
			Name:          "explore",
			Description:   "Gives back a list of pokemons in a location.",
			Function:      callExplore,
			AcceptsString: true,
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
		var inputCommand, inputArgument string

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command input: ", err)
		}
		input = strings.TrimSpace(input)
		words := strings.Fields(input)
		numberWords := len(words)

		if numberWords == 2 {
			inputCommand = words[0]
			inputArgument = words[1]
		} else if numberWords == 1 {
			inputCommand = words[0]
		} else {
			fmt.Printf("The command '%s' does not exist\n", input)
			continue
		}

		if commandStruct, ok := commands[inputCommand]; !ok {
			fmt.Printf("Command '%s' not found\n", inputCommand)
		} else {
			if commandStruct.AcceptsString && numberWords == 1 {
				fmt.Printf("Command '%s' requires an argument\n", inputCommand)
			} else if !commandStruct.AcceptsString && numberWords > 1 {
				fmt.Printf("Command '%s' does not accept an argument\n", inputCommand)
			} else {
				err := commandStruct.Function(config, inputArgument)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		fmt.Println(config.PointLocation)
	}
}
