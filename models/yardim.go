package models

import (
	"context"
	"log"

	bson "go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Yardim struct {
	YardimTipi      string       `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"` // Gıda, İlaç, Enkaz, Isınma, Kayıp
	AdSoyad         string       `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon         string       `json:"telefon,omitempty" bson:"telefon,omitempty"`
	YedekTelefonlar string       `json:"yedekTelefonlar,omitempty" bson:"yedekTelefonlar,omitempty"`
	Email           string       `json:"email,omitempty" bson:"email,omitempty"`
	Adres           string       `json:"adres,omitempty" bson:"adres,omitempty"`
	AdresTarifi     string       `json:"adresTarifi,omitempty" bson:"adresTarifi,omitempty"`
	AcilDurum       AcilDurum    `json:"acilDurum,omitempty" bson:"acilDurum,omitempty"`
	YardimDurumu    YardimDurumu `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	KisiSayisi      string       `json:"kisiSayisi,omitempty" bson:"kisiSayisi,omitempty"`
	FizikiDurum     string       `json:"fizikiDurum,omitempty" bson:"fizikiDurum,omitempty"`
	GoogleMapLink   string       `json:"googleMapLink,omitempty" bson:"googleMapLink,omitempty"`
	TweetLink       string       `json:"tweetLink,omitempty" bson:"tweetLink,omitempty"`
	Fields          any          `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	Ip              string       `json:"ip,omitempty" bson:"ip,omitempty"`
}

func yardim(ctx context.Context) (list []Yardim) {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection("yardim")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	if cursor.Err() == nil {
		for cursor.Next(ctx) {
			var data Yardim
			cursor.Decode(&data)
			list = append(list, data)
		}
	}
	return list
}
