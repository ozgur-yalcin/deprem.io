package cache

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const (
	redis_host = "localhost:6379"
	redis_pass = ""
	redis_db   = 0
)

func Get(ctx context.Context, key string) []byte {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	data, err := conn.Get(ctx, key).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

func Exists(ctx context.Context, key string) (exists bool) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	data, err := conn.Exists(ctx, key).Result()
	if err != nil {
		log.Println(err)
		exists = false
	}
	if data > 0 {
		exists = true
	} else {
		exists = false
	}
	return exists
}

func Delete(ctx context.Context, key string) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	err := conn.Del(ctx, key).Err()
	if err != nil {
		log.Println(err)
	}
}

func Set(ctx context.Context, key string, val interface{}, exp int) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	err := conn.Set(ctx, key, val, time.Duration(exp)*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
}

func Expire(ctx context.Context, key string, exp int) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	err := conn.Expire(ctx, key, time.Duration(exp)*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
}

func Flush(ctx context.Context) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	err := conn.FlushDB(ctx).Err()
	if err != nil {
		log.Println(err)
	}
}
