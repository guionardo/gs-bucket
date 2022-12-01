package handlers

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/guionardo/gs-bucket/pkg/responses"
)

// GET /file_hash
func getHandler(w http.ResponseWriter, r *http.Request) {
	file_hash := path.Base(r.URL.Path)
	switch file_hash {
	case "", "/":
		getHomeHandler(w, r)
		return
	case "markdown.css":
		getMarkdownCssHandler(w, r)
		return
	case "store":
		getIndexHandler(w, r)
		return
	}
	if len(file_hash) == 0 || file_hash == "/" {
	}

	file, err := _repository.GetFile(file_hash)
	if err != nil {
		responses.WriteResponse(w, http.StatusNotFound, responses.ErrorResponse{Error: err.Error()})
		return
	}
	if file.ValidUntil.Before(time.Now()) {
		_repository.DeleteFile(file)
		responses.WriteResponse(w, http.StatusNotFound, responses.ErrorResponse{Error: "file expired"})
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	w.Header().Set("Content-Type", file.MimeType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(file.Content)))
	w.Header().Set("Expires", file.ValidUntil.Format(time.RFC1123))

	w.WriteHeader(http.StatusOK)
	w.Write(file.Content)
}
