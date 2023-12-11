package leaderboard

import (
	"AocDiscordBot/internal/aoc"

	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func FormatLeaderboard(leaderboard *aoc.Leaderboard) *discordgo.MessageEmbed {
	if leaderboard == nil || len(leaderboard.Members) == 0 {
		return nil
	}

	// Convert map to a slice for sorting
	members := make([]aoc.Member, 0, len(leaderboard.Members))
	for _, member := range leaderboard.Members {
		members = append(members, member)
	}

	// Sort by local score
	sort.Slice(members, func(i, j int) bool {
		return members[i].LocalScore > members[j].LocalScore
	})

	var sb strings.Builder

	var prevScore int
	var rank int
	for i, member := range members {
		// Handle ties and first place
		if i == 0 || member.LocalScore < prevScore {
			rank = i + 1
			prevScore = member.LocalScore
		}
		line := fmt.Sprintf("%d. %s - %d points (%d stars)\n", rank, member.Name, member.LocalScore, member.Stars)
		sb.WriteString(line)
	}

	// Create the embed
	embed := &discordgo.MessageEmbed{
		Title:       "AoC Leaderboard:",
		Description: sb.String(),
		Color:       0x034F20,
	}

	return embed
}

func FormatStars(leaderboard *aoc.Leaderboard) *discordgo.MessageEmbed {

	// Convert map to a slice for sorting
	members := make([]aoc.Member, 0, len(leaderboard.Members))
	for _, member := range leaderboard.Members {
		members = append(members, member)
	}

	// Sort by local score
	sort.Slice(members, func(i, j int) bool {
		return members[i].LocalScore > members[j].LocalScore
	})

	var sb strings.Builder

	maxDays := 0
	longestNameLength := 0
	for _, member := range members {
		if len(member.CompletionDayLevels) > maxDays {
			maxDays = len(member.CompletionDayLevels)
		}
		if len(member.Name) > longestNameLength {
			longestNameLength = len(member.Name)
		}
	}

	sb.WriteString("Day")
	for i := 1; i <= maxDays; i++ {
		sb.WriteString(fmt.Sprintf(" %2d", i))
	}

	for _, member := range members {
		sb.WriteString("\n")
		sb.WriteString("    ")
		for i := 1; i <= maxDays; i++ {
			stars := 0
			day, exists := member.CompletionDayLevels[fmt.Sprint(i)]
			if exists {
				if day.Level2 != nil {
					stars++
				}
				if day.Level1 != nil {
					stars++
				}
			}

			switch stars {
			case 2:
				sb.WriteString(" ★ ")
			case 1:
				sb.WriteString(" ☆ ")
			default:
				sb.WriteString("   ")
			}
		}
		sb.WriteString(fmt.Sprintf(" %-*s", longestNameLength, member.Name))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "AoC Stars:",
		Description: "```" + sb.String() + "```",
		Color:       0xB22222,
	}

	return embed
}

func StoreLeaderboard(leaderboard *aoc.Leaderboard) error {
	file, err := os.Create("leaderboard.json")
	log.Println("Storing leaderboard")
	if err != nil {
		return fmt.Errorf("error creating leaderboard file: %w", err)
	}

	// change leadboard to json
	leaderboardJson, err := json.Marshal(leaderboard)
	if err != nil {
		return fmt.Errorf("error marshalling leaderboard: %w", err)
	}

	// write leaderboard to file
	_, err = file.Write(leaderboardJson)
	if err != nil {
		return fmt.Errorf("error writing leaderboard to file: %w", err)
	}

	return nil
}

func GetLeaderboardFromFile(file *os.File) (*aoc.Leaderboard, error) {
	leaderboard, err := os.ReadFile("leaderboard.json")
	log.Println("Getting leaderboard from file")
	if err != nil {
		return nil, fmt.Errorf("error reading leaderboard file: %w", err)
	}

	var lb aoc.Leaderboard

	err = json.Unmarshal(leaderboard, &lb)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling leaderboard: %w", err)
	}

	return &lb, nil
}
