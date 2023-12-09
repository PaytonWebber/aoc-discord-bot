package config

import (
	"os"
)

type Config struct {
	LeaderboardID string
	SessionCookie string
	DiscordToken  string
	ChannelID     string
}

func NewConfig() *Config {
	return &Config{
		LeaderboardID: os.Getenv("LEADERBOARD_ID"),
		SessionCookie: os.Getenv("SESSION_COOKIE"),
		DiscordToken:  os.Getenv("DISCORD_TOKEN"),
		ChannelID:     os.Getenv("CHANNEL_ID"),
	}
}

func NewTestConfig() *Config {
	return &Config{
		LeaderboardID: os.Getenv("LEADERBOARD_ID"),
		SessionCookie: os.Getenv("SESSION_COOKIE"),
		DiscordToken:  os.Getenv("TEST_DISCORD_TOKEN"),
		ChannelID:     os.Getenv("TEST_CHANNEL_ID"),
	}
}
