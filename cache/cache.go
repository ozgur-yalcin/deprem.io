package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func Get(ctx context.Context, key string) []byte {
	data, err := Redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil
	}
	return data
}

func Exists(ctx context.Context, key string) (exists bool) {
	data, err := Redis.Exists(ctx, key).Result()
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
	if err := Redis.Del(ctx, key).Err(); err != nil {
		log.Fatal(err)
	}
}

func Set(ctx context.Context, key string, val interface{}, exp int) {
	if err := Redis.Set(ctx, key, val, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Expire(ctx context.Context, key string, exp int) {
	if err := Redis.Expire(ctx, key, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Flush(ctx context.Context) {
	if err := Redis.FlushDB(ctx).Err(); err != nil {
		log.Fatal(err)
	}
}
