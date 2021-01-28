package discord

import (
	"fmt"
	"strings"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const (
	// Create commands create a new resource
	Create = "create"

	// Get commands return the most important information about the
	// specified resources
	Get = "get"

	// Describe commands return a detailed description of the selected
	// resources
	Describe = "describe"

	// Update commands modify properties of the specified resources
	Update = "update"

	// Delete commands remove the selected resources from the data store
	Delete = "delete"
)

// NewRouter returns a new router with default routes attached
func NewRouter() *exrouter.Route {
	root := exrouter.New()

	// Add default and ping commands
	root.On("ping", PingCmdRun).Desc("responds with pong")

	// Add user commands
	usrCmd := root.On("user", nil)
	usrCmd.On(Create, CreateUserCmd).Desc("register new discord user with given BattleTag™")

	// Assign default/help command last within scope of root command
	root.Default = root.On("help", func(ctx *exrouter.Context) {
		_, err := ctx.Reply(Help(ctx, root))
		if err != nil {
			log.WithFields(log.Fields{
				"guildId":   ctx.Msg.GuildID,
				"discordId": ctx.Msg.Author.ID,
			}).Errorf("Failed to send help menu. Error: %v\n", err)
		}
	}).Desc("prints this help menu")

	// Return router with default command paths configured
	return root
}

// Help creates a sting containing the help menu
func Help(ctx *exrouter.Context, root *exrouter.Route) string {
	var f func(depth int, r *exrouter.Route) string
	f = func(depth int, r *exrouter.Route) string {
		text := ""
		for _, v := range r.Routes {
			text += strings.Repeat("  ", depth) + v.Name + " : " + v.Description + "\n"
			text += f(depth+1, &exrouter.Route{Route: v})
		}
		return text
	}
	return "```" + f(0, root) + "```"
}

// PingCmdRun is executed when the ping command is called
func PingCmdRun(ctx *exrouter.Context) {
	_, err := ctx.Reply(fmt.Sprintf("%s pong", ctx.Msg.Author.Mention()))
	if err != nil {
		log.WithField("command", "ping").Errorf("Failed to send response. Error %v\n", err)
	}
}

// RegisterRouter registers router to discord session
func RegisterRouter(r *exrouter.Route, s *discordgo.Session, p string) {
	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		err := r.FindAndExecute(s, p, s.State.User.ID, m.Message)
		if err != nil {
			log.Fatalf("Failed to regsiter router. Error: %v\n", err)
		}
	})
}

// CreateUserCmd creates a new user
func CreateUserCmd(ctx *exrouter.Context) {
	id := ctx.Msg.Author.ID
	un := ctx.Msg.Author.Username
	bt := ctx.Args.Get(1)

	meta := log.Fields{
		"discordID":       id,
		"discordUsername": un,
		"battletag":       bt,
	}

	log.WithFields(meta).Info("registering new user")
	_, err := ctx.Reply(
		fmt.Sprintf("registered new Discord user '%s' with BattleTag™ '%s'", ctx.Msg.Author.Mention(), bt))
	if err != nil {
		log.WithFields(meta).Errorf("Failed to register new Discord user and BattleTag™. Error: %v\n", err)
	}
}
