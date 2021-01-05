package main

import (
	"flag"

	"github.com/owbot/pkg/discord"
	log "github.com/sirupsen/logrus"
)

const version = "v0.1.0"

// Command line flags
var (
	botToken = flag.String("t", "", "bot token")
	prefix   = flag.String("p", "!", "bot command prefix")
)

func main() {
	flag.Parse()

	// Configure package-level logger
	log.SetFormatter(&log.TextFormatter{})

	// New server with basic config & commands
	s, err := discord.NewServer(*botToken, *prefix)
	if err != nil {
		panic(err)
	}

	// Open the connection
	err = s.Ready()
	if err != nil {
		panic(err)
	}

	// Hold connection until user exits or fatal error occurs
	s.WaitToDie()
}
