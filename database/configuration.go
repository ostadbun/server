package database

import (
	"ostadbun/database/supabase"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
}

func Config() Database {
	return Database{
		supabase.Config(),
	}
}
