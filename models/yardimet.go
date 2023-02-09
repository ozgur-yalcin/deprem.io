package models

import (
	"context"
	"log"

	bson "go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Yardimet struct {
	YardimTipi   string       `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"`
	AdSoyad      string       `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon      string       `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Sehir        string       `json:"sehir,omitempty" bson:"sehir,omitempty"`
	Ilce         string       `json:"ilce,omitempty" bson:"ilce,omitempty"` // TODO: ilçe geçici required false yapıldı
	HedefSehir   string       `json:"hedefSehir,omitempty" bson:"hedefSehir,omitempty"`
	Aciklama     string       `json:"aciklama,omitempty" bson:"aciklama,omitempty"`
	Fields       any          `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	YardimDurumu YardimDurumu `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	Ip           string       `json:"ip,omitempty" bson:"ip,omitempty"`
}

func yardimet(ctx context.Context) (list []Yardimet) {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection("yardimet")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	if cursor.Err() == nil {
		for cursor.Next(ctx) {
			var data Yardimet
			cursor.Decode(&data)
			list = append(list, data)
		}
	}
	return list
}
