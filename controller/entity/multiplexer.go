package entity

import (
	entityrepository "ostadbun/repository/entity"
	entityService "ostadbun/service/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Controller struct {
	app fiber.Router
	Svc entityService.Entity
}

func Config(app fiber.Router, db *pgx.Conn, prefix ...string) {
	repo := entityrepository.Make(db)
	svc := entityService.Config(repo)

	m := Controller{Svc: svc, app: app}

	app.Post("/", m.Entity)

}
