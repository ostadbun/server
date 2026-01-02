package activityRepository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func (d DB) GetUserLevel(ctx context.Context, userid int) (int, error) {
	key := fmt.Sprintf("user-level:%d", userid)

	res, err := d.redis.Get(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return -1, nil
	}

	if err != nil {
		return -1, err
	}

	level, err := strconv.Atoi(res)
	if err != nil {

		return -1, err
	}

	return level, nil
}
