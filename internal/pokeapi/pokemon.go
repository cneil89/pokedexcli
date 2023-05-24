package pokeapi

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetPokemon(pokemon *string) (Pokemon, error) {
	if pokemon == nil {
		return Pokemon{}, errors.New("invalid pokemon: nil")
	}
	url := BASE_URL + "/pokemon/" + *pokemon

	if data, ok := c.cache.Get(url); ok {
		pokemonData := Pokemon{}
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonData, nil
	}

	body, err := c.httpRequestHelper(&url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)
	return pokemonData, nil
}
