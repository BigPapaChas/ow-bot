package discord

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/asaskevich/govalidator"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// ServerConfig holds args required from user
type ServerConfig struct {
	// The Discord bot token
	BotToken string `valid:"required"`
	// The bot command prefix
	Prefix string `valid:"required"`
}

func (c *ServerConfig) validate() error {
	valid, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	} else if !valid {
		return errors.New("invalid user input")
	}
	return nil
}

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
func NewServer(cfg *ServerConfig) (*Server, error) {
	err := cfg.validate()
	if err != nil {
		return nil, err
	}

	s, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		log.Error("Failed to initialize discordgo server", err)
		return nil, err
	}
	r := NewRouter()

	// create server object
	srv := &Server{Session: s, Router: r, Prefix: cfg.Prefix}
	return srv, nil
}

// Ready registers the routes to discord session then opens the session
func (s *Server) Ready() error {
	// add help command
	// s.Router.Default = s.Router.On("help", DefaultCmdRun).Desc("prints this help menu")
	// Assign default/help command last within scope of root command
	s.Router.Default = s.Router.On("help", func(ctx *exrouter.Context) {
		log.Debug("helpCmdRun")
		_, err := ctx.Reply("no")
		if err != nil {
			log.Fatal(err)
		}
	}).Desc("prints this help menu")

	// register router to discord session

	s.Session.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		_ = s.Router.FindAndExecute(s.Session, s.Prefix, s.Session.State.User.ID, m.Message)
	})

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
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
		log.Info("\nGracefully exiting...")
	}()
}
