package handlers

import (
	"net/http"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

func getMimeType(r *http.Request, body []byte) string {
	// if mime := r.Header.Get("Content-Type"); len(mime) > 0 {
	// 	return mime
	// }
	return mimetype.Detect(body).String()
}

func getTTL(r *http.Request) (time.Duration, error) {
	ttlRequest := ""
	if ttlRequest = r.URL.Query().Get("ttl"); len(ttlRequest) == 0 {
		if ttlRequest = r.Header.Get("ttl"); len(ttlRequest) == 0 {
			ttlRequest = "24h"
		}
	}

	return time.ParseDuration(ttlRequest)
}
