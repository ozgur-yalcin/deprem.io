package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ozgur-soft/deprem.io/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Iletisim(w http.ResponseWriter, r *http.Request) {
	model := new(models.Iletisim)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := model.Search(r.Context(), primitive.D{{Key: "_id", Value: id}}, 0, 1)
	if len(search) > 0 {
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
	model := new(models.Iletisim)
	data := models.Iletisim{}
	json.NewDecoder(r.Body).Decode(&data)
	exists := model.Search(r.Context(), primitive.D{{Key: "adSoyad", Value: data.AdSoyad}, {Key: "email", Value: data.Email}, {Key: "mesaj", Value: data.Mesaj}}, 0, 1)
	if len(exists) > 0 {
		response := models.Response{Error: "İletişim talebi zaten var, lütfen farklı bir talepte bulunun."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		w.Write(response.JSON())
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := model.Add(r.Context(), data)
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
	model := new(models.Iletisim)
	filter := primitive.D{}
	if r.Form.Get("adSoyad") != "" {
		filter = append(filter, primitive.E{Key: "adSoyad", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("adSoyad"), Options: "i"}}}})
	}
	if r.Form.Get("email") != "" {
		filter = append(filter, primitive.E{Key: "email", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("email"), Options: "i"}}}})
	}
	if r.Form.Get("telefon") != "" {
		filter = append(filter, primitive.E{Key: "telefon", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("telefon"), Options: "i"}}}})
	}
	if r.Form.Get("mesaj") != "" {
		filter = append(filter, primitive.E{Key: "mesaj", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("mesaj"), Options: "i"}}}})
	}
	if r.Form.Get("ip") != "" {
		filter = append(filter, primitive.E{Key: "ip", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("ip"), Options: "i"}}}})
	}
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
	search := model.Search(r.Context(), filter, (page-1)*limit, limit)
	response, _ := json.MarshalIndent(search, " ", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
