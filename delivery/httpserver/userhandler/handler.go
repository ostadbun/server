package userhandler

import (
	"database/sql"
	"ostadbun/service/activityService"
	"ostadbun/service/userservice"
)

type Handler struct {
	userSvc     userservice.User
	activitySvc activityService.Activity
	db          *sql.DB
}

func New(userSvc userservice.User, activitySvc activityService.Activity) Handler {
	return Handler{
		userSvc:     userSvc,
		activitySvc: activitySvc,
	}
}
