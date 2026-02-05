package middlewares

import (
	"os"
	"ostadbun/pkg/httpstorage"
	manipulation "ostadbun/service/manipulationService"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ManipulationPermission(m manipulation.Manipulation) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {

		userId, Err := httpstorage.Get(c, "user_id").Number()

		if Err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": Err.Error()})
		}

		level, errG := m.GetUserLevel(c.Context(), userId)

		if errG != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errG.Error())
		}

		scre := os.Getenv("BASICPERMISSIONSCORE")

		score, err := strconv.Atoi(scre)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "score not a valid number"})
		}

		if level < score {
			return c.Status(fiber.StatusForbidden).SendString("you can not manipulate")
		}

		return c.Next()

	}
}
