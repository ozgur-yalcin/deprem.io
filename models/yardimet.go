package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	cache "github.com/ozgur-soft/deprem.io/cache"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

const YardimetCollection = "yardimet"

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

func (model *Yardimet) Ara(ctx context.Context, data Yardimet, skip int64, limit int64) (list []Yardimet) {
	cachekey := fmt.Sprintf("%v_%v_%v", YardimetCollection, skip, limit)
	if cache.Get(ctx, cachekey) != nil {
		data := cache.Get(ctx, cachekey)
		reader := bytes.NewReader(data)
		decoder := json.NewDecoder(reader)
		decoder.Decode(&list)
		return list
	}
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection(YardimetCollection)
	cursor, err := collection.Find(ctx, data, mongooptions.Find().SetSkip(skip).SetLimit(limit))
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
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI(os.Getenv("MONGOURL")))
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
