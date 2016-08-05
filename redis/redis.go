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

// Set set a key with the a value if that key does not exists
func Set(key string, value []byte) (bool, error) {
	return globalClient.SetNX(key, value, 0).Result()
}

// Get get a value based on the key
func Get(key string) (string, error) {
	return globalClient.Get(key).Result()
}
