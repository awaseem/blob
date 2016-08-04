package constants

import (
	"os"
	"strconv"
)

// Default value
const portC string = "3232"

const redisHostC string = "localhost"
const redisPortC string = "6379"
const redisPasswordC string = ""
const redisDBC int = 0

// redis enviroment variables
const redisHostEnv = "REDIS_HOST"
const redisPortEnv = "REDIS_PORT"
const redisPasswordEnv = "REDIS_PASSWORD"
const redisDBEnv = "REDIS_DB"

// port enviroment variables
const portEnv = "PORT"

// GetRedisConfig get config variables for redis
func GetRedisConfig() (string, string, string, int) {
	redisHost := os.Getenv(redisHostEnv)
	redisPort := os.Getenv(redisPortEnv)
	redisPassword := os.Getenv(redisPasswordEnv)
	redisDB, err := strconv.Atoi(os.Getenv(redisDBEnv))
	if err != nil {
		redisDB = 0
	}
	if redisHost == "" && redisPort == "" && redisPassword == "" {
		redisHost = redisHostC
		redisPort = redisPortC
		redisPassword = redisPasswordC
		redisDB = redisDBC
	}
	return redisHost, redisPort, redisPassword, redisDB
}

// GetPortConfig get config variable for server
func GetPortConfig() string {
	port := os.Getenv(portEnv)
	if port == "" {
		port = portC
	}
	return port
}
