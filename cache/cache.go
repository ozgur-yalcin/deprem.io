package cache

import (
	"context"
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func Get(ctx context.Context, key string) []byte {
	uri, err := redis.ParseURL(os.Getenv("REDISURL"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	conn := redis.NewClient(uri)
	defer conn.Close()
	data, err := conn.Get(ctx, key).Bytes()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func Exists(ctx context.Context, key string) (exists bool) {
	uri, err := redis.ParseURL(os.Getenv("REDISURL"))
	if err != nil {
		log.Fatal(err)
	}
	conn := redis.NewClient(uri)
	defer conn.Close()
	data, err := conn.Exists(ctx, key).Result()
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
	uri, err := redis.ParseURL(os.Getenv("REDISURL"))
	if err != nil {
		log.Fatal(err)
	}
	conn := redis.NewClient(uri)
	defer conn.Close()
	if err := conn.Del(ctx, key).Err(); err != nil {
		log.Fatal(err)
	}
}

func Set(ctx context.Context, key string, val interface{}, exp int) {
	uri, err := redis.ParseURL(os.Getenv("REDISURL"))
	if err != nil {
		log.Fatal(err)
	}
	conn := redis.NewClient(uri)
	defer conn.Close()
	if err := conn.Set(ctx, key, val, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Expire(ctx context.Context, key string, exp int) {
	uri, err := redis.ParseURL(os.Getenv("REDISURL"))
	if err != nil {
		log.Fatal(err)
	}
	conn := redis.NewClient(uri)
	defer conn.Close()

	if err := conn.Expire(ctx, key, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Flush(ctx context.Context) {
	uri, err := redis.ParseURL(os.Getenv("REDISURL"))
	if err != nil {
		log.Fatal(err)
	}
	conn := redis.NewClient(uri)
	defer conn.Close()
	if err := conn.FlushDB(ctx).Err(); err != nil {
		log.Fatal(err)
	}
}
