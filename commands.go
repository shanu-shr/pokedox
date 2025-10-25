package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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

	supportedCommands["map"] = cliCommand{
		name:        "map",
		description: "Display the name of 20 location areas in pokemon world",
		callback:    mapHandler,
	}

	supportedCommands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Display the name of previous 20 location areas in pokemon world",
		callback:    mapbHandler,
	}
}
