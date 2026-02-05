package userhandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Level(c *fiber.Ctx) error {

	usId := c.Params("userid")

	UserID, err := strconv.Atoi(usId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	number, errE := h.activitySvc.LevelCalculator(c.Context(), UserID)

	if errE != nil {
		return c.SendString(errE.Error())
	}
	return c.SendString(fmt.Sprintf("count is : %d", number))
}
