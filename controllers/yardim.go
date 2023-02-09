package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/ozgur-soft/deprem.io/models"
)

func Yardim(w http.ResponseWriter, r *http.Request) {
	yardim := new(models.Yardim)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := yardim.Ara(r.Context(), models.Yardim{Id: id}, 0, 1)
	if len(search) == 1 {
		response, _ := json.MarshalIndent(search[0], " ", " ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
	response := models.Response{Error: "Yardım bildirimi bulunamadı!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(response.JSON())
}

func YardimEkle(w http.ResponseWriter, r *http.Request) {
	yardim := new(models.Yardim)
	data := models.Yardim{}
	json.NewDecoder(r.Body).Decode(&data)
	search := yardim.Ara(r.Context(), models.Yardim{AdSoyad: data.AdSoyad, Adres: data.Adres}, 0, 1)
	if len(search) > 0 {
		response := models.Response{Error: "Yardım bildirimi daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := yardim.Ekle(r.Context(), data)
	if id != "" {
		response := models.Response{Message: "Yardım bildirimi başarıyla alındı"}
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

func YardimAra(w http.ResponseWriter, r *http.Request) {
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
	yardim := new(models.Yardim)
	search := yardim.Ara(r.Context(), models.Yardim{
		AdSoyad: r.Form.Get("adSoyad"),
		Adres:   r.Form.Get("adres"),
	}, (page-1)*limit, limit)
	response, _ := json.MarshalIndent(search, " ", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
