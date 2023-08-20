package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	repo "github.com/guionardo/gs-bucket/backend/repository"
)

var ServerIsRunning bool

func Start(sig chan os.Signal, config *Config) error {
	var httpServer *http.Server
	if repository, err := repo.CreateLocalRepository(config.RepositoryFolder); err != nil {
		return err
	} else {
		httpServer = &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%d", config.HttpPort),
			Handler: Service(repository)}
	}
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

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
	log.Printf("Starting GS-Bucket service %s @ %s", Version, httpServer.Addr)
	// Run the server
	ServerIsRunning = true
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		ServerIsRunning = false
		return err
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	ServerIsRunning = false
	return nil
}
