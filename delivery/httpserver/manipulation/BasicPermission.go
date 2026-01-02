package manipulation

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) BasicPermission(c *fiber.Ctx) error {

	return c.SendString("you can manipulate !  \uF179")

}
