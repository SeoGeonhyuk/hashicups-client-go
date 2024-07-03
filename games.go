package hashicups

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetGame - Returns a specifc game
func (c *Client) GetGame(gameID int, authToken *string) (*Game, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/games/%s", c.HostURL, gameID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	game := Game{}
	err = json.Unmarshal(body, &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

// CreateGame - Create new game
func (c *Client) CreateGame(game Game, authToken *string) (*Game, error) {
	rb, err := json.Marshal(game)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/games", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gameJson := Game{}
	err = json.Unmarshal(body, &gameJson)
	if err != nil {
		return nil, err
	}

	return &gameJson, nil
}

// UpdateGame - Updates an game
func (c *Client) UpdateGame(gameID string, game Game, authToken *string) (*Game, error) {
	rb, err := json.Marshal(game)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/games/%s", c.HostURL, gameID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gameJson := Game{}
	err = json.Unmarshal(body, &gameJson)
	if err != nil {
		return nil, err
	}

	return &gameJson, nil
}

// DeleteGame - Deletes an game
func (c *Client) DeleteGame(gameID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/games/%s", c.HostURL, gameID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Deleted order" {
		return errors.New(string(body))
	}

	return nil
}
