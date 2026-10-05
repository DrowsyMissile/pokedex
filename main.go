package main

import (
	"time"

	"github.com/DrowsyMissile/pokedex/internal/commands"
	"github.com/DrowsyMissile/pokedex/internal/pokeapi"
	"github.com/DrowsyMissile/pokedex/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &commands.Config{
		Commands:      repl.GetCommands(),
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeClient,
	}

	repl.StartRepl(cfg)
}
