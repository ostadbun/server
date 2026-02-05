package userhandler

import (
	"ostadbun/delivery/httpserver/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/user")

	userGroup.Get("/oauth", h.redirector)

	userGroup.Get("/oauth/callback/:provider", h.acceptor)

	userGroup.Get("/switch-permission/:userid", middlewares.Auth(h.userSvc), middlewares.IsAdmin(h.userSvc), h.switchPermission)

	userGroup.Get("/level/:userid", h.Level)

	userGroup.Get("/ow", middlewares.Auth(h.userSvc), middlewares.IsAdmin(h.userSvc), test)

}

func test(c *fiber.Ctx) error {

	return c.SendString("you access here yeay")
}
