package database

import (
	"fmt"
	"log"
	"sync"

	envConf "ostadbun/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresDatabase struct {
	Db *sqlx.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresSql(env *envConf.Env) *postgresDatabase {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			env.Host,
			env.Username,
			env.Password,
			env.Database,
			env.Db_port,
		)

		db, err := sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Fatalf("❌ failed to connect database: %v", err)
		}

		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxIdleTime(5 * 60)

		log.Println("✅ connection successful")

		dbInstance = &postgresDatabase{Db: db}
	})

	return dbInstance
}

func (p *postgresDatabase) GetDB() *sqlx.DB {
	return p.Db
}
