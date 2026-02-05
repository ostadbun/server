package redisOauth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisOauth struct {
	redis *redis.Client
}

func New(client *redis.Client) *RedisOauth {
	return &RedisOauth{redis: client}
}

func (o *RedisOauth) Set(ctx context.Context, userdata []byte) (string, error) {

	newSession := uuid.New().String()

	key := fmt.Sprintf("osbn-o-auth-state:%s", newSession)

	if err := o.redis.Set(ctx, key, userdata, time.Minute*6).Err(); err != nil {
		fmt.Println(err)

		return "", err
	}

	return key, nil
}
