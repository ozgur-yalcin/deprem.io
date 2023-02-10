package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const IletisimCollection = "iletisim"

type Iletisim struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AdSoyad   string             `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Telefon   string             `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Mesaj     string             `json:"mesaj,omitempty" bson:"mesaj,omitempty"`
	IPv4      string             `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
