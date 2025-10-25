package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedox > ")
		if scanner.Scan() {
			words := cleanInput(scanner.Text())
			fmt.Println("Your command was:", words[0])
		}
	}
}

func cleanInput(text string) []string {
	data := strings.Fields(strings.ToLower(text))
	return data
}
