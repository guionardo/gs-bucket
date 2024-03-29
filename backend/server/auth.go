package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Authorization struct {
	masterKey string
	apiKeys   map[string]string
	lock      sync.RWMutex
}

const (
	masterUser = "master"
	keyFile    = "users.key"
	chars      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	auth *Authorization = &Authorization{
		apiKeys: make(map[string]string),
	}
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GetAuthorization() *Authorization {
	return auth
}

func (a *Authorization) SetMasterKeyFile(fileName string) error {
	if masterKey, err := os.ReadFile(fileName); err == nil {
		a.masterKey = strings.TrimRight(string(masterKey), "\n")
		if len(a.masterKey) == 0 {
			log.Printf("FAILED getting master key from empty file: %s", fileName)
			return fmt.Errorf("empty master key file %s", fileName)
		}
		log.Printf("Master key was set")
		return nil
	} else {
		log.Printf("FAILED to read master key file: %v", err)
		return err
	}
}

func (a *Authorization) IsAuthorized(r *http.Request) (owner string, err error) {
	if len(a.masterKey) == 0 {
		if err = a.SetMasterKeyFile("master.key"); err != nil {
			return
		}
	}
	key := r.Header.Get("api-key")
	if len(key) == 0 {
		err = errors.New("'api-key' header is required")
		return
	}
	if key == a.masterKey {
		log.Print("Accepted masterkey request")
		owner = masterUser
		return
	}
	if len(a.apiKeys) == 0 {
		_ = a.loadKeys()
	}
	if user, ok := a.apiKeys[key]; ok {
		log.Printf("Accepted key from user %s", user)
		owner = user
		return
	}
	log.Printf("Invalid key %s", key)
	err = errors.New("invalid key")
	return
}

func (a *Authorization) loadKeys() (err error) {
	var data []byte
	if data, err = os.ReadFile(keyFile); err == nil {
		if err = json.Unmarshal(data, &a.apiKeys); err != nil {
			log.Printf("Failed to read keys - %v", err)
			a.apiKeys = make(map[string]string)
		}
	}
	return
}

func (a *Authorization) saveKeys() error {
	if data, err := json.Marshal(a.apiKeys); err != nil {
		log.Printf("Failed to marshaling keys - %v", err)
		return err
	} else {
		err = os.WriteFile(keyFile, data, 0600)
		if err != nil {
			log.Printf("Failed to write keys file - %v", err)
			return err
		}
	}
	return nil
}

func (a *Authorization) SetKey(user string) (key string, err error) {
	if key == masterUser {
		err = errors.New("you cannot redefine the master key")
		return
	}
	a.lock.Lock()
	defer a.lock.Unlock()

	_ = a.loadKeys()

	newKey := generateKey(30)
	a.apiKeys[newKey] = user
	if err = a.saveKeys(); err != nil {
		delete(a.apiKeys, newKey)
		return "", err
	}

	return newKey, nil
}

func generateKey(size int) string {
	key := make([]byte, size)
	for i := range key {
		key[i] = chars[seededRand.Intn(len(chars)-1)]
	}
	return string(key)
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		owner, err := auth.IsAuthorized(r)
		if err != nil || owner != masterUser {
			renderError(w, r, http.StatusUnauthorized, "UNAUTHORIZED")
			return
		}
		newContext := context.WithValue(r.Context(), authOwnerKey, owner)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}
