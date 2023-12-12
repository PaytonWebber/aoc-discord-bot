package main

import (
	"github.com/PaytonWebber/AocDiscordBot/internal/aoc"
	"github.com/PaytonWebber/AocDiscordBot/internal/config"
	"github.com/PaytonWebber/AocDiscordBot/internal/discord"
	"github.com/PaytonWebber/AocDiscordBot/internal/leaderboard"

	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := loadConfig()

	session := createDiscordSession(cfg)

	storedLeaderboard := getLeaderboard(cfg)

	tracker := initTracker(cfg, storedLeaderboard)

	bot := initBotHandler(session, tracker, cfg)

	session.AddHandler(bot.MessageRecieved)

	setupSignalHandling(session, bot)
}

func loadConfig() *config.Config {
	cfg := config.NewConfig()
	if cfg == nil {
		log.Fatal("cfg is nil")
	}
	return cfg
}

func createDiscordSession(cfg *config.Config) *discordgo.Session {
	session, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalf("error creating discord session: %v", err)
	}
	err = session.Open()
	if err != nil {
		log.Fatalf("error opening connection to discord: %v", err)
	}
	return session
}

func getLeaderboard(cfg *config.Config) *aoc.Leaderboard {
	file, err := os.Open("leaderboard.json")
	if err != nil || file == nil {
		log.Printf("error opening leaderboard file: %v", err)
		log.Printf("getting leaderboard from AoC")
		client := aoc.NewClient(cfg.SessionCookie)
		storedLeaderboard, err := client.GetLeaderboard(cfg.LeaderboardID)
		return handleLeaderboardError(storedLeaderboard, err)
	}
	storedLeaderboard, err := leaderboard.GetLeaderboardFromFile(file)
	return handleLeaderboardError(storedLeaderboard, err)
}

func handleLeaderboardError(leaderboard *aoc.Leaderboard, err error) *aoc.Leaderboard {
	if err != nil {
		log.Printf("error getting leaderboard: %v", err)
	}
	return leaderboard
}

func initTracker(cfg *config.Config, storedLeaderboard *aoc.Leaderboard) *leaderboard.Tracker {
	tracker := leaderboard.NewTracker(cfg, storedLeaderboard)
	if tracker == nil {
		log.Fatal("tracker is nil")
	}
	return tracker
}

func initBotHandler(session *discordgo.Session, tracker *leaderboard.Tracker, cfg *config.Config) *discord.BotHandler {
	bot := discord.NewBotHandler(session, tracker, cfg)
	if bot == nil {
		log.Fatal("botHandler is nil")
	}
	checkForUpdates(bot)
	return bot
}

func checkForUpdates(bot *discord.BotHandler) {
	hadUpdates, err := bot.CheckForUpdates()
	if err != nil {
		log.Printf("error checking for updates: %v", err)
	}
	if !hadUpdates {
		log.Printf("no updates")
	}
}

func setupSignalHandling(session *discordgo.Session, bot *discord.BotHandler) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Start the periodic update check in a goroutine
	go periodicUpdateCheck(bot)

	// Wait for an interrupt signal to shutdown
	<-signals

	// Perform final actions before shutting down
	finalShutdownActions(session, bot)
}

func periodicUpdateCheck(bot *discord.BotHandler) {
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			checkForUpdates(bot)
		}
	}
}

func finalShutdownActions(session *discordgo.Session, bot *discord.BotHandler) {
	log.Printf("Shutting down...")
	checkForUpdates(bot)
	session.Close()
	log.Printf("Session closed")
	os.Exit(0)
}
