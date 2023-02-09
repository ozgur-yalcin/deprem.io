package controllers

import (
	"net/http"

	"github.com/ozgur-soft/deprem.io/models"
)

func Yardim(w http.ResponseWriter, r *http.Request) {
	yardim := new(models.Yardim)
	search := yardim.Ara(r.Context(), models.Yardim{
		AdSoyad: r.Form.Get("adSoyad"),
		Adres:   r.Form.Get("adres"),
	}, 0, 10)
	if len(search) > 0 {
		response := models.Response{Error: "Bu yardım bildirimi daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := yardim.Kaydet(r.Context(), models.Yardim{
		YardimTipi:    r.Form.Get("yardimTipi"),
		AdSoyad:       r.Form.Get("adSoyad"),
		Telefon:       r.Form.Get("telefon"),
		Email:         r.Form.Get("email"),
		Adres:         r.Form.Get("adres"),
		AcilDurum:     r.Form.Get("acilDurum"),
		AdresTarifi:   r.Form.Get("adresTarifi"),
		YardimDurumu:  "bekleniyor",
		KisiSayisi:    r.Form.Get("kisiSayisi"),
		FizikiDurum:   r.Form.Get("fizikiDurum"),
		TweetLink:     r.Form.Get("tweetLink"),
		GoogleMapLink: r.Form.Get("googleMapLink"),
		Fields:        r.Form.Get("fields"),
		Ip:            r.Header.Get("X-Forwarded-For"),
	})
	if id != "" {
		response := models.Response{Message: "Yardım talebiniz başarıyla alındı"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response.JSON())
		return
	}
	w.Write([]byte("Hata! Yardım dökümanı kaydedilemedi!"))
}
