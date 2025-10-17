package server

import (
	envConf "ostadbun/config"
	"ostadbun/database"

	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
	db  database.Database
	env *envConf.Env
}

func NewFiberServer(env *envConf.Env, db database.Database) Server {

	fiberApp := fiber.New()

	return &fiberServer{
		app: fiberApp,
		db:  db,
		env: env,
	}
}

func (s *fiberServer) Start() {
	s.app.Get("health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	s.app.Listen(":3000")
}

func (s *fiberServer) RegisterRoutes() {

	panic("")
}
