package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const YardimetCollection = "yardimet"

type Yardimet struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	YardimTipi   string             `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"`
	AdSoyad      string             `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon      string             `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Sehir        string             `json:"sehir,omitempty" bson:"sehir,omitempty"`
	Ilce         string             `json:"ilce,omitempty" bson:"ilce,omitempty"`
	HedefSehir   string             `json:"hedefSehir,omitempty" bson:"hedefSehir,omitempty"`
	Aciklama     string             `json:"aciklama,omitempty" bson:"aciklama,omitempty"`
	Fields       primitive.M        `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	YardimDurumu string             `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	IPv4         string             `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt    primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
