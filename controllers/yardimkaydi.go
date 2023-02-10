package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozgur-soft/deprem.io/database"
	"github.com/ozgur-soft/deprem.io/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func YardimKaydi(ctx *gin.Context) {
	search := database.Search(ctx, models.YardimKaydiCollection, primitive.D{{Key: "_id", Value: ctx.Param("id")}}, 0, 1)
	if len(search) > 0 {
		ctx.JSON(http.StatusOK, search[0])
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Yardım kaydı bulunamadı!"})
}

func YardimKaydiEkle(ctx *gin.Context) {
	data := models.YardimKaydi{}
	ctx.BindJSON(data)
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := database.Add(ctx, models.YardimKaydiCollection, data)
	if id != "" {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Yardım kaydı başarıyla alındı", "id": id})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucu hatası"})
}
