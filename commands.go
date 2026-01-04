package main

import (
	"fmt"
	"os"

	"github.com/quanchobi/pokedexcli/internal/pokeapi"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commandExit(c *config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMapForward(conf *config, _ ...string) error {
	locations, err := pokeapi.GetLocationPage(conf.next)
	if err != nil {
		return err
	}

	conf.next = locations.Next
	conf.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(conf *config, _ ...string) error {
	if conf.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := pokeapi.GetLocationPage(conf.previous)
	if err != nil {
		return err
	}

	conf.next = locations.Next
	conf.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(conf *config, names ...string) error {
	if len(names) < 1 {
		fmt.Println("explore requires a location name")
		return nil
	}
	name := names[0]
	url := "https://pokeapi.co/api/v2/location-area/" + name
	encounters, err := pokeapi.GetAreaEncounters(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range encounters {
		fmt.Printf("- %s\n", pokemon)
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
		"explore": {
			name:        "explore",
			description: "Shows all Pokemon encounters for a given area",
			callback:    commandExplore,
		},
	}
}
