package config

import (
	"os"
)

type Config struct {
	LeaderboardID string
	SessionCookie string
}

func NewConfig() *Config {
	return &Config{
		LeaderboardID: os.Getenv("LEADERBOARD_ID"),
		SessionCookie: os.Getenv("SESSION_COOKIE"),
	}
}
