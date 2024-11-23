package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	function    func() error
}

func callExit() error {
	fmt.Println("\nSee you Pokemaniaco!")
	os.Exit(0)
	return nil
}

func callHelp() error {
	commands := getCommands()
	fmt.Printf("Usage\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func callMap() error {
	return nil
}

func callMapb() error {
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
			description: "Shows 20 locations from the pokemon world.\nEach subsequent call shows the next 20.",
			function:    callMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows 20 previous locations from the pokemon world\nEach subsequent call shows the next 20.",
			function:    callMapb,
		},
	}
	return commands
}

func main() {
	fmt.Println("Pokego!")
	commands := getCommands()
	for {
		fmt.Printf("pokedex> ")
		var input string
		fmt.Scanln(&input)
		if command, ok := commands[input]; !ok {
			fmt.Printf("Command '%s' not found\n", input)
		} else {
			command.function()
		}
	}
}
