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

	// repl loop
	for {
		fmt.Print("Pokedex > ")

		eof := !scanner.Scan()
		if eof {
			commandExit()
		}

		input := scanner.Text()
		tokens := cleanInput(input)

		if len(tokens) == 0 {
			continue
		}

		command, ok := commands[tokens[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Print(err)
		}
	}
}
