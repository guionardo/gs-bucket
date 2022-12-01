package responses

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	ErrorResponse struct {
		Error string `json:"error"`
	}

	PostResponse struct {
		Success      bool      `json:"success"`
		Message      string    `json:"message"`
		FileName     string    `json:"fileName"`
		HashFileName string    `json:"hashFileName,omitempty"`
		ValidUntil   time.Time `json:"validUntil,omitempty"`
		URL          string    `json:"url,omitempty"`
	}
)

func WriteResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}