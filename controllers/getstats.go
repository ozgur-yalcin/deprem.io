package controllers

import (
	"net/http"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("-"))
}
