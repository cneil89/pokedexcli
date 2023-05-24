package pokeapi

import (
	"io"
)

func (c *Client) httpRequestHelper(url *string) ([]byte, error) {
	res, err := c.httpClient.Get(*url)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, nil
	}

	return body, nil
}
