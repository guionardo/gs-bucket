package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TelephoneTan/GoHTTPGzipServer/gzip"
	"github.com/guionardo/gs-bucket/pkg/config"
	"github.com/guionardo/gs-bucket/pkg/logger"
	"github.com/guionardo/gs-bucket/pkg/repositories"
)

var (
	_repository *repositories.Repository
	_config     *config.Config
	_gzipH      *gzip.Handler
	_logger     = logger.GetLogger("handlers")
)

func setupHost(r *http.Request) {
	if repositories.IsHostOk() {
		return
	}
	if strings.Contains(r.Host, "localhost") {
		repositories.SetLastHost(fmt.Sprintf("http://%s", r.Host))
	} else {
		repositories.SetLastHost(fmt.Sprintf("https://%s", r.Host))
	}

}
func MainHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	setupHost(r)
	defer func() {
		// _logger.Printf("%v", w)
		_logger.Printf("%s %-6s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path, time.Since(startTime))
	}()
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPost:
		postHandler(w, r)
	case http.MethodDelete:
		deleteHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupServer(repository *repositories.Repository, config *config.Config) {
	_repository = repository
	_config = config
	h := http.NewServeMux()
	h.HandleFunc("/", MainHandler)
	_gzipH = (&gzip.Handler{Handler: h}).Init()
}

func StartServer() {
	_logger.Printf("Starting server on %s:%d", _config.Host, _config.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", _config.Host, _config.Port), _gzipH))
}
