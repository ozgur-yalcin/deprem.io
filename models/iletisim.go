package deprem

import (
	"context"
	"log"

	bson "go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Iletisim struct {
	AdSoyad string `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Email   string `json:"email,omitempty" bson:"email,omitempty"`
	Telefon string `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Mesaj   string `json:"mesaj,omitempty" bson:"mesaj,omitempty"`
	Ip      string `json:"ip,omitempty" bson:"ip,omitempty"`
}

func iletisim(ctx context.Context) (list []Iletisim) {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("deprem").Collection("iletisim")
	cursor, err := collection.Find(ctx, bson.D{})
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
