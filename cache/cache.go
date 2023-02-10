package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ozgur-soft/deprem.io/environment"
)

func Get(ctx context.Context, key string) []byte {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Password: environment.RedisPass, DB: 0})
	defer conn.Close()
	data, err := conn.Get(ctx, key).Bytes()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func Exists(ctx context.Context, key string) (exists bool) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Password: environment.RedisPass, DB: 0})
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
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Password: environment.RedisPass, DB: 0})
	defer conn.Close()
	if err := conn.Del(ctx, key).Err(); err != nil {
		log.Fatal(err)
	}
}

func Set(ctx context.Context, key string, val interface{}, exp int) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Password: environment.RedisPass, DB: 0})
	defer conn.Close()
	if err := conn.Set(ctx, key, val, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Expire(ctx context.Context, key string, exp int) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Password: environment.RedisPass, DB: 0})
	defer conn.Close()

	if err := conn.Expire(ctx, key, time.Duration(exp)*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
}

func Flush(ctx context.Context) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: environment.RedisHost + ":" + environment.RedisPort, Password: environment.RedisPass, DB: 0})
	defer conn.Close()
	if err := conn.FlushDB(ctx).Err(); err != nil {
		log.Fatal(err)
	}
}
