package models

import (
	"context"
	"fmt"
	"log"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Yardimet struct {
	YardimTipi   string `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"`
	AdSoyad      string `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon      string `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Sehir        string `json:"sehir,omitempty" bson:"sehir,omitempty"`
	Ilce         string `json:"ilce,omitempty" bson:"ilce,omitempty"` // TODO: ilçe geçici required false yapıldı
	HedefSehir   string `json:"hedefSehir,omitempty" bson:"hedefSehir,omitempty"`
	Aciklama     string `json:"aciklama,omitempty" bson:"aciklama,omitempty"`
	Fields       any    `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	YardimDurumu string `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	Ip           string `json:"ip,omitempty" bson:"ip,omitempty"`
}

func (model *Yardimet) Ara(ctx context.Context, data Yardimet) (list []Yardimet) {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection(YardimetCollection)
	cursor, err := collection.Find(ctx, data)
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

func (model *Yardimet) Kaydet(ctx context.Context, data Yardimet) string {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection(YardimetCollection)
	insert, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", insert.InsertedID)
}
