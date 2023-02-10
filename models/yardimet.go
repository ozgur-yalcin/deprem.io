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

const YardimetCollection = "yardimet"

type Yardimet struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	YardimTipi   string             `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"`
	AdSoyad      string             `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon      string             `json:"telefon,omitempty" bson:"telefon,omitempty"`
	Sehir        string             `json:"sehir,omitempty" bson:"sehir,omitempty"`
	Ilce         string             `json:"ilce,omitempty" bson:"ilce,omitempty"` // TODO: ilçe geçici required false yapıldı
	HedefSehir   string             `json:"hedefSehir,omitempty" bson:"hedefSehir,omitempty"`
	Aciklama     string             `json:"aciklama,omitempty" bson:"aciklama,omitempty"`
	Fields       any                `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	YardimDurumu string             `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	IPv4         string             `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt    primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (model *Yardimet) Ara(ctx context.Context, search bson.D, skip int64, limit int64) (list []Yardimet) {
	key, err := json.Marshal(search)
	if err != nil {
		log.Fatal(err)
	}
	cachekey := cache.Key(YardimetCollection, string(key)+fmt.Sprintf("%v", skip)+fmt.Sprintf("%v", limit))
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
	collection := cli.Database(environment.MongoDb).Collection(YardimetCollection)
	cursor, err := collection.Find(ctx, search, options.Find().SetSkip(skip).SetLimit(limit))
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

func (model *Yardimet) Ekle(ctx context.Context, data Yardimet) string {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", environment.MongoUser, environment.MongoPass, environment.MongoHost, environment.MongoPort, environment.MongoAuth)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)
	collection := cli.Database(environment.MongoDb).Collection(YardimetCollection)
	insert, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", insert.InsertedID)
}
