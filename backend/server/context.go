package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guionardo/gs-bucket/domain"
)

type key int

const (
	fileKey key = iota
	authOwnerKey
)

func PadCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var file *domain.File
		var name string
		if name = chi.URLParam(r, "code"); name != "" {
			repository.Purge()
			file = repository.Get(name)
		}
		if file == nil {
			renderError(w, r, 404, "NOT FOUND")
			return
		}

		ctx := context.WithValue(r.Context(), fileKey, file)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func getFile(r *http.Request) *domain.File {
	return r.Context().Value(fileKey).(*domain.File)
}
