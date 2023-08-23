package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guionardo/gs-bucket/backend/server"
)

var (
	Build = "development"
)

//	@title			GS-Bucket API
//	@version		0.0.6
//	@description	This application will run a HTTP server to store files
//	@contact.name	Guionardo Furlan
//	@contact.url	https://github.com/guionardo/gs-bucket
//	@contact.email	guionardo@gmail.com
func main() {
	if buildTime, err := time.Parse("2006-01-02T15:04:05-0700", Build); err == nil {
		server.BuildTime = buildTime
	}
	log.Printf("Starting %s - built at %v", server.AppName, server.BuildTime)
	config, err := server.GetConfig()
	if err != nil {
		panic(err)
	}
	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	if err := server.Start(sig, config); err != nil {
		panic(err)
	}

}
