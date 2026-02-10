package academicRepository

import (
	"ostadbun/database"
)

type DB struct {
	conn *database.PostgresDB
}

func New(conn *database.PostgresDB) *DB {
	return &DB{
		conn: conn,
	}
}
