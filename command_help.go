package main

import "fmt"

func helpHandler() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, item := range supportedCommands {
		fmt.Println(item.description)
	}
	return nil
}
