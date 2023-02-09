package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/ozgur-soft/deprem.io/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Iletisim(w http.ResponseWriter, r *http.Request) {
	iletisim := new(models.Iletisim)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := iletisim.Ara(r.Context(), bson.D{{"_id", id}}, 0, 1)
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
	exists := iletisim.Ara(r.Context(), bson.D{{"adSoyad", r.Form.Get("adSoyad")}, {"email", r.Form.Get("email")}, {"mesaj", r.Form.Get("mesaj")}}, 0, 1)
	if len(exists) > 0 {
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
	filter := bson.D{}
	if r.Form.Get("adSoyad") != "" {
		filter = append(filter, bson.E{"adSoyad", bson.D{{"$regex", r.Form.Get("adSoyad")}}})
	}
	if r.Form.Get("email") != "" {
		filter = append(filter, bson.E{"email", bson.D{{"$regex", r.Form.Get("email")}}})
	}
	if r.Form.Get("telefon") != "" {
		filter = append(filter, bson.E{"telefon", bson.D{{"$regex", r.Form.Get("telefon")}}})
	}
	if r.Form.Get("mesaj") != "" {
		filter = append(filter, bson.E{"mesaj", bson.D{{"$regex", r.Form.Get("mesaj")}}})
	}
	if r.Form.Get("ip") != "" {
		filter = append(filter, bson.E{"ip", bson.D{{"$regex", r.Form.Get("ip")}}})
	}
	search := iletisim.Ara(r.Context(), filter, (page-1)*limit, limit)
	response, _ := json.MarshalIndent(search, " ", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
