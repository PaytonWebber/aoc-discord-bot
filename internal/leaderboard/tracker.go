package leaderboard

import (
	"github.com/Nebula5102/aoc-discord-bot-mk2/internal/aoc"
	"github.com/Nebula5102/aoc-discord-bot-mk2/internal/config"
	"log"
	"time"
)

type AOCClient interface {
	GetLeaderboard(leaderboardID string) (*aoc.Leaderboard, error)
}

// Tracker manages the leaderboard state
type Tracker struct {
	PreviousLeaderboard *aoc.Leaderboard
	CurrentLeaderboard  *aoc.Leaderboard
	Client              AOCClient
	Config              *config.Config
	LastUpdate          time.Time
}

func NewTracker(cfg *config.Config, StoredLeaderboard *aoc.Leaderboard, client AOCClient) *Tracker {
	return &Tracker{
		Client:             client,
		Config:             cfg,
		CurrentLeaderboard: StoredLeaderboard,
	}
}

func (t *Tracker) GetLeaderboard() (*aoc.Leaderboard, error) {
	leaderboard, err := t.Client.GetLeaderboard(t.Config.LeaderboardID)
	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}

func (t *Tracker) UpdateLeaderboard() error {
	leaderboard, err := t.GetLeaderboard()
	if err != nil {
		return err
	}

	t.PreviousLeaderboard = t.CurrentLeaderboard
	t.CurrentLeaderboard = leaderboard

	return nil
}

func (t *Tracker) CheckForNewStars() ([]string, error) {
	var newStars []string

	// TODO: Get the new star data from the current leaderboard
	for memberID, member := range t.CurrentLeaderboard.Members {
		previousMember, ok := t.PreviousLeaderboard.Members[memberID]
		if !ok {
			continue
		} else if member.Stars > previousMember.Stars {
			newStars = append(newStars, member.Name)
		}
	}

	return newStars, nil
}

func (t *Tracker) CheckForNewMembers() ([]string, error) {
	var newMembers []string

	for memberID, member := range t.CurrentLeaderboard.Members {
		_, ok := t.PreviousLeaderboard.Members[memberID]
		if !ok {
			log.Printf("New member: %s", member.Name)
			newMembers = append(newMembers, member.Name)
		}
	}

	return newMembers, nil
}
