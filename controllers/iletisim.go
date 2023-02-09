package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/ozgur-soft/deprem.io/models"
)

func Iletisim(w http.ResponseWriter, r *http.Request) {
	iletisim := new(models.Iletisim)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := iletisim.Ara(r.Context(), models.Iletisim{Id: id}, 0, 1)
	if len(search) == 1 {
		response, _ := json.MarshalIndent(search[0], " ", " ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
	response := models.Response{Error: "İletişim talebi bulunamadı!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(response.JSON())
}

func IletisimEkle(w http.ResponseWriter, r *http.Request) {
	iletisim := new(models.Iletisim)
	data := models.Iletisim{}
	json.NewDecoder(r.Body).Decode(&data)
	search := iletisim.Ara(r.Context(), models.Iletisim{AdSoyad: data.AdSoyad, Email: data.Email, Mesaj: data.Mesaj}, 0, 1)
	if len(search) > 0 {
		response := models.Response{Error: "İletişim talebi zaten var, lütfen farklı bir talepte bulunun."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := iletisim.Ekle(r.Context(), data)
	if id != "" {
		response := models.Response{Message: "İletişim talebi başarıyla alındı"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response.JSON())
		return
	}
	response := models.Response{Error: "Hata!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(response.JSON())
}

func IletisimAra(w http.ResponseWriter, r *http.Request) {
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
	id := iletisim.Ekle(r.Context(), models.Iletisim{
		AdSoyad: r.Form.Get("adSoyad"),
		Email:   r.Form.Get("email"),
		Telefon: r.Form.Get("telefon"),
		Mesaj:   r.Form.Get("mesaj"),
		IPv4:    r.Header.Get("X-Forwarded-For"),
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
