// internal/leaderboard/tracker_test.go

package leaderboard

import (
	"errors"
	"testing"

	"github.com/PaytonWebber/aoc-discord-bot/internal/aoc"
	"github.com/PaytonWebber/aoc-discord-bot/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAOCClient is a mock implementation of the AOCClient interface
type MockAOCClient struct {
	mock.Mock
}

func (m *MockAOCClient) GetLeaderboard(leaderboardID string) (*aoc.Leaderboard, error) {
	args := m.Called(leaderboardID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*aoc.Leaderboard), args.Error(1)
}

func TestNewTracker(t *testing.T) {
	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	initialLeaderboard := &aoc.Leaderboard{
		Members: make(map[string]aoc.Member),
		Event:   "2024",
		OwnerID: 12345,
	}

	mockClient := new(MockAOCClient)

	tracker := NewTracker(cfg, initialLeaderboard, mockClient)

	assert.Equal(t, cfg, tracker.Config, "Config should be set correctly")
	assert.Equal(t, initialLeaderboard, tracker.CurrentLeaderboard, "Initial leaderboard should be set correctly")
	assert.Equal(t, mockClient, tracker.Client, "AOCClient should be set correctly")
	assert.Nil(t, tracker.PreviousLeaderboard, "PreviousLeaderboard should be nil initially")
}

func TestGetLeaderboard_Success(t *testing.T) {
	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)

	expectedLeaderboard := &aoc.Leaderboard{
		Members: make(map[string]aoc.Member),
		Event:   "2024",
		OwnerID: 12345,
	}

	mockClient.On("GetLeaderboard", "test-leaderboard").Return(expectedLeaderboard, nil)

	tracker := NewTracker(cfg, nil, mockClient)

	leaderboard, err := tracker.GetLeaderboard()

	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedLeaderboard, leaderboard, "Expected leaderboard to match")
	mockClient.AssertExpectations(t)
}

func TestGetLeaderboard_Error(t *testing.T) {
	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)

	mockClient.On("GetLeaderboard", "test-leaderboard").Return(nil, errors.New("API error"))

	tracker := NewTracker(cfg, nil, mockClient)

	leaderboard, err := tracker.GetLeaderboard()

	assert.Error(t, err, "Expected an error")
	assert.Nil(t, leaderboard, "Expected leaderboard to be nil")
	mockClient.AssertExpectations(t)
}

func TestUpdateLeaderboard_Success(t *testing.T) {
	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	initialLeaderboard := &aoc.Leaderboard{
		Members: make(map[string]aoc.Member),
		Event:   "2024",
		OwnerID: 12345,
	}

	mockClient := new(MockAOCClient)

	newLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               3,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	mockClient.On("GetLeaderboard", "test-leaderboard").Return(newLeaderboard, nil)

	tracker := NewTracker(cfg, initialLeaderboard, mockClient)

	err := tracker.UpdateLeaderboard()

	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, initialLeaderboard, tracker.PreviousLeaderboard, "PreviousLeaderboard should be updated to initial")
	assert.Equal(t, newLeaderboard, tracker.CurrentLeaderboard, "CurrentLeaderboard should be updated to new")
	mockClient.AssertExpectations(t)
}

func TestUpdateLeaderboard_Error(t *testing.T) {
	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	initialLeaderboard := &aoc.Leaderboard{
		Members: make(map[string]aoc.Member),
		Event:   "2024",
		OwnerID: 12345,
	}

	mockClient := new(MockAOCClient)

	mockClient.On("GetLeaderboard", "test-leaderboard").Return(nil, errors.New("API error"))

	tracker := NewTracker(cfg, initialLeaderboard, mockClient)

	err := tracker.UpdateLeaderboard()

	assert.Error(t, err, "Expected an error")
	assert.Equal(t, initialLeaderboard, tracker.CurrentLeaderboard, "CurrentLeaderboard should remain unchanged")
	mockClient.AssertExpectations(t)
}

func TestCheckForNewStars(t *testing.T) {
	// Setup previous and current leaderboards
	previousLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			"2": {
				ID:                  2,
				Name:                "User2",
				LocalScore:          250,
				Stars:               3,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	currentLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               3, // Increased stars
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			"2": {
				ID:                  2,
				Name:                "User2",
				LocalScore:          250,
				Stars:               3,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			"3": {
				ID:                  3,
				Name:                "User3",
				LocalScore:          300,
				Stars:               1,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)
	tracker := NewTracker(cfg, previousLeaderboard, mockClient)
	mockClient.On("GetLeaderboard", "test-leaderboard").Return(currentLeaderboard, nil)

	// Simulate updating the leaderboard
	err := tracker.UpdateLeaderboard()
	assert.NoError(t, err, "Expected no error during UpdateLeaderboard")

	newStars, err := tracker.CheckForNewStars()

	assert.NoError(t, err, "Expected no error")
	assert.Len(t, newStars, 1, "Expected one new star")
	assert.Contains(t, newStars, "User1", "Expected User1 to have new stars")
}

func TestCheckForNewStars_NoNewStars(t *testing.T) {
	// Setup previous and current leaderboards with no new stars
	previousLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	currentLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)
	tracker := NewTracker(cfg, previousLeaderboard, mockClient)
	mockClient.On("GetLeaderboard", "test-leaderboard").Return(currentLeaderboard, nil)

	// Simulate updating the leaderboard
	err := tracker.UpdateLeaderboard()
	assert.NoError(t, err, "Expected no error during UpdateLeaderboard")

	newStars, err := tracker.CheckForNewStars()

	assert.NoError(t, err, "Expected no error")
	assert.Len(t, newStars, 0, "Expected no new stars")
}

func TestCheckForNewMembers(t *testing.T) {
	// Setup previous and current leaderboards
	previousLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	currentLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			"2": {
				ID:                  2,
				Name:                "User2",
				LocalScore:          250,
				Stars:               3,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)
	tracker := NewTracker(cfg, previousLeaderboard, mockClient)
	mockClient.On("GetLeaderboard", "test-leaderboard").Return(currentLeaderboard, nil)

	// Simulate updating the leaderboard
	err := tracker.UpdateLeaderboard()
	assert.NoError(t, err, "Expected no error during UpdateLeaderboard")

	newMembers, err := tracker.CheckForNewMembers()

	assert.NoError(t, err, "Expected no error")
	assert.Len(t, newMembers, 1, "Expected one new member")
	assert.Contains(t, newMembers, "User2", "Expected User2 to be identified as a new member")
}

func TestCheckForNewMembers_NoNewMembers(t *testing.T) {
	// Setup previous and current leaderboards with no new members
	previousLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	currentLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)
	tracker := NewTracker(cfg, previousLeaderboard, mockClient)
	mockClient.On("GetLeaderboard", "test-leaderboard").Return(currentLeaderboard, nil)

	// Simulate updating the leaderboard
	err := tracker.UpdateLeaderboard()
	assert.NoError(t, err, "Expected no error during UpdateLeaderboard")

	newMembers, err := tracker.CheckForNewMembers()

	assert.NoError(t, err, "Expected no error")
	assert.Len(t, newMembers, 0, "Expected no new members")
}

func TestCheckForNewStars_PartialPreviousLeaderboard(t *testing.T) {
	// Setup previous and current leaderboards where some members are missing
	previousLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			// User2 is missing in previous leaderboard
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	currentLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               3, // Increased stars
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			"2": {
				ID:                  2,
				Name:                "User2",
				LocalScore:          250,
				Stars:               3,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)
	tracker := NewTracker(cfg, previousLeaderboard, mockClient)
	mockClient.On("GetLeaderboard", "test-leaderboard").Return(currentLeaderboard, nil)

	// Simulate updating the leaderboard
	err := tracker.UpdateLeaderboard()
	assert.NoError(t, err, "Expected no error during UpdateLeaderboard")

	newStars, err := tracker.CheckForNewStars()

	assert.NoError(t, err, "Expected no error")
	assert.Len(t, newStars, 1, "Expected one new star")
	assert.Contains(t, newStars, "User1", "Expected User1 to have new stars")
}

func TestCheckForNewMembers_PartialPreviousLeaderboard(t *testing.T) {
	// Setup previous and current leaderboards where some members are missing
	previousLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			// User2 is missing in previous leaderboard
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	currentLeaderboard := &aoc.Leaderboard{
		Members: map[string]aoc.Member{
			"1": {
				ID:                  1,
				Name:                "User1",
				LocalScore:          200,
				Stars:               2,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
			"2": {
				ID:                  2,
				Name:                "User2",
				LocalScore:          250,
				Stars:               3,
				CompletionDayLevels: make(map[string]aoc.CompletionDayLevel),
			},
		},
		Event:   "2024",
		OwnerID: 12345,
	}

	cfg := &config.Config{
		LeaderboardID: "test-leaderboard",
		SessionCookie: "test-session-cookie",
		DiscordToken:  "test-discord-token",
		ChannelID:     "test-channel",
	}

	mockClient := new(MockAOCClient)
	tracker := NewTracker(cfg, previousLeaderboard, mockClient)
	mockClient.On("GetLeaderboard", "test-leaderboard").Return(currentLeaderboard, nil)

	// Simulate updating the leaderboard
	err := tracker.UpdateLeaderboard()
	assert.NoError(t, err, "Expected no error during UpdateLeaderboard")

	newMembers, err := tracker.CheckForNewMembers()

	assert.NoError(t, err, "Expected no error")
	assert.Len(t, newMembers, 1, "Expected one new member")
	assert.Contains(t, newMembers, "User2", "Expected User2 to be identified as a new member")
}
