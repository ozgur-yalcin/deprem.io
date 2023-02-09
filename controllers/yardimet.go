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
	"go.mongodb.org/mongo-driver/bson"
)

func Yardimet(w http.ResponseWriter, r *http.Request) {
	yardimet := new(models.Yardimet)
	id := path.Base(strings.TrimRight(r.URL.EscapedPath(), "/"))
	search := yardimet.Ara(r.Context(), bson.D{{"_id", id}}, 0, 1)
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
	yardimet := new(models.Yardimet)
	data := models.Yardimet{}
	json.NewDecoder(r.Body).Decode(&data)
	exists := yardimet.Ara(r.Context(), bson.D{{"adSoyad", r.Form.Get("adSoyad")}, {"sehir", r.Form.Get("sehir")}}, 0, 1)
	if len(exists) > 0 {
		response := models.Response{Error: "Yardım kaydı daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := yardimet.Ekle(r.Context(), data)
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
	yardimet := new(models.Yardimet)
	filter := bson.D{}
	if r.Form.Get("yardimTipi") != "" {
		filter = append(filter, bson.E{"yardimTipi", r.Form.Get("yardimTipi")})
	}
	if r.Form.Get("adSoyad") != "" {
		filter = append(filter, bson.E{"adSoyad", r.Form.Get("adSoyad")})
	}
	if r.Form.Get("telefon") != "" {
		filter = append(filter, bson.E{"telefon", r.Form.Get("telefon")})
	}
	if r.Form.Get("sehir") != "" {
		filter = append(filter, bson.E{"sehir", r.Form.Get("sehir")})
	}
	if r.Form.Get("hedefSehir") != "" {
		filter = append(filter, bson.E{"hedefSehir", r.Form.Get("hedefSehir")})
	}
	if r.Form.Get("aciklama") != "" {
		filter = append(filter, bson.E{"aciklama", r.Form.Get("aciklama")})
	}
	if r.Form.Get("yardimDurumu") != "" {
		filter = append(filter, bson.E{"yardimDurumu", r.Form.Get("yardimDurumu")})
	}
	if r.Form.Get("ip") != "" {
		filter = append(filter, bson.E{"ip", r.Form.Get("ip")})
	}
	search := yardimet.Ara(r.Context(), filter, (page-1)*limit, limit)
	response, _ := json.MarshalIndent(search, " ", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func YardimetRapor(w http.ResponseWriter, r *http.Request) {
	file := xlsx.NewFile()
	rows := []string{"Yardım tipi", "Ad soyad", "Telefon", "Şehir", "Hedef şehir", "Yardım durumu", "Açıklama", "IP adresi", "Oluşturma zamanı", "Güncelleme zamanı"}
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		log.Println(err.Error())
	}
	xlshead := sheet.AddRow()
	head := xlshead.AddCell()
	head.Value = "#"
	for i, row := range rows {
		head := xlshead.AddCell()
		head.Value = row
		sheet.SetColWidth(i+1, i+1, 20.0)
	}
	yardimet := new(models.Yardimet)
	list := yardimet.Ara(r.Context(), bson.D{}, 0, 100000)
	for id, data := range list {
		xlsdata := sheet.AddRow()
		cell := xlsdata.AddCell()
		cell.SetInt(id + 1)
		xlsdata.AddCell().SetString(data.YardimTipi)
		xlsdata.AddCell().SetString(data.AdSoyad)
		xlsdata.AddCell().SetString(data.Telefon)
		xlsdata.AddCell().SetString(data.Sehir)
		xlsdata.AddCell().SetString(data.HedefSehir)
		xlsdata.AddCell().SetString(data.Aciklama)
		xlsdata.AddCell().SetString(data.YardimDurumu)
		xlsdata.AddCell().SetString(data.IPv4)
		xlsdata.AddCell().SetString(data.CreatedAt.Format(time.RFC3339))
		xlsdata.AddCell().SetString(data.UpdatedAt.Format(time.RFC3339))
	}
	sheet.AutoFilter = &xlsx.AutoFilter{
		TopLeftCell:     "B1",
		BottomRightCell: xlsx.GetCellIDStringFromCoords(len(rows)+1, len(list)+1),
	}
	buffer := new(bytes.Buffer)
	defer buffer.Reset()
	file.Write(buffer)
	w.Header().Set("Content-Disposition", "attachment; filename=export.xlsx; size="+fmt.Sprintf("%v", buffer.Len()))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeContent(w, r, "export.xlsx", time.Now(), bytes.NewReader(buffer.Bytes()))
}
