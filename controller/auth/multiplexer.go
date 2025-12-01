package auth

import (
	authReposipory "ostadbun/repository/auth"
	authService "ostadbun/service/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Controller struct {
	app fiber.Router
	Svc authService.Auth
}

func Config(app fiber.Router, db *pgx.Conn, prefix ...string) {
	repo := authReposipory.Make(db)
	svcu := authService.Config(repo)

	m := Controller{Svc: svcu, app: app}

	app.Get("/s", m.Auth)

}
