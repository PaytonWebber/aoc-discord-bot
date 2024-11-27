package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("All Environment Variables Set", func(t *testing.T) {
		// Set environment variables
		t.Setenv("LEADERBOARD_ID", "prod-leaderboard")
		t.Setenv("SESSION_COOKIE", "prod-session-cookie")
		t.Setenv("DISCORD_TOKEN", "prod-discord-token")
		t.Setenv("CHANNEL_ID", "prod-channel-id")

		// Call NewConfig
		cfg := NewConfig()

		// Assertions
		assert.Equal(t, "prod-leaderboard", cfg.LeaderboardID, "LeaderboardID should match")
		assert.Equal(t, "prod-session-cookie", cfg.SessionCookie, "SessionCookie should match")
		assert.Equal(t, "prod-discord-token", cfg.DiscordToken, "DiscordToken should match")
		assert.Equal(t, "prod-channel-id", cfg.ChannelID, "ChannelID should match")
	})

	t.Run("Some Environment Variables Missing", func(t *testing.T) {
		// Set some environment variables
		t.Setenv("LEADERBOARD_ID", "prod-leaderboard")
		t.Setenv("SESSION_COOKIE", "prod-session-cookie")
		// DISCORD_TOKEN and CHANNEL_ID are not set

		// Call NewConfig
		cfg := NewConfig()

		// Assertions
		assert.Equal(t, "prod-leaderboard", cfg.LeaderboardID, "LeaderboardID should match")
		assert.Equal(t, "prod-session-cookie", cfg.SessionCookie, "SessionCookie should match")
		assert.Equal(t, "", cfg.DiscordToken, "DiscordToken should be empty")
		assert.Equal(t, "", cfg.ChannelID, "ChannelID should be empty")
	})

	t.Run("No Environment Variables Set", func(t *testing.T) {
		// No environment variables are set

		// Call NewConfig
		cfg := NewConfig()

		// Assertions
		assert.Equal(t, "", cfg.LeaderboardID, "LeaderboardID should be empty")
		assert.Equal(t, "", cfg.SessionCookie, "SessionCookie should be empty")
		assert.Equal(t, "", cfg.DiscordToken, "DiscordToken should be empty")
		assert.Equal(t, "", cfg.ChannelID, "ChannelID should be empty")
	})
}

func TestConfigStruct(t *testing.T) {
	t.Run("Config Struct Fields", func(t *testing.T) {
		// Set environment variables
		t.Setenv("LEADERBOARD_ID", "config-leaderboard")
		t.Setenv("SESSION_COOKIE", "config-session-cookie")
		t.Setenv("DISCORD_TOKEN", "config-discord-token")
		t.Setenv("CHANNEL_ID", "config-channel-id")

		// Initialize Config
		cfg := NewConfig()

		// Directly test struct fields
		expected := &Config{
			LeaderboardID: "config-leaderboard",
			SessionCookie: "config-session-cookie",
			DiscordToken:  "config-discord-token",
			ChannelID:     "config-channel-id",
		}

		assert.Equal(t, expected, cfg, "Config struct should match expected values")
	})
}
