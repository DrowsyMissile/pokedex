package commands

import (
	"errors"
	"fmt"
)

func CommandPokedex(cfg *Config, args ...string) error {

	if len(cfg.CaughtPokemon) == 0 {
		return errors.New("You have not caught any pokemon yet")
	}
	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
