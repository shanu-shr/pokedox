package main

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
