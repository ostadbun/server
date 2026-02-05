package activityService

import (
	"context"
	"errors"
	"fmt"
)

func (a Activity) UpdateRedisCash(ctx context.Context, userid int) error {
	MainlevelCounted, ErrPsg := a.repo.MainStoreCalculateAndFetch(userid)
	//TODO log postgres Error
	if errors.Is(ErrPsg, nil) && MainlevelCounted > -1 {
		fmt.Println("LevelCounted psql", MainlevelCounted)
		SetNewToRedis := a.redis.Set(ctx, userid, MainlevelCounted)
		if SetNewToRedis != nil {
			//TODO log this
			fmt.Println("SetNewToRedis", SetNewToRedis)
		}
		return nil
	}

	return fmt.Errorf("unhandled user level calculation %s", ErrPsg.Error())
}
