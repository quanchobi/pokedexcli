package main

import (
	"fmt"
	"github.com/quanchobi/pokedexcli/internal/pokeapi"
	"os"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMapForward(conf *config) error {
	locations, err := pokeapi.GetLocationPage(conf.next)
	if err != nil {
		return err
	}
	// we area expecting a single page of data, so unmarshal is fine.

	conf.next = locations.Next
	conf.previous = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(conf *config) error {
	if conf.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := pokeapi.GetLocationPage(conf.previous)
	if err != nil {
		return err
	}

	conf.next = locations.Next
	conf.previous = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 next location areas",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the 20 previous location areas",
			callback:    commandMapBack,
		},
	}
}
