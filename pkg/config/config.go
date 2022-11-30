package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/guionardo/gs-bucket/pkg/logger"
)

type Config struct {
	DataPath string
	Host     string
	Port     int
}

func (config *Config) String() string {
	return fmt.Sprintf("Config{DataPath: %s, Host: %s, Port: %d}", config.DataPath, config.Host, config.Port)
}

func NewConfig() (config *Config, err error) {
	config = &Config{
		Host: "",
		Port: 8080,
	}
	defer func() {
		logger.GetLogger("CONFIG").Printf("config: %v", config)
	}()
	args := getArgsMap()
	if config.DataPath, err = getValue(args, "data-path"); err != nil {
		return
	}
	if host, err := getValue(args, "host"); err == nil {
		config.Host = host
	}
	if port, err := getValue(args, "port"); err == nil {
		if config.Port, err = strconv.Atoi(port); err != nil {
			return config, err
		}
	}

	return
}

func getArgsMap() map[string]string {
	args := make(map[string]string)
	lastKey := ""
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			lastKey = strings.TrimPrefix(arg, "--")
			continue
		}
		if len(lastKey) > 0 {
			args[lastKey] = arg
			lastKey = ""
		}
	}
	return args
}

func getValue(args map[string]string, key string) (string, error) {
	if val, ok := args[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("missing argument --%s", key)
}
