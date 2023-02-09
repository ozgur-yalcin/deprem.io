package controllers

import (
	"net/http"

	"github.com/ozgur-soft/deprem.io/models"
)

func Yardimet(w http.ResponseWriter, r *http.Request) {
	yardimet := new(models.Yardimet)
	search := yardimet.Ara(r.Context(), models.Yardimet{
		AdSoyad: r.Form.Get("adSoyad"),
		Sehir:   r.Form.Get("sehir"),
	}, 0, 10)
	if len(search) > 0 {
		response := models.Response{Error: "Bu iletişim talebi zaten var, lütfen farklı bir talepte bulunun."}
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
	w.Write([]byte("Hata! Yardım dökümanı kaydedilemedi!"))
}
