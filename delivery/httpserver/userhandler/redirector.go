package userhandler

import (
	"encoding/json"
	"fmt"
	"ostadbun/delivery/httpserver/pkg/useragent"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) redirector(c *fiber.Ctx) error {

	provider := c.Query("provider", "")

	client := c.Query("client", "web")

	a := useragent.ReadDeviceInfo(c, client)
	data, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}

	url, err := h.authSvc.RedirectUrlGenerator(provider, client, data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("invalid provider called: %s", provider))
	}

	return c.RedirectBack(url)
}

func (h Handler) t(c *fiber.Ctx) error {

	return c.SendString("sw")
}
