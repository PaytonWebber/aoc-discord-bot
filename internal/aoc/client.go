package aoc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	SessionCookie string
}

func NewClient(sessionCookie string) *Client {
	return &Client{
		SessionCookie: sessionCookie,
	}
}

func (c *Client) GetLeaderboard(leaderboardID string) (*Leaderboard, error) {
	url := fmt.Sprintf("https://adventofcode.com/2023/leaderboard/private/view/%s.json", leaderboardID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request: %w", err)
	}

	req.Header.Set("cookie", fmt.Sprintf("session=%s", c.SessionCookie))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	var leaderboard Leaderboard
	err = json.Unmarshal(body, &leaderboard)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling body: %w", err)
	}

	return &leaderboard, nil
}
