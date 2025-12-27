package userhandler

import (
	"encoding/json"
	"fmt"
	"ostadbun/delivery/httpserver/pkg/useragent"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) redirector(c *fiber.Ctx) error {

	provder := c.Query("provider", "")

	data, err := json.Marshal(useragent.ReadDeviceInfo(c))

	if err != nil {
		fmt.Println(err)
	}

	url, err := h.authSvc.RedirectUrlGenerator(provder, data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid provider called: %s", provder))
	}

	return c.RedirectBack(url)
}

func (h Handler) t(c *fiber.Ctx) error {

	return c.SendString("sw")
}
