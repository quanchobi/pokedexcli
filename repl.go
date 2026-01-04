package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func repl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	conf := config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}

	// repl loop
	for {
		fmt.Print("Pokedex > ")

		eof := !scanner.Scan()
		if eof {
			os.Exit(0)
		}

		input := scanner.Text()
		tokens := cleanInput(input)

		if len(tokens) == 0 {
			continue
		}

		command, ok := commands[tokens[0]]
		var params []string
		if len(tokens) > 1 {
			params = tokens[1:]
		}
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(&conf, params...); err != nil {
			fmt.Print(err)
		}
	}
}
