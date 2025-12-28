package redisAdaptor

import (
	"github.com/redis/go-redis/v9"
)



func New() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	return rdb
}
