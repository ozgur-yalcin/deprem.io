package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ozgur-soft/deprem.io/database"
	"github.com/ozgur-soft/deprem.io/models"
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Yardimet(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimRight(r.URL.EscapedPath(), "/")
	if len(strings.Split(path, "/")) == 3 {
		id := regexp.MustCompile(`[^\/]+@`).FindString(path)
		search := database.Search(r.Context(), models.YardimetCollection, primitive.D{{Key: "_id", Value: id}}, 0, 1)
		if len(search) > 0 {
			response, _ := json.MarshalIndent(search[0], " ", " ")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
		response := models.Response{Error: "Yardım kaydı bulunamadı!"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response.JSON())
		return
	}
	http.NotFound(w, r)
}

func YardimetEkle(w http.ResponseWriter, r *http.Request) {
	data := models.Yardimet{}
	json.NewDecoder(r.Body).Decode(&data)
	exists := database.Search(r.Context(), models.YardimetCollection, primitive.D{{Key: "adSoyad", Value: data.AdSoyad}, {Key: "sehir", Value: data.Sehir}}, 0, 1)
	if len(exists) > 0 {
		response := models.Response{Error: "Yardım kaydı daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		w.Write(response.JSON())
		return
	}
	data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	id := database.Add(r.Context(), models.YardimetCollection, data)
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
	search := database.Search(r.Context(), models.YardimetCollection, filter, (page-1)*limit, limit)
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
	search := database.Search(r.Context(), models.YardimetCollection, filter, 0, 100000)
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
	w.Header().Set("Content-Disposition", "attachment; filename=export.xlsx; size="+fmt.Sprintf("%v", buffer.Len()))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeContent(w, r, "export.xlsx", time.Now(), bytes.NewReader(buffer.Bytes()))
}
