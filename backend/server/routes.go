package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/guionardo/gs-bucket/backend/docs"
	repo "github.com/guionardo/gs-bucket/backend/repository"
	"github.com/guionardo/gs-bucket/domain"
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

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Compress(5, "application/json", "text/css", "text/html"),
		middleware.Recoverer,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusPermanentRedirect)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	r.Route("/pads", func(r chi.Router) {
		r.Get("/", ListPads)
		r.Post("/", CreatePad)
		r.Route("/{code}", func(r chi.Router) {
			r.Use(PadCtx)
			r.Get("/", GetPad)
			r.Delete("/", DeletePad)
		})
	})
	return r
}

// Create pad godoc
//
//	@Summary		Create a pad
//	@Description	Post a file to a pad, accessible for anyone
//	@Tags			pads
//	@Accept			json
//	@Produce		json
//	@Param			name				query		string	true	"File name"
//	@Param			ttl					query		string	false	"Time to live"
//	@Param			delete-after-read	query		bool	false	"If informed, the file will be deleted after first download"
//	@Param			content				body		string	true	"Content"
//	@Success		201					{object}	domain.File
//	@Failure		400					{object}	server.ErrResponse
//	@Failure		404					{object}	server.ErrResponse
//	@Failure		500					{object}	server.ErrResponse
//	@Router			/pads [post]
func CreatePad(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		renderError(w, r, http.StatusBadRequest, "Required name argument")
		return
	}

	ttl, _ := time.ParseDuration(chi.URLParam(r, "ttl"))
	deleteAfterRead := r.URL.Query().Has("delete-after-read")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		renderError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if len(body) == 0 {
		renderError(w, r, http.StatusBadRequest, "Required body request")
		return
	}

	file, err := domain.CreateFileFromData(name, body, ttl)

	file.DeleteAfterRead = deleteAfterRead

	if err != nil {
		renderError(w, r, http.StatusInternalServerError,
			fmt.Sprintf("File creation error %v", err))
		return
	}
	if err = repository.SaveFile(file, body); err != nil {
		renderError(w, r, http.StatusInsufficientStorage,
			fmt.Sprintf("File saving error %v", err))
		return
	}
	file.StatusCode = http.StatusCreated
	render.Render(w, r, file)
}

// Download pad godoc
//
//	@Summary	Download a pad
//	@Tags		pads
//	@Accept		json
//	@Produce	json
//	@Param		code	path		string	true	"File code"
//	@Success	200		{string}	string
//	@Failure	404		{object}	server.ErrResponse
//	@Failure	500		{object}	server.ErrResponse
//	@Router		/pads/{code} [get]
func GetPad(w http.ResponseWriter, r *http.Request) {
	file := getFile(r)
	data, _ := repository.Load(file.Slug)
	file.SeenCount++
	file.LastSeen = time.Now()
	if err := repository.Save(); err != nil {
		renderError(w, r, http.StatusInternalServerError, fmt.Sprintf("Failed to save pad %v", err))
		return
	}
	if file.DeleteAfterRead {
		file.ValidUntil = time.Now()
	}
	w.Header().Add("Content-Type", file.ContentType)
	w.Header().Add("Date", file.CreationDate.Format(time.RFC1123Z))
	w.Header().Add("Expires", file.ValidUntil.Format(time.RFC1123Z))
	_, _ = w.Write(data)
}

// Delete pad godoc
//
//	@Summary	Delete a pad
//	@Tags		pads
//	@Accept		json
//	@Produce	json
//	@Param		code	path		string	true	"File code"
//	@Success	200		{string}	string
//	@Failure	404		{object}	server.ErrResponse
//	@Failure	500		{object}	server.ErrResponse
//	@Router		/pads/{code} [delete]
func DeletePad(w http.ResponseWriter, r *http.Request) {
	file := getFile(r)
	if err := repository.Delete(file.Slug); err == nil {
		w.WriteHeader(http.StatusAccepted)
	} else {
		renderError(w, r, http.StatusInternalServerError,
			fmt.Sprintf("Failed to delete file %v", err))
	}
}

// List pads godoc
//
//	@Summary	List pads
//	@Tags		pads
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]domain.File
//	@Failure	404	{object}	server.ErrResponse
//	@Failure	500	{object}	server.ErrResponse
//	@Router		/pads [get]
func ListPads(w http.ResponseWriter, r *http.Request) {
	repository.Purge()
	pads, err := repository.List()
	if err != nil {
		renderError(w, r, http.StatusInternalServerError,
			fmt.Sprintf("Failed to enumerate pads %v", err))
	} else {
		w.Header().Add("Content-Type", "application/json")
		body, _ := json.Marshal(pads)
		_, _ = w.Write(body)
	}
}
