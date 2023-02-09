package controllers

import (
	"net/http"
	"strconv"

	"github.com/ozgur-soft/deprem.io/models"
)

func Yardimet(w http.ResponseWriter, r *http.Request) {
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
	yardimet := new(models.Yardimet)
	search := yardimet.Ara(r.Context(), models.Yardimet{
		AdSoyad: r.Form.Get("adSoyad"),
		Sehir:   r.Form.Get("sehir"),
	}, (page-1)*limit, limit)
	if len(search) > 0 {
		response := models.Response{Error: "Bu yardım bildirimi daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := yardimet.Kaydet(r.Context(), models.Yardimet{
		YardimTipi:   r.Form.Get("yardimTipi"),
		AdSoyad:      r.Form.Get("adSoyad"),
		Telefon:      r.Form.Get("telefon"),
		Sehir:        r.Form.Get("sehir"),
		Ilce:         r.Form.Get("ilce"),
		HedefSehir:   r.Form.Get("hedefSehir"),
		YardimDurumu: r.Form.Get("yardimDurumu"),
		Aciklama:     r.Form.Get("aciklama"),
		Fields:       r.Form.Get("fields"),
		Ip:           r.Header.Get("X-Forwarded-For"),
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
