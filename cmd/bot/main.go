package main

import (
	"AocDiscordBot/internal/aoc"
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
	// Get the config from .env file
	cfg := config.NewConfig()
	if cfg == nil {
		log.Fatal("cfg is nil")
	}

	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalf("error creating discord session: %v", err)
	}

	// Get the leaderboard from file or AoC
	var storedLeaderboard *aoc.Leaderboard
	file, err := os.Open("leaderboard.json")
	if err != nil || file == nil {
		log.Printf("error opening leaderboard file: %v", err)
		log.Printf("getting leaderboard from AoC")
		client := aoc.NewClient(cfg.SessionCookie)
		storedLeaderboard, err = client.GetLeaderboard(cfg.LeaderboardID)
	} else {
		storedLeaderboard, err = leaderboard.GetLeaderboardFromFile(file)
		if err != nil {
			log.Printf("error getting leaderboard from file: %v", err)
		} else {
			log.Printf("got leaderboard from file")
		}
	}

	// Create the tracker
	tracker := leaderboard.NewTracker(cfg, storedLeaderboard)
	if tracker == nil {
		log.Fatal("tracker is nil")
	}

	// Create the bot handler
	bot := discord.NewBotHandler(session, tracker, cfg)
	if bot == nil {
		log.Fatal("botHandler is nil")
	}

	// Register the messageCreate func as a callback for MessageCreate events
	session.AddHandler(bot.MessageRecieved)

	// Setup exit signal handling to gracefully shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Start the bot
	err = session.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}

	// Check for updates on startup to see if any updates happened while the bot was offline
	err = bot.CheckForUpdates()
	if err != nil {
		log.Printf("error checking for updates: %v", err)
	}

	// Check for updates every 15 minutes
	ticker := time.NewTicker(15 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := bot.CheckForUpdates()
				if err != nil {
					log.Printf("error checking for updates: %v", err)
				}
			}
		}
	}()

	// Wait for an interrupt signal to shutdown
	<-signals

	finalLeaderboard, err := tracker.GetLeaderboard()
	if err != nil {
		log.Printf("Error getting final leaderboard: %v", err)
	}

	shutdown(session, finalLeaderboard)
}

func shutdown(session *discordgo.Session, finalLeaderboard *aoc.Leaderboard) {
	log.Printf("Shutting down...")
	session.Close()
	log.Printf("Session closed")
	err := leaderboard.StoreLeaderboard(finalLeaderboard)
	if err != nil {
		log.Printf("Error storing final leaderboard: %v", err)
	}
	log.Printf("Leaderboard stored")
	os.Exit(0)
}
