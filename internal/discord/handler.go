package discord

import (
	"AocDiscordBot/internal/leaderboard"
	"github.com/bwmarrin/discordgo"

	"log"
	"strings"
)

type BotHandler struct {
	Session *discordgo.Session
	Tracker *leaderboard.Tracker
}

func NewBotHandler(session *discordgo.Session, tracker *leaderboard.Tracker) *BotHandler {
	return &BotHandler{
		Session: session,
		Tracker: tracker,
	}
}

func (bh *BotHandler) MessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message is "!leaderboard" command
	if strings.ToLower(m.Content) == "!leaderboard" {
		log.Println("Leaderboard command received")

		// Get the leaderboard
		formattedLeaderboard := leaderboard.FormatLeaderboard(bh.Tracker.CurrentLeaderboard)
		bh.SendChannelMessage(m.ChannelID, formattedLeaderboard)
	}
}

func (bh *BotHandler) SendChannelMessage(channelID, message string) {
	_, err := bh.Session.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Printf("error sending message: %v", err)
	}
}
