package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozgur-soft/deprem.io/database"
	"github.com/ozgur-soft/deprem.io/models"
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Yardimet(ctx *gin.Context) {
	search := database.Search(ctx, models.YardimetCollection, primitive.D{{Key: "postId", Value: ctx.Param("id")}}, 0, 1)
	if len(search) > 0 {
		ctx.JSON(http.StatusOK, search[0])
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Yardım kaydı bulunamadı!"})
}

func YardimetEkle(ctx *gin.Context) {
	data := models.Yardimet{}
	ctx.BindJSON(data)
	exists := database.Search(ctx, models.YardimetCollection, primitive.D{{Key: "adSoyad", Value: data.AdSoyad}, {Key: "sehir", Value: data.Sehir}}, 0, 1)
	if len(exists) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Yardım kaydı daha önce veritabanımıza eklendi."})
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := database.Add(ctx, models.YardimetCollection, data)
	if id != "" {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Yardım kaydı başarıyla alındı", "id": id})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Sunucu hatası"})
}

func YardimetAra(ctx *gin.Context) {
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
	if ctx.Query("sehir") != "" {
		filter = append(filter, primitive.E{Key: "sehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("sehir"), Options: "i"}}}})
	}
	if ctx.Query("hedefSehir") != "" {
		filter = append(filter, primitive.E{Key: "hedefSehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("hedefSehir"), Options: "i"}}}})
	}
	if ctx.Query("aciklama") != "" {
		filter = append(filter, primitive.E{Key: "aciklama", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("aciklama"), Options: "i"}}}})
	}
	if ctx.Query("yardimDurumu") != "" {
		filter = append(filter, primitive.E{Key: "yardimDurumu", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("yardimDurumu"), Options: "i"}}}})
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
	search := database.Search(ctx, models.YardimetCollection, filter, (page-1)*limit, limit)
	ctx.JSON(http.StatusOK, search)
}

func YardimetRapor(ctx *gin.Context) {
	file := xlsx.NewFile()
	rows := []string{"Yardım tipi", "Ad soyad", "Telefon", "Şehir", "Hedef şehir", "Açıklama", "Yardım durumu", "IP adresi", "Oluşturma zamanı", "Güncelleme zamanı"}
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		log.Fatal(err.Error())
	}
	xlshead := sheet.AddRow()
	head := xlshead.AddCell()
	head.Value = "#"
	for i, row := range rows {
		head := xlshead.AddCell()
		head.Value = row
		sheet.SetColWidth(i+1, i+1, 20.0)
	}
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
	if ctx.Query("sehir") != "" {
		filter = append(filter, primitive.E{Key: "sehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("sehir"), Options: "i"}}}})
	}
	if ctx.Query("hedefSehir") != "" {
		filter = append(filter, primitive.E{Key: "hedefSehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("hedefSehir"), Options: "i"}}}})
	}
	if ctx.Query("aciklama") != "" {
		filter = append(filter, primitive.E{Key: "aciklama", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("aciklama"), Options: "i"}}}})
	}
	if ctx.Query("yardimDurumu") != "" {
		filter = append(filter, primitive.E{Key: "yardimDurumu", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("yardimDurumu"), Options: "i"}}}})
	}
	if ctx.Query("ip") != "" {
		filter = append(filter, primitive.E{Key: "ip", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: ctx.Query("ip"), Options: "i"}}}})
	}
	search := database.Search(ctx, models.YardimetCollection, filter, 0, 100000)
	for id, data := range search {
		if parse, err := json.Marshal(data); err == nil {
			info := models.Yardimet{}
			json.Unmarshal(parse, &info)
			xlsdata := sheet.AddRow()
			xlsdata.AddCell().SetInt(id + 1)
			xlsdata.AddCell().SetString(info.YardimTipi)
			xlsdata.AddCell().SetString(info.AdSoyad)
			xlsdata.AddCell().SetString(info.Telefon)
			xlsdata.AddCell().SetString(info.Sehir)
			xlsdata.AddCell().SetString(info.HedefSehir)
			xlsdata.AddCell().SetString(info.Aciklama)
			xlsdata.AddCell().SetString(info.YardimDurumu)
			xlsdata.AddCell().SetString(info.IPv4)
			xlsdata.AddCell().SetString(info.CreatedAt.Time().Format(time.RFC3339))
			xlsdata.AddCell().SetString(info.UpdatedAt.Time().Format(time.RFC3339))
		}
	}
	sheet.AutoFilter = &xlsx.AutoFilter{TopLeftCell: "B1", BottomRightCell: xlsx.GetCellIDStringFromCoords(len(rows)+1, len(search)+1)}
	buffer := new(bytes.Buffer)
	defer buffer.Reset()
	file.Write(buffer)
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
}
