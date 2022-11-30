package handlers

import (
	"net/http"

	"github.com/guionardo/gs-bucket/pkg/responses"
)

func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	files, err := _repository.GetAll()
	if err != nil {
		responses.WriteResponse(w, http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	responses.WriteResponse(w, http.StatusOK, files)
}
