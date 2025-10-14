package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	LeaderboardID string
	SessionCookie string
	DiscordToken  string
	ChannelID     string
	AOCYear       int
}

func NewConfig() *Config {
	// Default to current year if AOC_YEAR is not set
	year := time.Now().Year()
	if yearStr := os.Getenv("AOC_YEAR"); yearStr != "" {
		if parsedYear, err := strconv.Atoi(yearStr); err == nil && parsedYear > 2015 {
			year = parsedYear
		}
	}

	return &Config{
		LeaderboardID: os.Getenv("LEADERBOARD_ID"),
		SessionCookie: os.Getenv("SESSION_COOKIE"),
		DiscordToken:  os.Getenv("DISCORD_TOKEN"),
		ChannelID:     os.Getenv("CHANNEL_ID"),
		AOCYear:       year,
	}
}

// Validate checks that all required configuration values are present.
func (c *Config) Validate() error {
	if c.LeaderboardID == "" {
		return fmt.Errorf("LEADERBOARD_ID environment variable is required")
	}
	if c.SessionCookie == "" {
		return fmt.Errorf("SESSION_COOKIE environment variable is required")
	}
	if c.DiscordToken == "" {
		return fmt.Errorf("DISCORD_TOKEN environment variable is required")
	}
	if c.ChannelID == "" {
		return fmt.Errorf("CHANNEL_ID environment variable is required")
	}
	if c.AOCYear < 2015 {
		return fmt.Errorf("AOC_YEAR must be 2015 or later (Advent of Code started in 2015)")
	}
	return nil
}
