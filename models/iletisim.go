package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ozgur-soft/deprem.io/cache"
	"github.com/ozgur-soft/deprem.io/environment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (model *Iletisim) Ara(ctx context.Context, search bson.D, skip int64, limit int64) (list []Iletisim) {
	cachekey := fmt.Sprintf("%v_%v_%v", IletisimCollection, skip, limit)
	if cache.Get(ctx, cachekey) != nil {
		data := cache.Get(ctx, cachekey)
		reader := bytes.NewReader(data)
		decoder := json.NewDecoder(reader)
		decoder.Decode(&list)
		return list
	}
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", environment.MongoUser, environment.MongoPass, environment.MongoHost, environment.MongoPort, environment.MongoAuth)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)
	collection := cli.Database(environment.MongoDb).Collection(IletisimCollection)
	cursor, err := collection.Find(ctx, search, options.Find().SetSkip(skip).SetLimit(limit))
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
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", environment.MongoUser, environment.MongoPass, environment.MongoHost, environment.MongoPort, environment.MongoAuth)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)
	collection := cli.Database(environment.MongoDb).Collection(IletisimCollection)
	insert, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", insert.InsertedID)
}
