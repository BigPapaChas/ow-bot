// genuri provides a quick way to
package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type UserInput struct {
	ClientId    int
	Permissions int
}

func (u *UserInput) validate() {
	if u.ClientId == -1 || u.Permissions == -1 {
		fmt.Println("User input required:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

var input UserInput

func init() {
	flag.IntVarP(&input.ClientId, "client-id", "c", -1, "Bot client id")
	flag.IntVarP(&input.Permissions, "permissions", "p", -1, "Permissions")
	flag.Parse()
	input.validate()
}

func main() {

	fmt.Printf("https://discord.com/oauth2/authorize?scope=bot&client_id=%d&permissions=%d\n",
		input.ClientId, input.Permissions)
}
