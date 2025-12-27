package userhandler

import (
	"ostadbun/service/userservice"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	authSvc userservice.User
	redis   *redis.Client
}

func New(authSvc userservice.User, redis *redis.Client) Handler {
	return Handler{
		authSvc: authSvc,
		redis:   redis,
	}
}
