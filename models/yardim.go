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
	bson "go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

const YardimCollection = "yardim"

type Yardim struct {
	Id              any       `json:"_id,omitempty" bson:"_id,omitempty"`
	YardimTipi      string    `json:"yardimTipi,omitempty" bson:"yardimTipi,omitempty"` // Gıda, İlaç, Enkaz, Isınma, Kayıp
	AdSoyad         string    `json:"adSoyad,omitempty" bson:"adSoyad,omitempty"`
	Telefon         string    `json:"telefon,omitempty" bson:"telefon,omitempty"`
	YedekTelefonlar string    `json:"yedekTelefonlar,omitempty" bson:"yedekTelefonlar,omitempty"`
	Email           string    `json:"email,omitempty" bson:"email,omitempty"`
	Adres           string    `json:"adres,omitempty" bson:"adres,omitempty"`
	AdresTarifi     string    `json:"adresTarifi,omitempty" bson:"adresTarifi,omitempty"`
	AcilDurum       string    `json:"acilDurum,omitempty" bson:"acilDurum,omitempty"`
	YardimDurumu    string    `json:"yardimDurumu,omitempty" bson:"yardimDurumu,omitempty"`
	KisiSayisi      string    `json:"kisiSayisi,omitempty" bson:"kisiSayisi,omitempty"`
	FizikiDurum     string    `json:"fizikiDurum,omitempty" bson:"fizikiDurum,omitempty"`
	GoogleMapLink   string    `json:"googleMapLink,omitempty" bson:"googleMapLink,omitempty"`
	TweetLink       string    `json:"tweetLink,omitempty" bson:"tweetLink,omitempty"`
	Fields          any       `json:"fields,omitempty" bson:"fields,omitempty"` // Tüm alternatif kullanımlar için buraya json yollayın
	IPv4            string    `json:"ip,omitempty" bson:"ip,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (model *Yardim) Ara(ctx context.Context, search bson.D, skip int64, limit int64) (list []Yardim) {
	cachekey := fmt.Sprintf("%v_%v_%v", YardimCollection, skip, limit)
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
	collection := client.Database(os.Getenv("MONGODB")).Collection(YardimCollection)
	cursor, err := collection.Find(ctx, search, mongooptions.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	if cursor.Err() == nil {
		for cursor.Next(ctx) {
			var data Yardim
			cursor.Decode(&data)
			list = append(list, data)
		}
	}
	return list
}

func (model *Yardim) Ekle(ctx context.Context, data Yardim) string {
	client, err := mongodb.Connect(ctx, mongooptions.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database(os.Getenv("MONGODB")).Collection(YardimCollection)
	insert, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", insert.InsertedID)
}
