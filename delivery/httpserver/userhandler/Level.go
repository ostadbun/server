package userhandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Level(c *fiber.Ctx) error {

	UserID := c.Params("userid")

	number, errE := h.authSvc.LevelCalculator(UserID)

	if errE != nil {
		return c.SendString(errE.Error())
	}
	return c.SendString(fmt.Sprintf("count is : %d", number))
}
