package main

import (
	"time"

	"github.com/cneil89/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := config{
		pokedex:       make(map[string]pokeapi.Pokemon),
		pokeapiClient: client,
	}
	startRepl(&cfg)
}
