package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// exit command
func commandExit(args []string) error {
	fmt.Println("Closing the Math Fetcher... Goodbye!")
	os.Exit(0)
	return nil
}

//help command with args
func commandHelp(args []string) error {
	if len(args) == 0 {
		fmt.Println(`Welcome to the Math Fetcher! Here is a list of the available commands:`)
		for key := range commands {
			fmt.Println(key)
		}
		fmt.Println("For more details about each command, type help <command-name>")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments! Please ask for help for one command at a time.")
	} else {
		arg := args[0]
		command, ok := commands[arg]
		if !ok {
			fmt.Println("This command does nothing...")
			fmt.Println("because it doesn't exist!")
		} else {
			fmt.Println("Name: " + command.name)
			fmt.Println("Description: " + command.description)
		}
	}
	return nil
}