package database

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ozgur-soft/deprem.io/cache"
	"github.com/ozgur-soft/deprem.io/environment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", environment.MongoUser, environment.MongoPass, environment.MongoHost, environment.MongoPort, environment.MongoAuthDb)
	cli, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	client = cli
}

func Search(ctx context.Context, collection string, search bson.D, skip int64, limit int64) (list []any) {
	if key, err := json.Marshal(search); err == nil {
		cachekey := cache.Key(collection, fmt.Sprintf("%v_%v_%v", string(key), skip, limit))
		if cache.Get(ctx, cachekey) != nil {
			data := cache.Get(ctx, cachekey)
			reader := bytes.NewReader(data)
			decoder := json.NewDecoder(reader)
			decoder.Decode(&list)
			return list
		}
	}
	cur, err := client.Database(environment.MongoDb).Collection(collection).Find(ctx, search, options.Find().SetSkip(skip).SetLimit(limit).SetAllowDiskUse(true))
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	if cur.Err() == nil {
		cur.All(ctx, &list)
	}
	return list
}

func Add(ctx context.Context, collection string, data any) string {
	add, err := client.Database(environment.MongoDb).Collection(collection).InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", add.InsertedID)
}
