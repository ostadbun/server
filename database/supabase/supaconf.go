package supabase

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Config() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return conn
}
