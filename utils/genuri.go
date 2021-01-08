package main

import (
	"flag"
	"fmt"
)

var cfg = struct {
	ClientID    int `json:"client_id"`
	Permissions int `json:"permissions"`
}{}

func init() {
	flag.IntVar(&cfg.ClientID, "c", -1, "Bot client id")
	flag.IntVar(&cfg.Permissions, "p", -1, "Permissions")
	flag.Parse()
}

func main() {
	fmt.Printf("https://discord.com/oauth2/authorize?scope=bot&client_id=%d&permissions=%d\n",
		cfg.ClientID, cfg.Permissions)
}
