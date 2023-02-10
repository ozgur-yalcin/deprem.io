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

func Iletisim(ctx *gin.Context) {
	search := database.Search(ctx, models.IletisimCollection, primitive.D{{Key: "postId", Value: ctx.Param("id")}}, 0, 1)
	if len(search) > 0 {
		ctx.JSON(http.StatusOK, search[0])
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "İletişim talebi bulunamadı!"})
}

func IletisimEkle(ctx *gin.Context) {
	data := models.Iletisim{}
	ctx.BindJSON(data)
	exists := database.Search(ctx, models.IletisimCollection, primitive.D{{Key: "adSoyad", Value: data.AdSoyad}, {Key: "email", Value: data.Email}, {Key: "mesaj", Value: data.Mesaj}}, 0, 1)
	if len(exists) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "İletişim talebi zaten var, lütfen farklı bir talepte bulunun."})
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := database.Add(ctx, models.IletisimCollection, data)
	if id != "" {
		ctx.JSON(http.StatusCreated, gin.H{"message": "İletişim talebi başarıyla alındı"})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucu hatası"})
}

func IletisimAra(ctx *gin.Context) {
	filter := primitive.D{}
	if ctx.Query("adSoyad") != "" {
		filter = append(filter, primitive.E{Key: "adSoyad", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("adSoyad"), Options: "i"}}}})
	}
	if ctx.Query("email") != "" {
		filter = append(filter, primitive.E{Key: "email", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("email"), Options: "i"}}}})
	}
	if ctx.Query("telefon") != "" {
		filter = append(filter, primitive.E{Key: "telefon", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("telefon"), Options: "i"}}}})
	}
	if ctx.Query("mesaj") != "" {
		filter = append(filter, primitive.E{Key: "mesaj", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("mesaj"), Options: "i"}}}})
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
	search := database.Search(ctx, models.IletisimCollection, filter, (page-1)*limit, limit)
	ctx.JSON(http.StatusOK, search)
}
