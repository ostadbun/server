package userhandler

import (
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/user")

	userGroup.Get("/oauth", h.redirector)

	userGroup.Get("/oauth/callback/:provider", h.acceptor)

	userGroup.Get("/ow", middlewares.Auth(h.redis), test)

}

func test(c *fiber.Ctx) error {

	return c.SendString("you access here yeay")
}
