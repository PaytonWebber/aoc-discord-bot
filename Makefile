# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bot

all: build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/bot/main.go

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/bot/main.go
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/bwmarrin/discordgo
	$(GOGET) github.com/joho/godotenv

.PHONY: all build test clean run deps build-linux

