package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ozgur-soft/deprem.io/models"
	"github.com/tealeg/xlsx"
)

func Yardimet(w http.ResponseWriter, r *http.Request) {
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
	search := yardimet.Ara(r.Context(), models.Yardimet{
		AdSoyad: r.Form.Get("adSoyad"),
		Sehir:   r.Form.Get("sehir"),
	}, (page-1)*limit, limit)
	if len(search) > 0 {
		response := models.Response{Error: "Bu yardım bildirimi daha önce veritabanımıza eklendi."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.JSON())
		return
	}
	id := yardimet.Kaydet(r.Context(), models.Yardimet{
		YardimTipi:   r.Form.Get("yardimTipi"),
		AdSoyad:      r.Form.Get("adSoyad"),
		Telefon:      r.Form.Get("telefon"),
		Sehir:        r.Form.Get("sehir"),
		Ilce:         r.Form.Get("ilce"),
		HedefSehir:   r.Form.Get("hedefSehir"),
		YardimDurumu: r.Form.Get("yardimDurumu"),
		Aciklama:     r.Form.Get("aciklama"),
		Fields:       r.Form.Get("fields"),
		IPv4:         r.Header.Get("X-Forwarded-For"),
		CreatedAt:    time.Now(),
	})
	if id != "" {
		response := models.Response{Message: "İletişim talebiniz başarıyla alındı"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response.JSON())
		return
	}
	response := models.Response{Error: "Hata! Yardım dökümanı kaydedilemedi!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(response.JSON())
}

func YardimetExport(w http.ResponseWriter, r *http.Request) {
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
	list := yardimet.Ara(r.Context(), models.Yardimet{}, 0, 100000)
	for id, data := range list {
		xlsdata := sheet.AddRow()
		cell := xlsdata.AddCell()
		cell.SetInt(id + 1)
		xlsdata.AddCell().SetString(data.YardimTipi)
		xlsdata.AddCell().SetString(data.AdSoyad)
		xlsdata.AddCell().SetString(data.Telefon)
		xlsdata.AddCell().SetString(data.Sehir)
		xlsdata.AddCell().SetString(data.HedefSehir)
		xlsdata.AddCell().SetString(data.YardimDurumu)
		xlsdata.AddCell().SetString(data.Aciklama)
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
