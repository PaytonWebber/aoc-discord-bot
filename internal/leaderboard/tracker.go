package leaderboard

import (
	"AocDiscordBot/internal/aoc"
	"AocDiscordBot/internal/config"
)

type Tracker struct {
	PreviousLeaderboard *aoc.Leaderboard
	CurrentLeaderboard  *aoc.Leaderboard
	Client              *aoc.Client
	Config              *config.Config
}

func NewTracker(cfg *config.Config) *Tracker {
	leaderboard, err := aoc.NewClient(cfg.SessionCookie).GetLeaderboard(cfg.LeaderboardID)
	if err != nil {
		panic(err)
	}
	return &Tracker{
		Client:             aoc.NewClient(cfg.SessionCookie),
		Config:             cfg,
		CurrentLeaderboard: leaderboard,
	}
}

func (t *Tracker) GetLeaderboard() (*aoc.Leaderboard, error) {
	leaderboard, err := t.Client.GetLeaderboard(t.Config.LeaderboardID)
	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}

func (t *Tracker) CheckForNewStars() ([]string, error) {

	leaderboard, err := t.GetLeaderboard()
	if err != nil {
		return nil, err
	}

	t.PreviousLeaderboard = t.CurrentLeaderboard
	t.CurrentLeaderboard = leaderboard

	var newStars []string

	// TODO: Get the new star data from the current leaderboard
	for memberID, member := range leaderboard.Members {

		previousMember, ok := t.PreviousLeaderboard.Members[memberID]
		if !ok {
			continue
		}

		if member.Stars > previousMember.Stars {
			newStars = append(newStars, member.Name)
		}
	}

	return newStars, nil
}

func (t *Tracker) CheckForNewMembers() ([]string, error) {

	leaderboard, err := t.GetLeaderboard()
	if err != nil {
		return nil, err
	}

	t.PreviousLeaderboard = t.CurrentLeaderboard
	t.CurrentLeaderboard = leaderboard

	var newMembers []string

	for memberID, member := range leaderboard.Members {

		_, ok := t.PreviousLeaderboard.Members[memberID]
		if !ok {
			newMembers = append(newMembers, member.Name)
		}
	}

	return newMembers, nil
}
