package main

import (
	"AocDiscordBot/internal/config"
	"AocDiscordBot/internal/discord"
	"AocDiscordBot/internal/leaderboard"

	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.NewConfig()
	if cfg == nil {
		log.Fatal("cfg is nil")
	}

	session, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalf("error creating discord session: %v", err)
	}

	tracker := leaderboard.NewTracker(cfg)
	if tracker == nil {
		log.Fatal("tracker is nil")
	}

	bot := discord.NewBotHandler(session, tracker)
	if bot == nil {
		log.Fatal("botHandler is nil")
	}

	session.AddHandler(bot.MessageRecieved)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Start the bot
	err = session.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}
	log.Println("Bot is now running. Press CTRL-C to exit.")

	ticker := time.NewTicker(15 * time.Minute) // Set the interval

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Checking for new stars")
				stars, err := tracker.CheckForNewStars()
				if err != nil {
					log.Printf("Error checking for new stars: %v", err)
					continue
				}

				log.Println("Checking for new members")
				members, err := tracker.CheckForNewMembers()
				if err != nil {
					log.Printf("Error checking for new members: %v", err)
					continue
				}

				if len(members) > 0 {
					bot.SendChannelMessage(cfg.ChannelID, "CHALLENGER APPROACHING!")
					for _, name := range members {
						bot.SendChannelMessage(cfg.ChannelID, name+" joined the leaderboard!")
					}
				}

				if len(stars) > 0 {
					for _, name := range stars {
						bot.SendChannelMessage(cfg.ChannelID, name+" got a star!")
					}
				}

				if len(stars) > 0 || len(members) > 0 {
					formattedLeaderboard := leaderboard.FormatLeaderboard(tracker.CurrentLeaderboard)
					bot.SendChannelMessage(cfg.ChannelID, formattedLeaderboard)
				}
			}
		}
	}()

	<-signals

	shutdown(session)
}

func shutdown(session *discordgo.Session) {
	log.Println("shutting down")
	session.Close()
	os.Exit(0)
}
