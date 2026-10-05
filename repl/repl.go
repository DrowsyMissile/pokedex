package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DrowsyMissile/pokedex/internal/commands"
)

func StartRepl(cfg *commands.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := cfg.Commands[commandName]
		if exists {
			err := command.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func GetCommands() map[string]commands.CliCommand {
	return map[string]commands.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commands.CommandHelp,
		},
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Attempt to catch a pokemon",
			Callback:    commands.CommandCatch,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "Inspect caught pokemon",
			Callback:    commands.CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays all caught pokemon",
			Callback:    commands.CommandPokedex,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commands.CommandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commands.CommandMapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.CommandExit,
		},
		"explore": {
			Name:        "explore <location>",
			Description: "Explore what pokemons are in this area",
			Callback:    commands.CommandExplore,
		},
	}
}
