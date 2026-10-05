package commands

import (
	"errors"
	"fmt"
)

func CommandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	if pokemon, caught := cfg.CaughtPokemon[args[0]]; caught {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			if stat.Stat.Name == "hp" ||
				stat.Stat.Name == "attack" ||
				stat.Stat.Name == "defense" ||
				stat.Stat.Name == "special-attack" ||
				stat.Stat.Name == "special-defense" ||
				stat.Stat.Name == "speed" {
				fmt.Printf(" * %s: %d\n", stat.Stat.Name, stat.BaseStat)
			}
		}
		fmt.Printf("Types:\n")
		for _, typ := range pokemon.Types {
			fmt.Printf(" * %s\n", typ.Type.Name)
		}
		return nil
	}
	return errors.New("you have not caught that pokemon")
}
