package manipulation

import (
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/manipulation")

	//test you have basic permission
	userGroup.Get("/permission", middlewares.Auth(h.usersvc), middlewares.ManipulationPermission(h.manipulSVC), h.BasicPermission)

	userGroup.Post("/lesson", middlewares.Auth(h.usersvc), middlewares.ManipulationPermission(h.manipulSVC), h.AddLesson)

}
