package controllers

import (
	"net/http"

	"github.com/ozgur-soft/deprem.io/cache"
)

func Flushall(w http.ResponseWriter, r *http.Request) {
	cache.Flush(r.Context())
	w.Write([]byte("OK"))
}
