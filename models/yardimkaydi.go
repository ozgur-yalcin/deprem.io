package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const YardimKaydiCollection = "yardimkaydi"

type YardimKaydi struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PostId    primitive.ObjectID `json:"postId,omitempty" bson:"postId,omitempty"`
	AdSoyad   string             `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon   string             `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	SonDurum  string             `json:"sonDurum,omitempty" bson:"sonDurum,omitempty"`
	Aciklama  string             `json:"aciklama,omitempty" bson:"aciklama,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
