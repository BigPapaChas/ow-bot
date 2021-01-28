package main

import (
	flag "github.com/spf13/pflag"

	"github.com/owbot/pkg/discord"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const version = "v0.1.0"

// Command line flags
var (
	botToken = flag.StringP("token", "t", "", "bot token")
	prefix   = flag.StringP("prefix", "p", "!", "bot command prefix")
)

type BotInfoHook struct {
	version string
}

func (h *BotInfoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *BotInfoHook) Fire(e *logrus.Entry) error {
	e.Data["version"] = version
	return nil
}

func main() {
	flag.Parse()

	// Configure package-level logger
	log.SetFormatter(&log.TextFormatter{})
	log.AddHook(&BotInfoHook{version})

	// New server with basic config & commands
	s, err := discord.NewServer(&discord.ServerConfig{BotToken: *botToken, Prefix: *prefix})
	if err != nil {
		panic(err)
	}

	// Open the connection
	err = s.Ready()
	if err != nil {
		log.Fatalf("failed to open session: %v\n", err)
	}

	// Hold connection until user exits or fatal error occurs
	s.WaitToDie()
}
