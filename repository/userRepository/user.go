package userRepository

import (
	"database/sql"
	"fmt"
	"ostadbun/entity"
	"ostadbun/pkg/hash"
)

type AuthRepo struct {
	db *sql.DB
}

func Make(db *sql.DB) AuthRepo {
	return AuthRepo{
		db: db,
	}
}

func (a AuthRepo) ExistingCheck(email string) (int, bool) {
	var id int

	email = hash.Hasher(email)

	err := a.db.QueryRow(
		`SELECT id FROM users WHERE email = $1 LIMIT 1`,
		email,
	).Scan(&id)

	fmt.Println(email, id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, false

	}

	return id, id > 0
}

func (a AuthRepo) RegisterUser(user entity.User) (int, error) {

	var id int

	err := a.db.QueryRow(`
		INSERT INTO users (email, name)
		VALUES ($1, $2)
		RETURNING id
	`,
		user.Email_Hashe(),
		user.Name,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
