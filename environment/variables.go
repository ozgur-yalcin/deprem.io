package environment

import "os"

var (
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisUser = os.Getenv("REDIS_USER")
	RedisPass = os.Getenv("REDIS_PASS")
)

var (
	MongoHost   = os.Getenv("MONGO_HOST")
	MongoPort   = os.Getenv("MONGO_PORT")
	MongoUser   = os.Getenv("MONGO_USER")
	MongoPass   = os.Getenv("MONGO_PASS")
	MongoAuthDb = os.Getenv("MONGO_AUTH_DB")
	MongoDb     = os.Getenv("MONGO_DB")
)
