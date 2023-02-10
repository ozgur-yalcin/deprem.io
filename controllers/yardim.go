package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ozgur-soft/deprem.io/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Yardim(w http.ResponseWriter, r *http.Request) {
	yardim := new(models.Yardim)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := yardim.Ara(r.Context(), bson.D{{"_id", id}}, 0, 1)
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
	exists := yardim.Ara(r.Context(), bson.D{{"adSoyad", r.Form.Get("adSoyad")}, {"adres", r.Form.Get("adres")}}, 0, 1)
	if len(exists) > 0 {
		response := models.Response{Error: "Yardım bildirimi daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
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
	yardim := new(models.Yardim)
	filter := bson.D{}
	if r.Form.Get("yardimTipi") != "" {
		filter = append(filter, bson.E{"yardimTipi", bson.D{{"$regex", primitive.Regex{r.Form.Get("yardimTipi"), "i"}}}})
	}
	if r.Form.Get("adSoyad") != "" {
		filter = append(filter, bson.E{"adSoyad", bson.D{{"$regex", primitive.Regex{r.Form.Get("adSoyad"), "i"}}}})
	}
	if r.Form.Get("telefon") != "" {
		filter = append(filter, bson.E{"telefon", bson.D{{"$regex", primitive.Regex{r.Form.Get("telefon"), "i"}}}})
	}
	if r.Form.Get("yedekTelefonlar") != "" {
		filter = append(filter, bson.E{"yedekTelefonlar", bson.D{{"$regex", primitive.Regex{r.Form.Get("yedekTelefonlar"), "i"}}}})
	}
	if r.Form.Get("email") != "" {
		filter = append(filter, bson.E{"email", bson.D{{"$regex", primitive.Regex{r.Form.Get("email"), "i"}}}})
	}
	if r.Form.Get("adres") != "" {
		filter = append(filter, bson.E{"adres", bson.D{{"$regex", primitive.Regex{r.Form.Get("adres"), "i"}}}})
	}
	if r.Form.Get("adresTarifi") != "" {
		filter = append(filter, bson.E{"adresTarifi", bson.D{{"$regex", primitive.Regex{r.Form.Get("adresTarifi"), "i"}}}})
	}
	if r.Form.Get("acilDurum") != "" {
		filter = append(filter, bson.E{"acilDurum", bson.D{{"$regex", primitive.Regex{r.Form.Get("acilDurum"), "i"}}}})
	}
	if r.Form.Get("yardimDurumu") != "" {
		filter = append(filter, bson.E{"yardimDurumu", bson.D{{"$regex", primitive.Regex{r.Form.Get("yardimDurumu"), "i"}}}})
	}
	if r.Form.Get("kisiSayisi") != "" {
		filter = append(filter, bson.E{"kisiSayisi", bson.D{{"$regex", primitive.Regex{r.Form.Get("kisiSayisi"), "i"}}}})
	}
	if r.Form.Get("fizikiDurum") != "" {
		filter = append(filter, bson.E{"fizikiDurum", bson.D{{"$regex", primitive.Regex{r.Form.Get("fizikiDurum"), "i"}}}})
	}
	if r.Form.Get("googleMapLink") != "" {
		filter = append(filter, bson.E{"googleMapLink", bson.D{{"$regex", primitive.Regex{r.Form.Get("googleMapLink"), "i"}}}})
	}
	if r.Form.Get("tweetLink") != "" {
		filter = append(filter, bson.E{"tweetLink", bson.D{{"$regex", primitive.Regex{r.Form.Get("tweetLink"), "i"}}}})
	}
	if r.Form.Get("ip") != "" {
		filter = append(filter, bson.E{"ip", bson.D{{"$regex", primitive.Regex{r.Form.Get("ip"), "i"}}}})
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
	search := yardim.Ara(r.Context(), filter, (page-1)*limit, limit)
	response, _ := json.MarshalIndent(search, " ", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
