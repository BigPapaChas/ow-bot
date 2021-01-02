package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
)

// BotConfig defines the configuration & properties of the server
type BotConfig struct {
	Version string
	Port    int
}

var config *BotConfig
var startTime = time.Now()

// StartServer starts the http server
func StartServer(b *BotConfig) {
	config = b

	fmt.Printf("Version: %s\n", config.Version)
	r := http.NewServeMux()
	r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(defaultResponse)))

	fmt.Printf("Starting server on port %d\n", config.Port)
	portString := ":" + strconv.Itoa(config.Port)
	http.ListenAndServe(portString, handlers.CompressHandler(r))
}

func defaultResponse(w http.ResponseWriter, req *http.Request) {
	uptime := time.Now().Sub(startTime)
	fmt.Fprintf(w, "Overwatch bot version: %s\n", config.Version)
	fmt.Fprintf(w, "Server port: %d\n", config.Port)
	fmt.Fprintf(w, "Uptime: %s", uptime.Truncate(time.Second))
}
