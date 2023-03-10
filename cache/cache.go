package cache

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"

	"github.com/ozgur-soft/deprem.io/environment"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Connect() {
	if environment.RedisUser != "" {
		client = redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Username: environment.RedisUser, Password: environment.RedisPass})
	} else {
		client = redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Password: environment.RedisPass})
	}
}

func Key(prefix, key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	sum := hash.Sum(nil)
	cachekey := prefix + ":" + hex.EncodeToString(sum)
	return cachekey
}

func Get(ctx context.Context, key string) []byte {
	data, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return nil
	}
	return data
}

func Exists(ctx context.Context, key string) (exists bool) {
	data, err := client.Exists(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	if data > 0 {
		exists = true
	} else {
		exists = false
	}
	return exists
}

func Delete(ctx context.Context, key string) {
	if err := client.Del(ctx, key).Err(); err != nil {
		log.Fatal(err)
	}
}

func Set(ctx context.Context, key string, val interface{}, exp int) {
	if err := client.Set(ctx, key, val, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Expire(ctx context.Context, key string, exp int) {
	if err := client.Expire(ctx, key, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Flush(ctx context.Context) {
	if err := client.FlushDB(ctx).Err(); err != nil {
		log.Fatal(err)
	}
}
