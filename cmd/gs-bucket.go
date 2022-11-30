package main

import (
	"github.com/guionardo/gs-bucket/pkg/config"
	"github.com/guionardo/gs-bucket/pkg/handlers"
	"github.com/guionardo/gs-bucket/pkg/logger"
	"github.com/guionardo/gs-bucket/pkg/repositories"
)

func main() {
	logger := logger.GetLogger("main")
	logger.Printf("Starting gs-bucket")
	config, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}
	repository, err := repositories.NewRepository(config)
	if err != nil {
		logger.Fatalf("Error loading repository: %v", err)
	}

	handlers.SetupServer(repository, config)
	handlers.StartServer()
}
