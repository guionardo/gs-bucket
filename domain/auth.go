package domain

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type AuthResponse struct {
	UserName string `json:"user"`
	Key      string `json:"key"`
}

func (e *AuthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

func (e *AuthResponse) String() string {
	return fmt.Sprintf("%s:%s", e.UserName, e.Key)
}
