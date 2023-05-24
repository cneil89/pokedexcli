package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cneil89/pokedex/internal/pokeapi"
)

type config struct {
	pokedex       map[string]pokeapi.Pokemon
	pokeapiClient pokeapi.Client
	nextLocURL    *string
	prevLocURL    *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		if command, ok := getCommands()[commandName]; ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// clean input so it can be used to compare against commands
func cleanInput(s string) []string {
	s = strings.ToLower(s)
	tmp := strings.Fields(s)
	return tmp
}
