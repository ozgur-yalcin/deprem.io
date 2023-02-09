package cache

import (
	"log"
	"net/http"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const (
	redis_host = "localhost:6379"
	redis_pass = ""
	redis_db   = 0
)

func Get(req *http.Request, key string) []byte {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	data, err := conn.Get(req.Context(), key).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

func Exists(req *http.Request, key string) (exists bool) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	data, err := conn.Exists(req.Context(), key).Result()
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

func Del(req *http.Request, key string) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	err := conn.Del(req.Context(), key).Err()
	if err != nil {
		log.Println(err)
	}
}

func Set(req *http.Request, key string, val interface{}, exp int) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	err := conn.Set(req.Context(), key, val, time.Duration(exp)*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
}

func Exp(req *http.Request, key string, exp int) {
	conn := redis.NewClient(&redis.Options{Network: "tcp", Addr: redis_host, Password: redis_pass, DB: redis_db})
	defer conn.Close()
	err := conn.Expire(req.Context(), key, time.Duration(exp)*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
}
