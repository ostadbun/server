package manipulation

import (
	manipulationParam "ostadbun/param/manipulation"
	"ostadbun/pkg/httpstorage"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) AddLesson(c *fiber.Ctx) error {

	var lesson manipulationParam.PendingLesson

	err := c.BodyParser(&lesson)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userId, errG := httpstorage.Get(c, "user_id").Number()

	if errG != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errG.Error(),
		})
	}

	errA := h.manipulSVC.AddLesson(lesson, userId)

	if errA != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errA,
		})
	}

	return c.JSON("add success")

}
