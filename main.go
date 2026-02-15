package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

var commands map[string]cliCommand 

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Math Fetcher",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Helps you understand the CLI",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Math Fetcher > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		args := input[1:]
		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback(args)
		}
	}
}
