package discord

import (
	"AocDiscordBot/internal/config"
	"AocDiscordBot/internal/leaderboard"
	"github.com/bwmarrin/discordgo"

	"log"
	"strings"
)

type BotHandler struct {
	Session *discordgo.Session
	Tracker *leaderboard.Tracker
	cfg     *config.Config
}

func NewBotHandler(session *discordgo.Session, tracker *leaderboard.Tracker, cfg *config.Config) *BotHandler {
	return &BotHandler{
		Session: session,
		Tracker: tracker,
		cfg:     cfg,
	}
}

func (bh *BotHandler) CheckForUpdates() error {
	log.Println("Checking for updates...")
	newStars, err := bh.Tracker.CheckForNewStars()
	if err != nil {
		return err
	}

	newMembers, err := bh.Tracker.CheckForNewMembers()
	if err != nil {
		return err
	}

	if len(newStars) > 0 {
		log.Printf("new stars: %v", newStars)
		for _, member := range newStars {
			bh.SendChannelMessage(bh.cfg.ChannelID, member+" has got a star!")
		}
	}

	if len(newMembers) > 0 {
		log.Printf("new members: %v", newMembers)
		bh.SendChannelMessage(bh.cfg.ChannelID, "CHALLENGER APPROACHING!")
		for _, member := range newMembers {
			bh.SendChannelMessage(bh.cfg.ChannelID, member+" has joined the leaderboard!")
		}
	}

	if len(newStars) > 0 || len(newMembers) > 0 {
		formattedLeaderboard := leaderboard.FormatLeaderboard(bh.Tracker.CurrentLeaderboard)
		bh.SendChannelMessageEmbed(bh.cfg.ChannelID, formattedLeaderboard)
	}

	return nil
}

func (bh *BotHandler) MessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.ChannelID != bh.cfg.ChannelID {
		return
	}

	// Check if the message is "!leaderboard" command
	if strings.ToLower(m.Content) == "!leaderboard" {
		log.Println("Leaderboard command received")

		// Get the leaderboard
		formattedLeaderboard := leaderboard.FormatLeaderboard(bh.Tracker.CurrentLeaderboard)
		bh.SendChannelMessageEmbed(m.ChannelID, formattedLeaderboard)
	}
}

func (bh *BotHandler) SendChannelMessage(channelID, message string) {
	_, err := bh.Session.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Printf("error sending message: %v", err)
	}
}

func (bh *BotHandler) SendChannelMessageEmbed(channelID string, embed *discordgo.MessageEmbed) {
	_, err := bh.Session.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Printf("error sending message: %v", err)
	}
}
