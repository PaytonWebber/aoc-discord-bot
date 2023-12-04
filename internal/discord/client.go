package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	Session *discordgo.Session
}

func NewBot(token string) *Bot {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating discord session: %v", err)
	}

	return &Bot{
		Session: session,
	}
}

func (b *Bot) Start() {
	err := b.Session.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	log.Println("Bot is now running.")
}

func SendChannelMessage(s *discordgo.Session, channelID string, message string) {
	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Printf("error sending message: %v", err)
	}
}
