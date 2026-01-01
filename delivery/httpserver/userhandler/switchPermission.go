package userhandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) switchPermission(c *fiber.Ctx) error {

	usida := c.Params("userid")

	userId, errN := strconv.Atoi(usida)

	if errN != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("userID should be a number")
	}

	fmt.Println(c.Locals("user_id"))
	mstrid := c.Locals("user_id")

	masterID, ok := strconv.Atoi(mstrid.(string))

	fmt.Println(masterID, ok)
	if ok != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("masterID not found wdaij")
	}

	err := h.authSvc.SwitchPermission(userId, masterID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Render("user/switchPermission", fiber.Map{})
}
