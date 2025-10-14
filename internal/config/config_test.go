package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("All Environment Variables Set", func(t *testing.T) {
		// Set environment variables
		t.Setenv("LEADERBOARD_ID", "prod-leaderboard")
		t.Setenv("SESSION_COOKIE", "prod-session-cookie")
		t.Setenv("DISCORD_TOKEN", "prod-discord-token")
		t.Setenv("CHANNEL_ID", "prod-channel-id")
		t.Setenv("AOC_YEAR", "2023")

		// Call NewConfig
		cfg := NewConfig()

		// Assertions
		assert.Equal(t, "prod-leaderboard", cfg.LeaderboardID, "LeaderboardID should match")
		assert.Equal(t, "prod-session-cookie", cfg.SessionCookie, "SessionCookie should match")
		assert.Equal(t, "prod-discord-token", cfg.DiscordToken, "DiscordToken should match")
		assert.Equal(t, "prod-channel-id", cfg.ChannelID, "ChannelID should match")
		assert.Equal(t, 2023, cfg.AOCYear, "AOCYear should match")
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
		assert.Equal(t, time.Now().Year(), cfg.AOCYear, "AOCYear should default to current year")
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
		assert.Equal(t, time.Now().Year(), cfg.AOCYear, "AOCYear should default to current year")
	})

	t.Run("Custom AOC Year Set", func(t *testing.T) {
		// Set custom year
		t.Setenv("AOC_YEAR", "2022")

		// Call NewConfig
		cfg := NewConfig()

		// Assertions
		assert.Equal(t, 2022, cfg.AOCYear, "AOCYear should be 2022")
	})

	t.Run("Invalid AOC Year Defaults to Current Year", func(t *testing.T) {
		// Set invalid year
		t.Setenv("AOC_YEAR", "not-a-number")

		// Call NewConfig
		cfg := NewConfig()

		// Assertions
		assert.Equal(t, time.Now().Year(), cfg.AOCYear, "AOCYear should default to current year when invalid")
	})

	t.Run("AOC Year Below 2015 Defaults to Current Year", func(t *testing.T) {
		// Set year before Advent of Code existed
		t.Setenv("AOC_YEAR", "2014")

		// Call NewConfig
		cfg := NewConfig()

		// Assertions
		assert.Equal(t, time.Now().Year(), cfg.AOCYear, "AOCYear should default to current year when below 2015")
	})
}

func TestConfigStruct(t *testing.T) {
	t.Run("Config Struct Fields", func(t *testing.T) {
		// Set environment variables
		t.Setenv("LEADERBOARD_ID", "config-leaderboard")
		t.Setenv("SESSION_COOKIE", "config-session-cookie")
		t.Setenv("DISCORD_TOKEN", "config-discord-token")
		t.Setenv("CHANNEL_ID", "config-channel-id")
		t.Setenv("AOC_YEAR", "2024")

		// Initialize Config
		cfg := NewConfig()

		// Directly test struct fields
		expected := &Config{
			LeaderboardID: "config-leaderboard",
			SessionCookie: "config-session-cookie",
			DiscordToken:  "config-discord-token",
			ChannelID:     "config-channel-id",
			AOCYear:       2024,
		}

		assert.Equal(t, expected, cfg, "Config struct should match expected values")
	})
}

func TestValidate(t *testing.T) {
	t.Run("Valid Config", func(t *testing.T) {
		cfg := &Config{
			LeaderboardID: "test-leaderboard",
			SessionCookie: "test-cookie",
			DiscordToken:  "test-token",
			ChannelID:     "test-channel",
			AOCYear:       2024,
		}

		err := cfg.Validate()
		assert.NoError(t, err, "Valid config should not return an error")
	})

	t.Run("Missing LeaderboardID", func(t *testing.T) {
		cfg := &Config{
			SessionCookie: "test-cookie",
			DiscordToken:  "test-token",
			ChannelID:     "test-channel",
			AOCYear:       2024,
		}

		err := cfg.Validate()
		assert.Error(t, err, "Should return error for missing LeaderboardID")
		assert.Contains(t, err.Error(), "LEADERBOARD_ID", "Error should mention LEADERBOARD_ID")
	})

	t.Run("Missing SessionCookie", func(t *testing.T) {
		cfg := &Config{
			LeaderboardID: "test-leaderboard",
			DiscordToken:  "test-token",
			ChannelID:     "test-channel",
			AOCYear:       2024,
		}

		err := cfg.Validate()
		assert.Error(t, err, "Should return error for missing SessionCookie")
		assert.Contains(t, err.Error(), "SESSION_COOKIE", "Error should mention SESSION_COOKIE")
	})

	t.Run("Missing DiscordToken", func(t *testing.T) {
		cfg := &Config{
			LeaderboardID: "test-leaderboard",
			SessionCookie: "test-cookie",
			ChannelID:     "test-channel",
			AOCYear:       2024,
		}

		err := cfg.Validate()
		assert.Error(t, err, "Should return error for missing DiscordToken")
		assert.Contains(t, err.Error(), "DISCORD_TOKEN", "Error should mention DISCORD_TOKEN")
	})

	t.Run("Missing ChannelID", func(t *testing.T) {
		cfg := &Config{
			LeaderboardID: "test-leaderboard",
			SessionCookie: "test-cookie",
			DiscordToken:  "test-token",
			AOCYear:       2024,
		}

		err := cfg.Validate()
		assert.Error(t, err, "Should return error for missing ChannelID")
		assert.Contains(t, err.Error(), "CHANNEL_ID", "Error should mention CHANNEL_ID")
	})

	t.Run("AOCYear Below 2015", func(t *testing.T) {
		cfg := &Config{
			LeaderboardID: "test-leaderboard",
			SessionCookie: "test-cookie",
			DiscordToken:  "test-token",
			ChannelID:     "test-channel",
			AOCYear:       2014,
		}

		err := cfg.Validate()
		assert.Error(t, err, "Should return error for AOCYear below 2015")
		assert.Contains(t, err.Error(), "AOC_YEAR", "Error should mention AOC_YEAR")
	})
}
