package controller

import (
	"ostadbun/controller/auth"
	"ostadbun/controller/entity"
	"ostadbun/database"

	"github.com/gofiber/fiber/v2"
)

func Config(app *fiber.App, db database.Database) {

	auth.Config(app.Group("auth"), db.Conn)

	entity.Config(app.Group("entity"), db.Conn)

}
