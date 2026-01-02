package userhandler

import (
	"database/sql"
	"ostadbun/service/activityService"
	"ostadbun/service/userservice"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	authSvc     userservice.User
	activitySvc activityService.Activity
	redis       *redis.Client
	db          *sql.DB
}

func New(authSvc userservice.User, activitySvc activityService.Activity, redis *redis.Client) Handler {
	return Handler{
		authSvc:     authSvc,
		redis:       redis,
		activitySvc: activitySvc,
	}
}
