package activityService

import (
	"context"
	"ostadbun/pkg/constants"
)

func (a Activity) Trigger(ctx context.Context, userid int, activity Activityconstants.ActivityTriggersName) error {

	defer func() {
		errC := a.UpdateRedisCash(ctx, userid)
		if errC != nil {
			//TODO log here
		}
	}()

	return a.repo.TriggerSetter(userid, activity)

}
