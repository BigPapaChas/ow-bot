package cmd

import (
	"fmt"
	"strings"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
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

	// DefaultPrefix for a new command router
	DefaultPrefix = "!"
)

var cmdPrefix *string
var root *exrouter.Route

// NewRouter returns a new router with default routes attached
func NewRouter(p string) *exrouter.Route {
	cmdPrefix = &p
	root = exrouter.New()

	// Add default and ping commands
	root.On("ping", PingCmdRun).Desc("responds with pong")

	// Add user commands
	usrCmd := root.On("user", nil)
	usrCmd.On("create", CreateUserCmd).Desc("register new discord user with given BattleTag™")

	// Add help command
	root.Default = root.On("help", HelpCmdRun).Desc("prints this help menu")

	// Return router with default command paths configured
	return root
}

// HelpCmdRun is executed when the default command is called
func HelpCmdRun(ctx *exrouter.Context) {
	var f func(depth int, r *exrouter.Route) string
	f = func(depth int, r *exrouter.Route) string {
		text := ""
		for _, v := range r.Routes {
			text += strings.Repeat("  ", depth) + v.Name + " : " + v.Description + "\n"
			text += f(depth+1, &exrouter.Route{Route: v})
		}
		return text
	}
	ctx.Reply("```" + f(0, root) + "```")

}

// PingCmdRun is executed when the ping command is called
func PingCmdRun(ctx *exrouter.Context) {
	ctx.Reply(fmt.Sprintf("%s pong", ctx.Msg.Author.Mention()))
}

// RegisterRouter registers router to discord session
func RegisterRouter(r *exrouter.Route, s *discordgo.Session, p string) {
	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		r.FindAndExecute(s, p, s.State.User.ID, m.Message)
	})
}
