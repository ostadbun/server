package userhandler

import (
	"ostadbun/service/userservice"
)

type Handler struct {
	authSvc userservice.Auth
}

func New(authSvc userservice.Auth) Handler {
	return Handler{
		authSvc: authSvc,
	}
}
