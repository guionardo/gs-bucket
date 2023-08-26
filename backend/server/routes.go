package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-chi/telemetry"
	_ "github.com/guionardo/gs-bucket/backend/docs"
	repo "github.com/guionardo/gs-bucket/backend/repository"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var repository repo.Repository

func renderError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	render.Render(w, r, &ErrResponse{
		HTTPStatusCode: statusCode,
		StatusText:     message,
	})
}

func Service(_repository repo.Repository) http.Handler {
	repository = _repository
	_repository.Purge()
	RecordMetrics(repository)

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		telemetry.Collector(telemetry.Config{
			AllowAny: true,
		}, []string{"/pads", "/swagger"}), // path prefix filters basically records generic http request metrics
		middleware.Compress(5, "application/json", "text/css", "text/html"),
		middleware.Recoverer,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusPermanentRedirect)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	//TODO: Implementar rotas /pads/USUARIO para evitar colisão de hashes/slugs
	r.Route("/pads", func(r chi.Router) {
		r.Get("/", ListPads)
		r.Post("/", CreatePad)
		r.Route("/{code}", func(r chi.Router) {
			r.Use(PadCtx)
			r.Get("/", GetPad)
			r.Delete("/", DeletePad)
		})
	})
	r.Route("/auth", func(r chi.Router) {
		r.Use(BasicAuth)
		r.Post("/{user}", CreateUser)
		r.Get("/", GetUsers)
		r.Delete("/{user}", DeleteUser)
	})
	return r
}

func RecordMetrics(repo repo.Repository) {
	BucketMetrics.RecordFileCount(repo.GetFileCount())
	BucketMetrics.RecordFileSize(repo.GetFileSize())
	BucketMetrics.RecordUserCount(len(auth.apiKeys))
}
