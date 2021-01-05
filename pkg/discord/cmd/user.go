package cmd

import (
	"fmt"

	"github.com/Necroforger/dgrouter/exrouter"
	log "github.com/sirupsen/logrus"
)

var userCmdRoot *exrouter.Route

func init() {
	userCmdRoot = exrouter.New().On("user", nil)
	userCmdRoot.On(Create, CreateUserCmd).Desc("register new discord user with given BattleTag™")
}

var userCmd = exrouter.New().On("user", nil).Desc("must call child command")

// CreateUserCmd creates a new user
func CreateUserCmd(ctx *exrouter.Context) {
	id := ctx.Msg.Author.ID
	un := ctx.Msg.Author.Username
	bt := ctx.Args.Get(1)
	log.WithFields(log.Fields{
		"discordID":       id,
		"discordUsername": un,
		"battletag":       bt,
	}).Info("registering new user")
	ctx.Reply(fmt.Sprintf("registered new Discord user '%s' with BattleTag™ '%s'", ctx.Msg.Author.Mention(), bt))
}
