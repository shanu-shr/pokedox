package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	registerCommands()
	initializeConfig()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedox > ")
		if scanner.Scan() {
			words := cleanInput(scanner.Text())

			command, ok := supportedCommands[words[0]]
			if ok {
				command.callback(&cfg)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	data := strings.Fields(strings.ToLower(text))
	return data
}

var cfg Config

func initializeConfig() {
	cfg = Config{
		Next:     "",
		previous: "",
	}
}
