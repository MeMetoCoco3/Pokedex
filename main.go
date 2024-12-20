package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	"github.com/MeMetoCoco3/Pokedex/callAPI"
	"github.com/MeMetoCoco3/Pokedex/internal"
	"github.com/MeMetoCoco3/Pokedex/types"
)

func callExit(c *types.Config, argument string) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func callHelp(c *types.Config, argument string) error {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println()
	return nil
}

func callMap(c *types.Config, argument string) error {
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

func callMapb(c *types.Config, argument string) error {

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

func callExplore(c *types.Config, argument string) error {
	var err error
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", argument)
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

func callCatch(c *types.Config, argument string) error {
	var err error
	fullURL := "https://pokeapi.co/api/v2/pokemon/" + argument
	fmt.Println(fullURL)

	reader, ok := c.Cache.Get(fullURL)
	if !ok {
		fmt.Println("NOT CACHE")
		reader, err = callapi.GetPokeInfo(fullURL)
		if err != nil {
			return err
		}
		c.Cache.Add(fullURL, reader)
	} else {
		fmt.Println("Cache")
	}

	var pokemon types.Pokemon
	err = json.Unmarshal(reader, &pokemon)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)

	time.Sleep(1 * time.Second)
	for i := 1; i < 4; i++ {
		fmt.Println(strings.Repeat("*", i))
		time.Sleep(1 * time.Second)
	}

	if capturePokemon(pokemon.BaseExperience) {
		c.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("You got a %s!!\n", pokemon.Name)
		types.PrintStats(pokemon)

	} else {
		fmt.Printf("%s escaped...\n", pokemon.Name)
	}

	return nil
}

func callInspection(c *types.Config, pokemon string) error {
	if value, ok := c.Pokedex[pokemon]; ok {
		types.PrintStats(value)
		return nil
	} else {
		return fmt.Errorf("%s is not catched...YET!!\n", pokemon)
	}
}

func callPokedex(c *types.Config, argument string) error {
	if len(c.Pokedex) == 0 {
		return fmt.Errorf("Pokedex is empty.\n")
	}

	for _, pokemon := range c.Pokedex {
		fmt.Printf("\t- %s. \n", pokemon.Name)
	}
	return nil
}

func capturePokemon(chances int) bool {
	n := rand.IntN(400)
	fmt.Printf("You got: %d\n", n)
	if n > chances {
		return true
	}
	return false
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func getCommands() map[string]types.CliCommand {
	commands := map[string]types.CliCommand{
		"help": {
			Name:            "help",
			Description:     "Gives instructions on Pokedex usage.",
			Function:        callHelp,
			AcceptsArgument: false,
		},
		"exit": {
			Name:            "exit",
			Description:     "Closes Pokedex.",
			Function:        callExit,
			AcceptsArgument: false,
		},
		"map": {
			Name:            "map",
			Description:     "Shows 20 locations from the pokemon world.\n\tEach subsequent call shows the next 20.",
			Function:        callMap,
			AcceptsArgument: false,
		},
		"mapb": {
			Name:            "mapb",
			Description:     "Shows 20 previous locations from the pokemon world.\n\tEach subsequent call shows the next 20.",
			Function:        callMapb,
			AcceptsArgument: false,
		},
		"explore": {
			Name:            "explore",
			Description:     "Gives back a list of pokemons in a location.",
			Function:        callExplore,
			AcceptsArgument: true,
		},
		"catch": {
			Name:            "Catch",
			Description:     "Gives user a chance to catch a pokemon.",
			Function:        callCatch,
			AcceptsArgument: true,
		},
		"inspect": {
			Name:            "Inspect",
			Description:     "Prints stats from a pokemon you already catched.",
			Function:        callInspection,
			AcceptsArgument: true,
		},
		"pokedex": {
			Name:            "Pokedex",
			Description:     "Prints all catched pokemons.",
			Function:        callPokedex,
			AcceptsArgument: false,
		},
	}
	return commands
}

func getDefaultConfig() *types.Config {
	c := &types.Config{
		PointLocation: 0,
		Cache:         pokecache.NewCache(5000),
		Pokedex:       make(map[string]types.Pokemon),
	}
	return c
}

func main() {
	fmt.Println("Welcome to the Pokedex!")
	reader := bufio.NewScanner(os.Stdin)

	commands := getCommands()
	config := getDefaultConfig()
	for {
		fmt.Printf("pokedex> ")
		reader.Scan()

		var inputCommand, inputArgument string
		words := cleanInput(reader.Text())

		if len(words) == 2 {
			inputCommand = words[0]
			inputArgument = words[1]
		} else if len(words) == 1 {
			inputCommand = words[0]
		} else {
			fmt.Printf("The command '%s' does not exist\n", words)
			continue
		}

		if commandStruct, ok := commands[inputCommand]; !ok {
			fmt.Printf("Command '%s' not found\n", inputCommand)
		} else {
			if commandStruct.AcceptsArgument && len(words) == 1 {
				fmt.Printf("Command '%s' requires an argument\n", inputCommand)
			} else if !commandStruct.AcceptsArgument && len(words) > 1 {
				fmt.Printf("Command '%s' does not accept an argument\n", inputCommand)
			} else {
				err := commandStruct.Function(config, inputArgument)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
