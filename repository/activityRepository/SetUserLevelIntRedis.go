package activityRepository

import (
	"context"
	"fmt"
	"time"
)

func (d DB) SetUserLevelIntRedis(ctx context.Context, userid int, level int) error {

	if level < 1 {
		return fmt.Errorf("level must be greater than or equal to 1")
	}

	key := fmt.Sprintf("user-level:%d", userid)

	return d.redis.Set(ctx, key, level, time.Hour).Err()

}
