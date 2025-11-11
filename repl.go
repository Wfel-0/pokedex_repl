package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next string
	Previous string
}

var commands map[string]cliCommand

func startRepl() {
	commands = map[string]cliCommand{
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
			description: "Prints 20location areas in the Pokemon World",
			callback:    commandMap,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for {
		fmt.Print("Pokedex >")

		scanner.Scan()
		input = scanner.Text()
		cl_input := cleanInput(input)
		if len(cl_input) == 0 {
			continue
		}
		val, ok := commands[cl_input[0]]
		if !ok {
			fmt.Printf("Unknown command \n")
		} else {
			val.callback()
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	output := []string{}

	str := ""

	for i := range text {
		if text[i] == ' ' && str != "" {
			output = append(output, str)
			str = ""
		} else if text[i] == ' ' {
			continue
		} else {
			str += string(text[i])
		}
	}

	if str != "" {
		output = append(output, str)
	}

	return output
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}

func commandMap() error {
	

	return nil
}
