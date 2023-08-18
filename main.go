package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	repo "github.com/guionardo/gs-bucket/backend/repository"
	"github.com/guionardo/gs-bucket/backend/server"
)

// @title			GS-Bucket API
// @version		0.3
// @description	This application will run a HTTP server to store files
// @contact.name	Guionardo Furlan
// @contact.url	https://github.com/guionardo/gs-bucket
// @contact.email	guionardo@gmail.com
func main() {
	config, err := server.GetConfig()
	if err != nil {
		panic(err)
	}

	repository, err := repo.CreateLocalRepository(config.RepositoryFolder)
	if err != nil {
		panic(err)
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.HttpPort),
		Handler: server.Service(repository)}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
		cancel()
	}()
	log.Printf("Starting GS-Bucket service @ %s", httpServer.Addr)
	// Run the server
	if err = httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
