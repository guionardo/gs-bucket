package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var _loggers = make(map[string]*log.Logger)

func GetLogger(name string) (logger *log.Logger) {
	name = fmt.Sprintf("%-9s", strings.ToUpper(name))
	var ok bool
	if logger, ok = _loggers[name]; ok {
		return
	}
	logger = log.New(os.Stdout, name, log.LstdFlags)
	_loggers[name] = logger
	return
}
