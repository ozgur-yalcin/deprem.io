package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const YardimCollection = "yardim"

type Yardim struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	YardimTipi      string             `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"` // Gıda, İlaç, Enkaz, Isınma, Kayıp
	AdSoyad         string             `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon         string             `json:"telefon,omitempty" bson:"telefon,omitempty"`
	YedekTelefonlar []string           `json:"yedekTelefonlar,omitempty" bson:"yedekTelefonlar,omitempty"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty"`
	Adres           string             `json:"adres,omitempty" bson:"adres,omitempty"`
	AdresTarifi     string             `json:"adresTarifi,omitempty" bson:"adresTarifi,omitempty"`
	AcilDurum       string             `json:"acilDurum,omitempty" bson:"acilDurum,omitempty"`
	YardimDurumu    string             `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	KisiSayisi      string             `json:"kisiSayisi,omitempty" bson:"kisiSayisi,omitempty"`
	FizikiDurum     string             `json:"fizikiDurum,omitempty" bson:"fizikiDurum,omitempty"`
	GoogleMapLink   string             `json:"googleMapLink,omitempty" bson:"googleMapLink,omitempty"`
	TweetLink       string             `json:"tweetLink,omitempty" bson:"tweetLink,omitempty"`
	Fields          primitive.M        `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	IPv4            string             `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt       primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt       primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
