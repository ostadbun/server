package activityService

import (
	"context"
	"ostadbun/repository/postgres/activityRepository"
)

type RedisActivity interface {
	Check(ctx context.Context, userid int) (int, error)
	Set(ctx context.Context, userid, level int) error
}

type Activity struct {
	repo  *activityRepository.DB
	redis RedisActivity
}

func New(repo *activityRepository.DB, redis RedisActivity) Activity {
	return Activity{
		repo:  repo,
		redis: redis,
	}
}
