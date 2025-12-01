package authService

import authReposipory "ostadbun/repository/auth"

type Auth struct {
	repo authReposipory.AuthRepo
}

func Config(r authReposipory.AuthRepo) Auth {
	return Auth{
		repo: r,
	}
}

func (a Auth) CreateUser(username string, password string) error { panic("") }
