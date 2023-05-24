package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "inspect details of caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all pokemon stored in Pokedex",
			callback:    commandPokedex,
		},
		"map": {
			name:        "map",
			description: "Next locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Previous locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("You have not caught any Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf("  - %s\n", k)
	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid input... Usage: inspect <pokemon>")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return errors.New("pokemon not yet caught")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid input... Usage: catch <pokemon>")
	}
	pokemon := &args[0]
	pokemonData, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing pokeball at %s... ", pokemonData.Name)
	if rand.Intn(pokemonData.BaseExperience) < 40 {
		fmt.Println("CAUGHT")
		cfg.pokedex[pokemonData.Name] = pokemonData
	} else {
		fmt.Println("ESCAPED")
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid input... Usage: explore <location>")
	}
	loc := &args[0]
	fmt.Printf("Exploring %s...\n", *loc)
	locData, err := cfg.pokeapiClient.ExploreLocation(loc)
	if err != nil {
		return err
	}
	// TODO: put gaurds incase area is void of pokemon
	fmt.Println("Pokemon Found:")
	for _, pokemon := range locData.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = locResp.Next
	cfg.prevLocURL = locResp.Previous

	for _, location := range locResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocURL == nil {
		return errors.New("you're on the first page")
	}

	locResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocURL)
	if err != nil {
		return err
	}
	cfg.nextLocURL = locResp.Next
	cfg.prevLocURL = locResp.Previous

	for _, location := range locResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range getCommands() {
		line := fmt.Sprintf("%s:\t%s", v.name, v.description)
		fmt.Println(line)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
