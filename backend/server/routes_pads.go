package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/render"
	"github.com/guionardo/gs-bucket/domain"
)

// Create pad godoc
//
//	@Summary		Create a pad
//	@Description	Post a file to a pad, accessible for anyone
//	@Tags			pads
//	@Accept			json
//	@Produce		json
//	@Param			api-key				header		string	true	"API Key"
//	@Param			name				query		string	true	"File name"
//	@Param			slug				query		string	false	"Slug or easy name (if not informed, will be used a hash value)"
//	@Param			ttl					query		string	false	"Time to live (i.Ex 300s, 1.5h or 2h45m). Valid time units are: 's', 'm', 'h')"
//	@Param			delete-after-read	query		bool	false	"If informed, the file will be deleted after first download."
//	@Param			private				query		bool	false	"Should use API-KEY header to download file"
//	@Param			content				body		string	true	"Content"
//	@Success		201					{object}	domain.File
//	@Failure		400					{object}	server.ErrResponse
//	@Failure		500					{object}	server.ErrResponse
//	@Failure		507					{object}	server.ErrResponse
//	@Router			/pads [post]
func CreatePad(w http.ResponseWriter, r *http.Request) {
	owner, err := auth.IsAuthorized(r)
	if err != nil {
		renderError(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		renderError(w, r, http.StatusBadRequest, "Required name argument")
		return
	}
	if contentLengthHeader := r.Header.Get("Content-Length"); len(contentLengthHeader) == 0 {
		renderError(w, r, http.StatusBadRequest, "Required Content-Length header")
		return
	}

	ttl, _ := time.ParseDuration(r.URL.Query().Get("ttl"))
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

	file, err := domain.CreateFileFromData(name, body, ttl, owner)
	if r.URL.Query().Has("private") {
		file.SetPrivate(true)
	}
	if slug := r.URL.Query().Get("slug"); len(slug) > 0 {
		_ = file.SetSlug(slug)
	}

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
	RecordMetrics(repository)

	file.StatusCode = http.StatusCreated
	render.Render(w, r, file)
}

// Download pad godoc
//
//	@Summary	Download a pad
//	@Tags		pads
//	@Accept		json
//	@Produce	json
//	@Param		api-key	header		string	false	"API Key used for private pads"
//	@Param		code	path		string	true	"File code"
//	@Success	200		{string}	string
//	@Failure	404		{object}	server.ErrResponse
//	@Failure	500		{object}	server.ErrResponse
//	@Router		/pads/{code} [get]
func GetPad(w http.ResponseWriter, r *http.Request) {
	file := getFile(r)
	if file.Private {
		if _, err := auth.IsAuthorized(r); err != nil {
			renderError(w, r, http.StatusUnauthorized, err.Error())
			return
		}
	}
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
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.Name))
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
//	@Param		api-key	header		string	false	"API Key used for private pads"
//	@Success	200		{object}	[]domain.File
//	@Failure	500		{object}	server.ErrResponse
//	@Router		/pads [get]
func ListPads(w http.ResponseWriter, r *http.Request) {
	repository.Purge()
	pads, err := repository.List()
	if err != nil {
		renderError(w, r, http.StatusInternalServerError,
			fmt.Sprintf("Failed to enumerate pads %v", err))
		return
	}
	filtered := make([]*domain.File, 0, len(pads))
	owner, _ := auth.IsAuthorized(r)
	// Get only the public and private pads
	for _, pad := range pads {
		if !pad.Private || pad.Owner == owner {
			filtered = append(filtered, pad)
		}
	}

	body, _ := json.Marshal(filtered)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", strconv.Itoa(len(body)))
	_, _ = w.Write(body)
}
