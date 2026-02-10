package academic

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/academic")

	userGroup.Get("/university", h.University)

	userGroup.Get("", h.Academics)

}
