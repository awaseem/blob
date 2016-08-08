package constants

import (
	"os"
	"strconv"
)

// Default value
const portC string = "3232"

// redis constants
const redisHostC string = "localhost"
const redisPortC string = "6379"
const redisPasswordC string = ""
const redisDBC int = 0

// jwtea constants
const jwteaURLC string = "https://token.aliwaseem.com"

// redis enviroment variables
const redisHostEnv = "REDIS_HOST"
const redisPortEnv = "REDIS_PORT"
const redisPasswordEnv = "REDIS_PASSWORD"
const redisDBEnv = "REDIS_DB"

// port enviroment variables
const portEnv = "PORT"

// jwtea enviroment variables
const jwteaURLEnv = "JWTEA_URL"

// GetRedisConfig get config variables for redis
func GetRedisConfig() (string, string, string, int) {
	redisHost := os.Getenv(redisHostEnv)
	redisPort := os.Getenv(redisPortEnv)
	redisPassword := os.Getenv(redisPasswordEnv)
	redisDB, err := strconv.Atoi(os.Getenv(redisDBEnv))
	if err != nil {
		redisDB = redisDBC
	}
	if redisHost == "" {
		redisHost = redisHostC
	}
	if redisPort == "" {
		redisPort = redisPortC
	}
	if redisPassword == "" {
		redisPassword = redisPasswordC
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

// GetJWTeaConfig get config variable for jwtea
func GetJWTeaConfig() string {
	url := os.Getenv(jwteaURLEnv)
	if url == "" {
		url = jwteaURLC
	}
	return url
}
