package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozgur-soft/deprem.io/database"
	"github.com/ozgur-soft/deprem.io/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Yardim(ctx *gin.Context) {
	search := database.Search(ctx, models.YardimCollection, primitive.D{{Key: "_id", Value: ctx.Param("id")}}, 0, 1)
	if len(search) > 0 {
		ctx.JSON(http.StatusOK, search[0])
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Yardım bildirimi bulunamadı!"})
}

func YardimEkle(ctx *gin.Context) {
	data := models.Yardim{}
	ctx.BindJSON(data)
	exists := database.Search(ctx, models.YardimCollection, primitive.D{{Key: "adSoyad", Value: data.AdSoyad}, {Key: "adres", Value: data.Adres}}, 0, 1)
	if len(exists) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Yardım bildirimi daha önce veritabanımıza eklendi."})
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := database.Add(ctx, models.YardimCollection, data)
	if id != "" {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Yardım bildirimi başarıyla alındı", "id": id})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucu hatası"})
}

func YardimAra(ctx *gin.Context) {
	filter := primitive.D{}
	if ctx.Query("yardimTipi") != "" {
		filter = append(filter, primitive.E{Key: "yardimTipi", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("yardimTipi"), Options: "i"}}}})
	}
	if ctx.Query("adSoyad") != "" {
		filter = append(filter, primitive.E{Key: "adSoyad", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("adSoyad"), Options: "i"}}}})
	}
	if ctx.Query("telefon") != "" {
		filter = append(filter, primitive.E{Key: "telefon", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("telefon"), Options: "i"}}}})
	}
	if ctx.Query("yedekTelefonlar") != "" {
		filter = append(filter, primitive.E{Key: "yedekTelefonlar", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("yedekTelefonlar"), Options: "i"}}}})
	}
	if ctx.Query("email") != "" {
		filter = append(filter, primitive.E{Key: "email", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("email"), Options: "i"}}}})
	}
	if ctx.Query("adres") != "" {
		filter = append(filter, primitive.E{Key: "adres", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("adres"), Options: "i"}}}})
	}
	if ctx.Query("adresTarifi") != "" {
		filter = append(filter, primitive.E{Key: "adresTarifi", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("adresTarifi"), Options: "i"}}}})
	}
	if ctx.Query("acilDurum") != "" {
		filter = append(filter, primitive.E{Key: "acilDurum", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("acilDurum"), Options: "i"}}}})
	}
	if ctx.Query("yardimDurumu") != "" {
		filter = append(filter, primitive.E{Key: "yardimDurumu", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("yardimDurumu"), Options: "i"}}}})
	}
	if ctx.Query("kisiSayisi") != "" {
		filter = append(filter, primitive.E{Key: "kisiSayisi", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("kisiSayisi"), Options: "i"}}}})
	}
	if ctx.Query("fizikiDurum") != "" {
		filter = append(filter, primitive.E{Key: "fizikiDurum", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("fizikiDurum"), Options: "i"}}}})
	}
	if ctx.Query("googleMapLink") != "" {
		filter = append(filter, primitive.E{Key: "googleMapLink", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("googleMapLink"), Options: "i"}}}})
	}
	if ctx.Query("tweetLink") != "" {
		filter = append(filter, primitive.E{Key: "tweetLink", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("tweetLink"), Options: "i"}}}})
	}
	if ctx.Query("ip") != "" {
		filter = append(filter, primitive.E{Key: "ip", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("ip"), Options: "i"}}}})
	}
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if page < 0 {
		page = 0
	}
	if limit <= 10 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	search := database.Search(ctx, models.YardimCollection, filter, (page-1)*limit, limit)
	ctx.JSON(http.StatusOK, search)
}
