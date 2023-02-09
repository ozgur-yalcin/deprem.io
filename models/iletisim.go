package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	cache "github.com/ozgur-soft/deprem.io/cache"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

const IletisimCollection = "iletisim"

type Iletisim struct {
	Id        any       `json:"_id,omitempty" bson:"_id,omitempty"`
	AdSoyad   string    `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Telefon   string    `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Mesaj     string    `json:"mesaj,omitempty" bson:"mesaj,omitempty"`
	IPv4      string    `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (model *Iletisim) Ara(ctx context.Context, search Iletisim, skip int64, limit int64) (list []Iletisim) {
	cachekey := fmt.Sprintf("%v_%v_%v", IletisimCollection, skip, limit)
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
	collection := client.Database("deprem").Collection(IletisimCollection)
	cursor, err := collection.Find(ctx, search, mongooptions.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	if cursor.Err() == nil {
		for cursor.Next(ctx) {
			var data Iletisim
			cursor.Decode(&data)
			list = append(list, data)
		}
	}
	return list
}

func (model *Iletisim) Ekle(ctx context.Context, data Iletisim) string {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection(IletisimCollection)
	insert, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", insert.InsertedID)
}
