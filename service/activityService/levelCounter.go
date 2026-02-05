package activityService

import (
	"context"
	"errors"
	"fmt"
)

//1) check on Redis is Exist?

//2) if exist pass them

//3) if not exist calculate and set on redis then pass that

func (a Activity) LevelCalculator(ctx context.Context, userid int) (int, error) {

	LevelCounted, errRedis := a.redis.Check(ctx, userid)
	//TODO log Redis Error
	if errors.Is(errRedis, nil) && LevelCounted > -1 {
		fmt.Println("LevelCounted rds", LevelCounted)
		return LevelCounted, nil
	}

	MainlevelCounted, ErrPsg := a.repo.MainStoreCalculateAndFetch(userid)
	//TODO log postgres Error
	if errors.Is(ErrPsg, nil) && MainlevelCounted > -1 {
		fmt.Println("LevelCounted psql", MainlevelCounted)
		SetNewToRedis := a.redis.Set(ctx, userid, MainlevelCounted)
		if SetNewToRedis != nil {
			//TODO log this
			fmt.Println("SetNewToRedis", SetNewToRedis)
		}
		return MainlevelCounted, nil
	}

	if ErrPsg != nil {
		return -1, ErrPsg
	}

	return -1, fmt.Errorf("unhandled user level calculation")

}
