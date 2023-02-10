package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ozgur-soft/deprem.io/models"
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Yardimet(w http.ResponseWriter, r *http.Request) {
	model := new(models.Yardimet)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := model.Search(r.Context(), primitive.D{{Key: "_id", Value: id}}, 0, 1)
	if len(search) > 0 {
		response, _ := json.MarshalIndent(search[0], " ", " ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
	response := models.Response{Error: "Yardım kaydı bulunamadı!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(response.JSON())
}

func YardimetEkle(w http.ResponseWriter, r *http.Request) {
	model := new(models.Yardimet)
	data := models.Yardimet{}
	json.NewDecoder(r.Body).Decode(&data)
	exists := model.Search(r.Context(), primitive.D{{Key: "adSoyad", Value: data.AdSoyad}, {Key: "sehir", Value: data.Sehir}}, 0, 1)
	if len(exists) > 0 {
		response := models.Response{Error: "Yardım kaydı daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		w.Write(response.JSON())
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := model.Add(r.Context(), data)
	if id != "" {
		response := models.Response{Message: "Yardım kaydı başarıyla alındı"}
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

func YardimetAra(w http.ResponseWriter, r *http.Request) {
	model := new(models.Yardimet)
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
	if r.Form.Get("sehir") != "" {
		filter = append(filter, primitive.E{Key: "sehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("sehir"), Options: "i"}}}})
	}
	if r.Form.Get("hedefSehir") != "" {
		filter = append(filter, primitive.E{Key: "hedefSehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("hedefSehir"), Options: "i"}}}})
	}
	if r.Form.Get("aciklama") != "" {
		filter = append(filter, primitive.E{Key: "aciklama", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("aciklama"), Options: "i"}}}})
	}
	if r.Form.Get("yardimDurumu") != "" {
		filter = append(filter, primitive.E{Key: "yardimDurumu", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("yardimDurumu"), Options: "i"}}}})
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

func YardimetRapor(w http.ResponseWriter, r *http.Request) {
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
	model := new(models.Yardimet)
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
	if r.Form.Get("sehir") != "" {
		filter = append(filter, primitive.E{Key: "sehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("sehir"), Options: "i"}}}})
	}
	if r.Form.Get("hedefSehir") != "" {
		filter = append(filter, primitive.E{Key: "hedefSehir", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("hedefSehir"), Options: "i"}}}})
	}
	if r.Form.Get("aciklama") != "" {
		filter = append(filter, primitive.E{Key: "aciklama", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("aciklama"), Options: "i"}}}})
	}
	if r.Form.Get("yardimDurumu") != "" {
		filter = append(filter, primitive.E{Key: "yardimDurumu", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("yardimDurumu"), Options: "i"}}}})
	}
	if r.Form.Get("ip") != "" {
		filter = append(filter, primitive.E{Key: "ip", Value: primitive.D{{Key: "$regex", Value: primitive.Regex{Pattern: r.Form.Get("ip"), Options: "i"}}}})
	}
	search := model.Search(r.Context(), filter, 0, 100000)
	for id, data := range search {
		xlsdata := sheet.AddRow()
		xlsdata.AddCell().SetInt(id + 1)
		xlsdata.AddCell().SetString(data.YardimTipi)
		xlsdata.AddCell().SetString(data.AdSoyad)
		xlsdata.AddCell().SetString(data.Telefon)
		xlsdata.AddCell().SetString(data.Sehir)
		xlsdata.AddCell().SetString(data.HedefSehir)
		xlsdata.AddCell().SetString(data.Aciklama)
		xlsdata.AddCell().SetString(data.YardimDurumu)
		xlsdata.AddCell().SetString(data.IPv4)
		xlsdata.AddCell().SetString(data.CreatedAt.Time().Format(time.RFC3339))
		xlsdata.AddCell().SetString(data.UpdatedAt.Time().Format(time.RFC3339))
	}
	sheet.AutoFilter = &xlsx.AutoFilter{TopLeftCell: "B1", BottomRightCell: xlsx.GetCellIDStringFromCoords(len(rows)+1, len(search)+1)}
	buffer := new(bytes.Buffer)
	defer buffer.Reset()
	file.Write(buffer)
	w.Header().Set("Content-Disposition", "attachment; filename=export.xlsx; size="+fmt.Sprintf("%v", buffer.Len()))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeContent(w, r, "export.xlsx", time.Now(), bytes.NewReader(buffer.Bytes()))
}
