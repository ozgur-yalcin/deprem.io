package environment

import "os"

var (
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisUser = os.Getenv("REDIS_USER")
	RedisPass = os.Getenv("REDIS_PASS")
	RedisDb   = os.Getenv("REDIS_DB")
)

var (
	MongoHost = os.Getenv("MONGO_HOST")
	MongoPort = os.Getenv("MONGO_PORT")
	MongoUser = os.Getenv("MONGO_USER")
	MongoPass = os.Getenv("MONGO_PASS")
	MongoAuth = os.Getenv("MONGO_AUTH")
	MongoDb   = os.Getenv("MONGO_DB")
)
