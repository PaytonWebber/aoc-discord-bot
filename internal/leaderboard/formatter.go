package leaderboard

import (
	"AocDiscordBot/internal/aoc"
	"fmt"
	"sort"
	"strings"
)

func FormatLeaderboard(leaderboard *aoc.Leaderboard) string {
	if leaderboard == nil || len(leaderboard.Members) == 0 {
		return "Leaderboard is empty or not available."
	}

	// Convert map to a slice for sorting
	members := make([]aoc.Member, 0, len(leaderboard.Members))
	for _, member := range leaderboard.Members {
		members = append(members, member)
	}

	sort.Slice(members, func(i, j int) bool {
		return members[i].LocalScore > members[j].LocalScore // Sorting by LocalScore, descending
	})

	var sb strings.Builder
	sb.WriteString("Leaderboard:\n")

	var prevScore int
	var rank int
	for i, member := range members {
		if i == 0 || member.LocalScore < prevScore {
			rank = i + 1
			prevScore = member.LocalScore
		}
		line := fmt.Sprintf("%d. %s - %d points (%d stars)\n", rank, member.Name, member.LocalScore, member.Stars)
		sb.WriteString(line)
	}

	return sb.String()
}
