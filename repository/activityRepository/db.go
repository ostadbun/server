package activityRepository

import (
	"ostadbun/database"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	conn  *database.PostgresDB
	redis *redis.Client
}

func New(conn *database.PostgresDB, redis *redis.Client) *DB {
	return &DB{
		conn:  conn,
		redis: redis,
	}
}
