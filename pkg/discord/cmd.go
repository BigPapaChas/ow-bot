package discord

import (
	"fmt"

	"github.com/Necroforger/dgrouter/exrouter"
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
	// usrCmd := root.On("user", nil)

	root.On("user", nil).On(Create, CreateUserCmd).Desc("register new discord user with given BattleTag™")

	// Return router with default command paths configured
	return root
}

// Help creates a sting containing the help menu
// func Help(ctx *exrouter.Context, root *exrouter.Route) string {
// 	var f func(depth int, r *exrouter.Route) string
// 	f = func(depth int, r *exrouter.Route) string {
// 		text := ""
// 		for _, v := range r.Routes {
// 			text += strings.Repeat("  ", depth) + v.Name + " : " + v.Description + "\n"
// 			text += f(depth+1, &exrouter.Route{Route: v})
// 		}
// 		return text
// 	}
// 	return "```" + f(0, root) + "```"
// }

// PingCmdRun is executed when the ping command is called
func PingCmdRun(ctx *exrouter.Context) {
	log.Debug("pingCmdRun")
	// _, err := ctx.Reply(fmt.Sprintf("%s pong", ctx.Msg.Author.Mention()))
	_, err := ctx.Reply("pong")
	if err != nil {
		log.WithField("command", "ping").Errorf("Failed to send response. Error %v\n", err)
	}
}

// // RegisterRouter registers router to discord session
// func RegisterRouter(r *exrouter.Route, s *discordgo.Session, p string) {
// 	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
// 		lc := log.Fields{
// 			"route": r.Route.Name,
// 		}
// 		log.WithFields(lc).Debug("finding route...")
// 		err := r.FindAndExecute(s, p, s.State.User.ID, m.Message)
// 		if err != nil {
// 			log.WithFields(lc).Fatalf("Could not find route. Error: %v", err)
// 		}
// 	})
// }

// CreateUserCmd creates a new user
func CreateUserCmd(ctx *exrouter.Context) {
	log.Debug("CreateUserCmd")
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
