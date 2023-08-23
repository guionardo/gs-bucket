package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/guionardo/gs-bucket/domain"
)

// Create user godoc
//
//	@Summary		Create a key for a user
//	@Description	Post a file to a pad, accessible for anyone
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			api-key	header		string	true	"API Key"
//	@Param			user	path		string	true	"User name"
//	@Success		201		{object}	domain.AuthResponse
//	@Failure		400		{object}	server.ErrResponse	"Required user name"
//	@Failure		401		{object}	server.ErrResponse
//	@Failure		500		{object}	server.ErrResponse
//	@Router			/auth/{user} [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if len(auth.masterKey) == 0 {
		renderError(w, r, http.StatusInternalServerError, "Master key is undefined")
		return
	}
	if user := chi.URLParam(r, "user"); user == "" {
		renderError(w, r, http.StatusBadRequest, "user is required")
	} else {
		if key, err := auth.SetKey(user); err == nil {
			render.Render(w, r, &domain.AuthResponse{
				UserName: user,
				Key:      key,
			})
		} else {
			renderError(w, r, http.StatusInternalServerError, err.Error())
		}
	}
}

// List users godoc
//
//	@Summary	List all users allowed to publish
//	@Tags		auth
//	@Produce	json
//	@Param		api-key	header		string	true	"API Key (master key)"
//	@Success	200		{object}	[]domain.AuthResponse
//	@Failure	401		{object}	server.ErrResponse
//	@Failure	500		{object}	server.ErrResponse
//	@Router		/auth/ [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	if owner := r.Context().Value(authOwnerKey); owner != masterUser {
		renderError(w, r, http.StatusForbidden, "Feature allowed only for master user")
		return
	}
	if err := auth.loadKeys(); err != nil {
		renderError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	keys := make([]domain.AuthResponse, len(auth.apiKeys))
	index := 0
	for key, user := range auth.apiKeys {
		keys[index] = domain.AuthResponse{
			UserName: user,
			Key:      key,
		}
		index++
	}
	if content, err := json.Marshal(keys); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(content)
	} else {
		renderError(w, r, http.StatusInternalServerError, err.Error())
	}
}

// Delete user godoc
//
//	@Summary	Delete all keys of user
//	@Tags		auth
//	@Produce	json
//	@Param		api-key	header		string	true	"API Key"
//	@Param		user	path		string	true	"User name"
//	@Success	201		{object}	server.ErrResponse
//	@Failure	401		{object}	server.ErrResponse
//	@Failure	500		{object}	server.ErrResponse
//	@Router		/auth/{user} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if owner := r.Context().Value(authOwnerKey); owner != masterUser {
		renderError(w, r, http.StatusForbidden, "Feature allowed only for master user")
		return
	}
	if err := auth.loadKeys(); err != nil {
		renderError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if userArg := chi.URLParam(r, "user"); userArg == "" {
		renderError(w, r, http.StatusBadRequest, "user is required")
	} else {
		removed := 0
		for key, user := range auth.apiKeys {
			if user == userArg {
				delete(auth.apiKeys, key)
				removed++
			}
		}
		if removed > 0 {
			if err := auth.saveKeys(); err != nil {
				_ = auth.loadKeys()
				removed = 0
			}
		}
		renderError(w, r, http.StatusAccepted, fmt.Sprintf("%d keys removed from user %s", removed, userArg))
	}
}
