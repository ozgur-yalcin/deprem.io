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

func Yardim(w http.ResponseWriter, r *http.Request) {
	model := new(models.Yardim)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := model.Search(r.Context(), primitive.D{{Key: "_id", Value: id}}, 0, 1)
	if len(search) > 0 {
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
	model := new(models.Yardim)
	data := models.Yardim{}
	json.NewDecoder(r.Body).Decode(&data)
	exists := model.Search(r.Context(), primitive.D{{Key: "adSoyad", Value: data.AdSoyad}, {Key: "adres", Value: data.Adres}}, 0, 1)
	if len(exists) > 0 {
		response := models.Response{Error: "Yardım bildirimi daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := model.Insert(r.Context(), data)
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
	model := new(models.Yardim)
	filter := primitive.D{}
	if r.Form.Get("yardimTipi") != "" {
		filter = append(filter, primitive.E{Key: "yardimTipi", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("yardimTipi"), Options: "i"}}}})
	}
	if r.Form.Get("adSoyad") != "" {
		filter = append(filter, primitive.E{Key: "adSoyad", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("adSoyad"), Options: "i"}}}})
	}
	if r.Form.Get("telefon") != "" {
		filter = append(filter, primitive.E{Key: "telefon", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("telefon"), Options: "i"}}}})
	}
	if r.Form.Get("yedekTelefonlar") != "" {
		filter = append(filter, primitive.E{Key: "yedekTelefonlar", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("yedekTelefonlar"), Options: "i"}}}})
	}
	if r.Form.Get("email") != "" {
		filter = append(filter, primitive.E{Key: "email", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("email"), Options: "i"}}}})
	}
	if r.Form.Get("adres") != "" {
		filter = append(filter, primitive.E{Key: "adres", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("adres"), Options: "i"}}}})
	}
	if r.Form.Get("adresTarifi") != "" {
		filter = append(filter, primitive.E{Key: "adresTarifi", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("adresTarifi"), Options: "i"}}}})
	}
	if r.Form.Get("acilDurum") != "" {
		filter = append(filter, primitive.E{Key: "acilDurum", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("acilDurum"), Options: "i"}}}})
	}
	if r.Form.Get("yardimDurumu") != "" {
		filter = append(filter, primitive.E{Key: "yardimDurumu", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("yardimDurumu"), Options: "i"}}}})
	}
	if r.Form.Get("kisiSayisi") != "" {
		filter = append(filter, primitive.E{Key: "kisiSayisi", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("kisiSayisi"), Options: "i"}}}})
	}
	if r.Form.Get("fizikiDurum") != "" {
		filter = append(filter, primitive.E{Key: "fizikiDurum", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("fizikiDurum"), Options: "i"}}}})
	}
	if r.Form.Get("googleMapLink") != "" {
		filter = append(filter, primitive.E{Key: "googleMapLink", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("googleMapLink"), Options: "i"}}}})
	}
	if r.Form.Get("tweetLink") != "" {
		filter = append(filter, primitive.E{Key: "tweetLink", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("tweetLink"), Options: "i"}}}})
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
