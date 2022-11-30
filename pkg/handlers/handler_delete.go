package handlers

import (
	"net/http"
	"path"

	"github.com/guionardo/gs-bucket/pkg/responses"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	file, err := _repository.GetFile(path.Base(r.URL.Path))
	if err != nil {
		responses.WriteResponse(w, http.StatusNotFound, responses.ErrorResponse{Error: err.Error()})
		return
	}
	_repository.DeleteFile(file)
	responses.WriteResponse(w, http.StatusAccepted, responses.PostResponse{
		Success:      true,
		Message:      "File deleted successfully",
		FileName:     file.Name,
		HashFileName: file.Code,
	})
	return
}
