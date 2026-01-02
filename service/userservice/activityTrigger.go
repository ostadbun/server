package userservice

import (
	"ostadbun/service/activity/activityconstants"
)

func (r User) ActivityTrigger(userID int, activity activityconstants.ActivityTriggersName) error {

	return r.activity.Trigger(userID, activity)

}
