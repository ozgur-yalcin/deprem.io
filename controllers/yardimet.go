package controllers

import (
	"net/http"
	"strconv"

	"github.com/ozgur-soft/deprem.io/models"
)

func Yardimet(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.ParseInt(r.Form.Get("skip"), 10, 64)
	limit, _ := strconv.ParseInt(r.Form.Get("limit"), 10, 64)
	if skip < 0 {
		skip = 0
	}
	if limit <= 10 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	yardimet := new(models.Yardimet)
	search := yardimet.Ara(r.Context(), models.Yardimet{
		AdSoyad: r.Form.Get("adSoyad"),
		Sehir:   r.Form.Get("sehir"),
	}, skip, limit)
	if len(search) > 0 {
		response := models.Response{Error: "Bu yardım bildirimi daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := yardimet.Kaydet(r.Context(), models.Yardimet{
		AdSoyad: r.Form.Get("adSoyad"),
		Telefon: r.Form.Get("telefon"),
		Ip:      r.Header.Get("X-Forwarded-For"),
	})
	if id != "" {
		response := models.Response{Message: "İletişim talebiniz başarıyla alındı"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response.JSON())
		return
	}
	response := models.Response{Error: "Hata! Yardım dökümanı kaydedilemedi!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(response.JSON())
	w.Write([]byte(""))
}
