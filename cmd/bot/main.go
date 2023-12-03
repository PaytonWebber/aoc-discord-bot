package main

import (
	"AocDiscordBot/internal/config"
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

	tracker := leaderboard.NewTracker(cfg)
	if tracker == nil {
		log.Fatal("tracker is nil")
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	newStars, err := tracker.CheckForNewStars()
	if err != nil {
		log.Fatalf("error checking for new stars: %v", err)
	}

	if len(newStars) > 0 {
		log.Printf("new stars: %v", newStars)
	}

	for _, member := range tracker.CurrentLeaderboard.Members {
		log.Printf("Name: %s, Stars: %d Local Score: %d", member.Name, member.Stars, member.LocalScore)
	}

	<-signals

	shutdown()
}

func shutdown() {
	log.Println("shutting down")
	os.Exit(0)
}
