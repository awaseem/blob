package redis

import (
	"blob/constants"

	redis "gopkg.in/redis.v4"
)

var globalClient *redis.Client

// CreateClient create a redis client, make sure this is called in init or main before using the client
func CreateClient() {
	redisHost, redisPort, redisPass, redisDB := constants.GetRedisConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       redisDB,
	})
	globalClient = client
}
