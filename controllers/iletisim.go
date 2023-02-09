package controllers

import (
	"net/http"
	"strconv"

	"github.com/ozgur-soft/deprem.io/models"
)

func Iletisim(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.ParseInt(r.Form.Get("page"), 10, 64)
	limit, _ := strconv.ParseInt(r.Form.Get("limit"), 10, 64)
	if page < 0 {
		page = 0
	}
	if limit <= 10 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	iletisim := new(models.Iletisim)
	search := iletisim.Ara(r.Context(), models.Iletisim{
		AdSoyad: r.Form.Get("adSoyad"),
		Email:   r.Form.Get("email"),
		Mesaj:   r.Form.Get("mesaj"),
	}, (page-1)*limit, limit)
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
	response := models.Response{Error: "Hata! Yardım dökümanı kaydedilemedi!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(response.JSON())
}
