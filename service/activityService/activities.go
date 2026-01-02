package activityService

import "ostadbun/pkg/constants"

func (a Activity) Trigger(userid int, activity Activityconstants.ActivityTriggersName) error {

	defer func() {
		errC := a.UpdateRedisCash(userid)
		if errC != nil {
			//TODO log here
		}
	}()

	return a.repo.TriggerSetter(userid, activity)

}
