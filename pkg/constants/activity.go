package Activityconstants

type ActivityTriggersName string

const (
	Trigger_RegisterAccount ActivityTriggersName = "register-account"
	Trigger_MakeAdmin       ActivityTriggersName = "up-to-admin"
	Trigger_UnMakeAdmin     ActivityTriggersName = "down-from-admin"
)
