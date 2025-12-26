package userhandler

import (
	"ostadbun/service/userservice"
)

type Handler struct {
	authSvc userservice.User
}

func New(authSvc userservice.User) Handler {
	return Handler{
		authSvc: authSvc,
	}
}
