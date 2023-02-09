package controllers

import (
	"net/http"

	"github.com/ozgur-soft/deprem.io/models"
)

func Iletisim(w http.ResponseWriter, r *http.Request) {
	iletisim := new(models.Iletisim)
	search := iletisim.Ara(r.Context(), models.Iletisim{
		AdSoyad: r.Form.Get("adSoyad"),
		Email:   r.Form.Get("email"),
		Mesaj:   r.Form.Get("mesaj"),
	})
	if len(search) > 0 {
		response := models.Response{Error: "Bu iletişim talebi zaten var, lütfen farklı bir talepte bulunun."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := iletisim.Kaydet(r.Context(), models.Iletisim{
		AdSoyad: r.Form.Get("adSoyad"),
		Email:   r.Form.Get("email"),
		Telefon: r.Form.Get("telefon"),
		Mesaj:   r.Form.Get("mesaj"),
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
