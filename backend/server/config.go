package server

import (
	"flag"
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
	repoFolder, httpPort := parseFlags()
	if len(repoFolder) == 0 {
		if repoFolder = os.Getenv(RepositoryFolder); len(repoFolder) == 0 {
			err = fmt.Errorf("MISSING %s ENVIRONMENT", RepositoryFolder)
			return
		}
	}

	if httpPort < 80 || httpPort > 65535 {
		if httpPort, _ = strconv.Atoi(os.Getenv(HttpPort)); httpPort < 80 || httpPort > 65535 {
			httpPort = 8080
		}
	}
	cfg = &Config{
		RepositoryFolder: repoFolder,
		HttpPort:         httpPort,
	}
	return
}

var (
	repoFolder string
	httpPort   int
)

func init() {
	flag.StringVar(&repoFolder, "repository", "", "Repository folder")
	flag.IntVar(&httpPort, "port", 0, "HTTP server port")
}

func parseFlags() (repoFolder string, httpPort int) {
	flag.Parse()
	return repoFolder, httpPort
}
