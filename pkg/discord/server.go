package discord

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"github.com/owbot/pkg/discord/cmd"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

// Server defines base server configuration.
type Server struct {
	// The discordgo session
	Session *discordgo.Session
	// The discordgo router
	Router *exrouter.Route
	// Prefix for bot commands
	Prefix string
}

// NewServer creates a server with the default config.
func NewServer(t, p string) (*Server, error) {
	s, err := discordgo.New("Bot " + t)
	if err != nil {
		log.Error("failed to initialize discordgo server", err)
		return nil, err
	}
	r := cmd.NewRouter(p)

	// create server object
	srv := &Server{Session: s, Router: r, Prefix: p}
	return srv, nil
}

// Ready registers the routes to discord session then opens the session
func (s *Server) Ready() error {
	// add help command
	// s.Router.Default = s.Router.On("help", cmd.DefaultCmdRun).Desc("prints this help menu")

	// register router
	cmd.RegisterRouter(s.Router, s.Session, s.Prefix)

	// open session
	err := s.Session.Open()
	if err != nil {
		log.Error("failed to open discordgo session", err)
	}
	return err
}

// WaitToDie will run the server until a kill signal is received
func (s *Server) WaitToDie() {
	defer func() {
		log.Info("Bot is now running.  Press CTRL-C to exit.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
		log.Info("\nGracefully exiting...")
	}()
}
