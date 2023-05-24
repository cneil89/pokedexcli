package pokeapi

import (
	"encoding/json"
	"errors"
)

func (c *Client) ListLocations(pageURL *string) (RespLocData, error) {
	url := BASE_URL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		locResp := RespLocData{}
		err := json.Unmarshal(data, &locResp)
		if err != nil {
			return RespLocData{}, err
		}
		return locResp, nil
	}

	body, err := c.httpRequestHelper(&url)
	if err != nil {
		return RespLocData{}, err
	}

	locResp := RespLocData{}
	err = json.Unmarshal(body, &locResp)
	if err != nil {
		return RespLocData{}, err
	}

	c.cache.Add(url, body)
	return locResp, nil
}

func (c *Client) ExploreLocation(location *string) (locExploreData, error) {
	if location == nil {
		return locExploreData{}, errors.New("invalid location: nil")
	}
	url := BASE_URL + "/location-area/" + *location

	if data, ok := c.cache.Get(url); ok {
		locData := locExploreData{}
		err := json.Unmarshal(data, &locData)
		if err != nil {
			return locExploreData{}, err
		}
		return locData, nil
	}

	body, err := c.httpRequestHelper(&url)
	if err != nil {
		return locExploreData{}, err
	}

	locData := locExploreData{}
	err = json.Unmarshal(body, &locData)
	if err != nil {
		return locExploreData{}, err
	}

	c.cache.Add(url, body)
	return locData, nil
}
