package config

import "os"

func GetRedisUrl() string {
	var defaultUrl = "localhost:6379"
	var redisUrl = os.Getenv("REDIS_URL")

	// Use localhost and well known port if env is not set
	if(redisUrl == "") {
		redisUrl = defaultUrl
	}

	return redisUrl
}