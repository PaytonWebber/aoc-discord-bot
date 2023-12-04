package main

import (
	"AocDiscordBot/internal/config"
	"AocDiscordBot/internal/discord"
	"AocDiscordBot/internal/leaderboard"

	"log"
	"os"
	"os/signal"
	"syscall"

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

	formattedLeaderboard := leaderboard.FormatLeaderboard(tracker.CurrentLeaderboard)
	discord.SendChannelMessage(bot.Session, cfg.ChannelID, formattedLeaderboard)

	<-signals

	shutdown(bot)
}

func shutdown(bot *discord.Bot) {
	log.Println("shutting down")
	bot.Session.Close()
	os.Exit(0)
}
