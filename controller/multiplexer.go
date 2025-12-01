package controller

import (
	"ostadbun/controller/auth"
	"ostadbun/database"

	"github.com/gofiber/fiber/v2"
)

func Config(app *fiber.App, db database.Database) {

	auth.Config(app.Group("ali"), db.Conn)

}
