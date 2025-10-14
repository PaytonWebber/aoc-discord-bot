package aoc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	SessionCookie string
	HTTPClient    *http.Client
	Year          int
}

// NewClient creates a new AOC client with the provided session cookie and year.
func NewClient(sessionCookie string, year int) *Client {
	return &Client{
		SessionCookie: sessionCookie,
		HTTPClient:    http.DefaultClient,
		Year:          year,
	}
}

// For testing purposes, SetHTTPClient allows you to set the HTTP client used by the client.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.HTTPClient = client
}

func (c *Client) GetLeaderboard(leaderboardID string) (*Leaderboard, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/leaderboard/private/view/%s.json", c.Year, leaderboardID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "github.com/PaytonWebber/aoc-discord-bot by paytonwebber@gmail.com")
	req.Header.Set("cookie", fmt.Sprintf("session=%s", c.SessionCookie))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var leaderboard Leaderboard
	err = json.Unmarshal(body, &leaderboard)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return &leaderboard, nil
}
