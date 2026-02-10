package redisAdaptor

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func New() *redis.Client {

	dbNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		dbNumber = 0
	}

	ProNumber, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))

	if err != nil {
		ProNumber = 2
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
		Protocol: ProNumber,
	})

	return rdb
}

//REDIS_ADDR=localhost:6379
//#REDIS_ADDR=redis-17273.c311.eu-central-1-1.ec2.cloud.redislabs.com:17273
//
//REDIS_USERNAME=
//#REDIS_USERNAME=default
//
//REDIS_PASSWORD=
//#REDIS_PASSWORD=DeCiC4nD9zCgNscpCd40SlkHRTj0WufT
//
//REDIS_DB=0
