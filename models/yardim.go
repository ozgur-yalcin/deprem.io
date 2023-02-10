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

const YardimCollection = "yardim"

type Yardim struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	YardimTipi      string             `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"` // Gıda, İlaç, Enkaz, Isınma, Kayıp
	AdSoyad         string             `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon         string             `json:"telefon,omitempty" bson:"telefon,omitempty"`
	YedekTelefonlar []string           `json:"yedekTelefonlar,omitempty" bson:"yedekTelefonlar,omitempty"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty"`
	Adres           string             `json:"adres,omitempty" bson:"adres,omitempty"`
	AdresTarifi     string             `json:"adresTarifi,omitempty" bson:"adresTarifi,omitempty"`
	AcilDurum       string             `json:"acilDurum,omitempty" bson:"acilDurum,omitempty"`
	YardimDurumu    string             `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	KisiSayisi      string             `json:"kisiSayisi,omitempty" bson:"kisiSayisi,omitempty"`
	FizikiDurum     string             `json:"fizikiDurum,omitempty" bson:"fizikiDurum,omitempty"`
	GoogleMapLink   string             `json:"googleMapLink,omitempty" bson:"googleMapLink,omitempty"`
	TweetLink       string             `json:"tweetLink,omitempty" bson:"tweetLink,omitempty"`
	Fields          any                `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	IPv4            string             `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt       primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt       primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (model *Yardim) Search(ctx context.Context, search bson.D, skip int64, limit int64) (list []Yardim) {
	if key, err := json.Marshal(search); err == nil {
		cachekey := cache.Key(YardimCollection, fmt.Sprintf("%v_%v_%v", string(key), skip, limit))
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
	col := cli.Database(environment.MongoDb).Collection(YardimCollection)
	cur, err := col.Find(ctx, search, options.Find().SetSkip(skip).SetLimit(limit).SetAllowDiskUse(true))
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	if cur.Err() == nil {
		for cur.Next(ctx) {
			data := Yardim{}
			cur.Decode(&data)
			list = append(list, data)
		}
	}
	return list
}

func (model *Yardim) Add(ctx context.Context, data Yardim) string {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", environment.MongoUser, environment.MongoPass, environment.MongoHost, environment.MongoPort, environment.MongoAuthDb)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Disconnect(ctx)
	col := cli.Database(environment.MongoDb).Collection(YardimCollection)
	add, err := col.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", add.InsertedID)
}
