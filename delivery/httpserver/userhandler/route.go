package userhandler

import "github.com/gofiber/fiber/v2"

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/user")

	userGroup.Get("/oauth", h.redirector)

	userGroup.Get("/oauth/callback/:id", h.acceptor)

}
