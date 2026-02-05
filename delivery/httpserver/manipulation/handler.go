package manipulation

import (
	manipulation "ostadbun/service/manipulationService"
	"ostadbun/service/userservice"
)

type Handler struct {
	manipulSVC manipulation.Manipulation
	usersvc    userservice.User
}

func New(manipulSVC manipulation.Manipulation, usersvc userservice.User) Handler {
	return Handler{
		usersvc:    usersvc,
		manipulSVC: manipulSVC,
	}
}
