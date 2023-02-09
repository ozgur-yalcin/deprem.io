package models

import (
	"context"
	"fmt"
	"log"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

const IletisimCollection = "iletisim"

type Iletisim struct {
	AdSoyad string `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Email   string `json:"email,omitempty" bson:"email,omitempty"`
	Telefon string `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Mesaj   string `json:"mesaj,omitempty" bson:"mesaj,omitempty"`
	Ip      string `json:"ip,omitempty" bson:"ip,omitempty"`
}

func (model *Iletisim) Ara(ctx context.Context, data Iletisim, skip int64, limit int64) (list []Iletisim) {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection(IletisimCollection)
	cursor, err := collection.Find(ctx, data, mongooptions.Find().SetSkip(skip).SetLimit(limit))
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

func (model *Iletisim) Kaydet(ctx context.Context, data Iletisim) string {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
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
