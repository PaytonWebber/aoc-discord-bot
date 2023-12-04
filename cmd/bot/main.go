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

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	cfg := config.NewConfig()
	if cfg == nil {
		log.Fatal("cfg is nil")
	}

	bot := discord.NewBot(cfg.DiscordToken)
	if bot == nil {
		log.Fatal("bot is nil")
	}

	tracker := leaderboard.NewTracker(cfg)
	if tracker == nil {
		log.Fatal("tracker is nil")
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(5 * time.Minute) // Set the interval

	go func() {
		for {
			select {
			case <-ticker.C:
				stars, err := tracker.CheckForNewStars()
				if err != nil {
					log.Printf("Error checking for new stars: %v", err)
					continue
				}

				members, err := tracker.CheckForNewMembers()
				if err != nil {
					log.Printf("Error checking for new members: %v", err)
					continue
				}

				if len(members) > 0 {
					for _, name := range members {
						discord.SendChannelMessage(bot.Session, cfg.ChannelID, name+" joined the leaderboard!")
					}
				}

				if len(stars) > 0 {
					for _, name := range stars {
						discord.SendChannelMessage(bot.Session, cfg.ChannelID, name+" got a star!")
					}
				}

				if len(stars) > 0 || len(members) > 0 {
					formattedLeaderboard := leaderboard.FormatLeaderboard(tracker.CurrentLeaderboard)
					discord.SendChannelMessage(bot.Session, cfg.ChannelID, formattedLeaderboard)
				}
			}
		}
	}()

	bot.Start()

	<-signals

	shutdown(bot)
}

func shutdown(bot *discord.Bot) {
	log.Println("shutting down")
	bot.Session.Close()
	os.Exit(0)
}
