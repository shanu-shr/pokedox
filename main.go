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

var supportedCommands map[string]cliCommand

func registerCommands() {
	supportedCommands = make(map[string]cliCommand)

	supportedCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedox",
		callback:    commandExit,
	}

	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    helpHandler,
	}
}

func main() {

	registerCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedox > ")
		if scanner.Scan() {
			words := cleanInput(scanner.Text())

			command, ok := supportedCommands[words[0]]
			if ok {
				command.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpHandler() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, item := range supportedCommands {
		fmt.Println(item.description)
	}
	return nil
}

func cleanInput(text string) []string {
	data := strings.Fields(strings.ToLower(text))
	return data
}
