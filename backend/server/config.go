package server

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	RepositoryFolder string
	HttpPort         int
}

const (
	RepositoryFolder = "GS_BUCKET_REPOSITORY_FOLDER"
	HttpPort         = "GS_BUCKET_HTTP_PORT"
)

func GetConfig() (cfg *Config, err error) {
	repoFolder := os.Getenv(RepositoryFolder)
	if len(repoFolder) == 0 {
		err = fmt.Errorf("MISSING %s ENVIRONMENT", RepositoryFolder)
		return
	}

	httpPort := 0
	if httpPort, _ = strconv.Atoi(os.Getenv(HttpPort)); httpPort < 80 {
		httpPort = 8080
	}
	cfg = &Config{
		RepositoryFolder: repoFolder,
		HttpPort:         httpPort,
	}
	return
}
