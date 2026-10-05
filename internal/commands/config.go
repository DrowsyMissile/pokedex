package commands

import "github.com/DrowsyMissile/pokedex/internal/pokeapi"

type Config struct {
	Commands         map[string]CliCommand
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]pokeapi.Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}
