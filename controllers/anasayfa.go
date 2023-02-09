package controllers

import "net/http"

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("depremio backend"))
}
