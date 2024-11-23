package main

import (
	"fmt"
	"github.com/MeMetoCoco3/Pokedex/callAPI"
	"os"
)

type cliCommand struct {
	name        string
	description string
	function    func(*config) error
}

type config struct {
	pointLocation int
}

func callExit(c *config) error {
	fmt.Println("\nSee you Pokemaniaco!")
	os.Exit(0)
	return nil
}

func callHelp(c *config) error {
	commands := getCommands()
	fmt.Printf("Usage\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func callMap(c *config) error {
	areas, err := callapi.GetArea(c.pointLocation)
	if err != nil {
		return err
	}
	for _, areaInformation := range areas {
		for _, area := range areaInformation.Results {
			fmt.Println(area.Name)
		}
	}
	c.pointLocation = c.pointLocation + 20
	return nil
}

func callMapb(c *config) error {
	c.pointLocation = c.pointLocation - 20
	areas, err := callapi.GetArea(c.pointLocation)
	if err != nil {
		return err
	}
	for _, areaInformation := range areas {
		for _, area := range areaInformation.Results {
			fmt.Println(area.Name)
		}
	}
	return nil
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Gives instructions on Pokedex usage",
			function:    callHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes Pokedex",
			function:    callExit,
		},
		"map": {
			name:        "map",
			description: "Shows 20 locations from the pokemon world.\n\tEach subsequent call shows the next 20.",
			function:    callMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows 20 previous locations from the pokemon world\n\tEach subsequent call shows the next 20.",
			function:    callMapb,
		},
	}
	return commands
}

func getDefaultConfig() *config {
	c := &config{
		pointLocation: 1,
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
			command.function(config)
		}
	}
}
