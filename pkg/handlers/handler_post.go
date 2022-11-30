package handlers

import (
	"io/ioutil"
	"net/http"
	"path"

	"github.com/guionardo/gs-bucket/pkg/repositories"
	"github.com/guionardo/gs-bucket/pkg/responses"
)

// POST /file_name
func postHandler(w http.ResponseWriter, r *http.Request) {
	originalFile := path.Base(r.URL.Path)
	if len(originalFile) == 0 {
		responses.WriteResponse(w, http.StatusBadRequest, responses.ErrorResponse{Error: "expected POST /[file name]"})
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.WriteResponse(w, http.StatusBadRequest, responses.ErrorResponse{Error: "expected BODY in request"})
		return
	}

	ttl, err := getTTL(r)
	if err != nil {
		responses.WriteResponse(w, http.StatusBadRequest, responses.ErrorResponse{Error: "invalid ttl"})
		return
	}
	mimetype := getMimeType(r, body)
	file := repositories.CreateFile(originalFile, body, mimetype, ttl)

	if err = _repository.SaveFile(file); err != nil {
		responses.WriteResponse(w, http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	responses.WriteResponse(w, http.StatusOK, responses.PostResponse{
		Success:      true,
		Message:      "File uploaded successfully",
		FileName:     file.Name,
		HashFileName: file.Code,
	})
	return
}