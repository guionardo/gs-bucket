package domain

import (
	"fmt"
	"hash/crc64"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-chi/render"
)

type (
	File struct {
		Name            string    `json:"name"`
		Slug            string    `json:"slug"`
		Size            int       `json:"size"`
		CreationDate    time.Time `json:"creation_date"`
		ValidUntil      time.Time `json:"valid_until"`
		ContentType     string    `json:"content_type"`
		DeleteAfterRead bool      `json:"delete_after_read"`
		LastSeen        time.Time `json:"last_seen"`
		SeenCount       int       `json:"seen_count"`
		StatusCode      int       `json:"-"`
		Owner           string    `json:"owner"`
		Private         bool      `json:"private"`
	}
	FileList []*File
)

func (e *File) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func CreateFileFromFile(existantFileName string, ttl time.Duration, owner string) (*File, error) {
	if content, err := os.ReadFile(existantFileName); err != nil {
		return nil, err
	} else {
		return CreateFileFromData(path.Base(existantFileName), content, ttl, owner)
	}
}

func CreateFileFromData(name string, data []byte, ttl time.Duration, owner string) (*File, error) {

	if ttl <= 0 || ttl > 24*time.Hour {
		ttl = time.Hour * 24
	}
	contentType := mimetype.Detect(data).String()

	return &File{
		Name:         path.Base(name),
		Slug:         createSlug(path.Base(name)),
		Size:         len(data),
		CreationDate: time.Now(),
		ValidUntil:   time.Now().Add(ttl),
		ContentType:  contentType,
		Owner:        owner,
	}, nil
}

func createSlug(fileName string) string {
	h := crc64.Checksum([]byte(fileName), crc64.MakeTable(crc64.ECMA))
	return fmt.Sprintf("%x", h)
}

func (file *File) Expired() bool {
	return (!file.ValidUntil.IsZero() && file.ValidUntil.Before(time.Now())) ||
		(file.DeleteAfterRead && file.SeenCount > 0)
}

func (file *File) SetSlug(slug string) error {
	file.Slug = url.PathEscape(strings.ReplaceAll(slug, " ", "_"))
	return nil
}

func (file *File) SetPrivate(isPrivate bool) *File {
	file.Private = isPrivate
	return file
}
