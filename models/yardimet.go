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
	Ilce         string             `json:"ilce,omitempty" bson:"ilce,omitempty"`
	HedefSehir   string             `json:"hedefSehir,omitempty" bson:"hedefSehir,omitempty"`
	Aciklama     string             `json:"aciklama,omitempty" bson:"aciklama,omitempty"`
	Fields       any                `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	YardimDurumu string             `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	IPv4         string             `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt    primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (model *Yardimet) Search(ctx context.Context, search bson.D, skip int64, limit int64) (list []Yardimet) {
	if key, err := json.Marshal(search); err == nil {
		cachekey := cache.Key(YardimetCollection, fmt.Sprintf("%v_%v_%v", string(key), skip, limit))
		if cache.Get(ctx, cachekey) != nil {
			data := cache.Get(ctx, cachekey)
			reader := bytes.NewReader(data)
			decoder := json.NewDecoder(reader)
			decoder.Decode(&list)
			return list
		}
	}
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", environment.MongoUser, environment.MongoPass, environment.MongoHost, environment.MongoPort, environment.MongoAuthDb)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)
	col := cli.Database(environment.MongoDb).Collection(YardimetCollection)
	cur, err := col.Find(ctx, search, options.Find().SetSkip(skip).SetLimit(limit).SetAllowDiskUse(true))
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	if cur.Err() == nil {
		cur.All(ctx, &list)
	}
	return list
}

func (model *Yardimet) Add(ctx context.Context, data Yardimet) string {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", environment.MongoUser, environment.MongoPass, environment.MongoHost, environment.MongoPort, environment.MongoAuthDb)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)
	col := cli.Database(environment.MongoDb).Collection(YardimetCollection)
	add, err := col.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", add.InsertedID)
}
