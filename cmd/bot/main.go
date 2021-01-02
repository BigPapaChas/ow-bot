package main

import (
	"fmt"

	"github.com/owbot/server"
)

const version = "v0.1.0"

func main() {
	fmt.Println("Starting Overwatch Bot...")
	b := &server.BotConfig{
		Version: version,
		Port:    8080,
	}
	server.StartServer(b)
}
