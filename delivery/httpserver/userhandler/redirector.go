package userhandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) redirector(c *fiber.Ctx) error {

	provder := c.Query("provider", "")

	url, err := h.authSvc.RedirectUrlGenerator(provder)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid provider called: %s", provder))
	}

	fmt.Println(url)
	return c.RedirectBack(url)
}

func (h Handler) t(c *fiber.Ctx) error {

	return c.SendString("sw")
}
