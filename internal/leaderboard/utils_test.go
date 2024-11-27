// internal/leaderboard/utils_test.go

package leaderboard

import (
	"testing"

	"github.com/PaytonWebber/aoc-discord-bot/internal/aoc"
	"github.com/stretchr/testify/assert"
)

func TestFormatLeaderboard(t *testing.T) {
	// Setup a sample leaderboard
	leaderboardData := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:         1,
				Name:       "Alice",
				LocalScore: 300,
				Stars:      5,
			},
			"2": {
				ID:         2,
				Name:       "Bob",
				LocalScore: 250,
				Stars:      4,
			},
			"3": {
				ID:         3,
				Name:       "Charlie",
				LocalScore: 250,
				Stars:      4,
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	// Call the function
	embed := FormatLeaderboard(leaderboardData)

	// Assertions
	assert.NotNil(t, embed, "Embed should not be nil")
	assert.Equal(t, "AoC Leaderboard:", embed.Title, "Embed title should match")

	// Verify the description content
	expectedDescription := "1. Alice - 300 points (5 stars)\n" +
		"2. Bob - 250 points (4 stars)\n" +
		"2. Charlie - 250 points (4 stars)\n"

	assert.Equal(t, expectedDescription, embed.Description, "Embed description should match expected formatted leaderboard")
	assert.Equal(t, 0x034F20, embed.Color, "Embed color should match expected value")
}

func TestFormatLeaderboard_EmptyLeaderboard(t *testing.T) {
	// Setup an empty leaderboard
	leaderboardData := &aoc.Leaderboard{
		Members: make(map[string]aoc.Member),
		Event:   "2024",
		OwnerID: 12345,
	}

	// Call the function
	embed := FormatLeaderboard(leaderboardData)

	// Assertions
	assert.Nil(t, embed, "Embed should be nil for empty leaderboard")
}

func TestFormatLeaderboard_NilLeaderboard(t *testing.T) {
	// Call the function with nil
	embed := FormatLeaderboard(nil)

	// Assertions
	assert.Nil(t, embed, "Embed should be nil for nil leaderboard")
}

func TestFormatStars(t *testing.T) {
	// Setup a sample leaderboard
	leaderboardData := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:   1,
				Name: "Alice",
				CompletionDayLevels: map[string]aoc.CompletionDayLevel{
					"1": {
						Level1: &aoc.StarDetail{
							GetStarTs: 1672444800,
							StarIndex: 1,
						},
						Level2: &aoc.StarDetail{
							GetStarTs: 1672444900,
							StarIndex: 2,
						},
					},
					"2": {
						Level1: &aoc.StarDetail{
							GetStarTs: 1672445000,
							StarIndex: 3,
						},
					},
				},
				Stars: 3,
			},
			"2": {
				ID:                  2,
				Name:                "Bob",
				CompletionDayLevels: map[string]aoc.CompletionDayLevel{},
				Stars:               0,
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	// Call the function
	embed := FormatStars(leaderboardData)

	// Assertions
	assert.NotNil(t, embed, "Embed should not be nil")
	assert.Equal(t, "AoC Stars:", embed.Title, "Embed title should match")
	assert.Equal(t, 0xB22222, embed.Color, "Embed color should match expected value")

	// Verify the description content
	expectedDescription := "```Day  1  2\n" +
		"     ★  ☆  Alice\n" +
		"           Bob  ```"

	assert.Equal(t, expectedDescription, embed.Description, "Embed description should match expected formatted stars")
}

func TestFormatStars_EmptyLeaderboard(t *testing.T) {
	// Setup an empty leaderboard
	leaderboardData := &aoc.Leaderboard{
		Members: make(map[string]aoc.Member),
		Event:   "2024",
		OwnerID: 12345,
	}

	// Call the function
	embed := FormatStars(leaderboardData)

	// Assertions
	assert.NotNil(t, embed, "Embed should not be nil even for empty leaderboard")
	assert.Equal(t, "AoC Stars:", embed.Title, "Embed title should match")
	assert.Equal(t, "```Day```", embed.Description, "Embed description should match expected formatted stars for empty leaderboard")
}

func TestFormatStars_NilLeaderboard(t *testing.T) {
	// Call the function with nil
	embed := FormatStars(nil)

	// Assertions
	assert.Nil(t, embed, "Embed should be nil for nil leaderboard")
}
