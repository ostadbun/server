package authReposipory

import (
	"database/sql"
)

type AuthRepo struct {
	db *sql.DB
}

func Make(db *sql.DB) AuthRepo {
	return AuthRepo{
		db: db,
	}
}

func (a AuthRepo) CreateUser(username string, password string) error {
	panic("implement me")
}
