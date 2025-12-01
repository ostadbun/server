package authReposipory

import "github.com/jackc/pgx/v5"

type AuthRepo struct {
	db *pgx.Conn
}

func Make(db *pgx.Conn) AuthRepo {
	return AuthRepo{
		db: db,
	}
}

func (a AuthRepo) createUser(username string, password string) error {
	panic("implement me")
}
