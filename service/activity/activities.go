package activityService

import "ostadbun/service/activity/activityconstants"

func (a Activity) Trigger(userid int, activity activityconstants.ActivityTriggersName) error {

	return a.repo.TriggerSetter(userid, activity)

}
