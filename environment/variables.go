package environment

import "os"

var (
	RedisHost = os.Getenv("REDIS_HOST")
	RedisUser = os.Getenv("REDIS_USER")
	RedisPass = os.Getenv("REDIS_PASS")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisDb   = os.Getenv("REDIS_DB")
)

var (
	MongoHost = os.Getenv("MONGO_HOST")
	MongoUser = os.Getenv("MONGO_USER")
	MongoPass = os.Getenv("MONGO_PASS")
	MongoPort = os.Getenv("MONGO_PORT")
	MongoDb   = os.Getenv("MONGO_DB")
)
