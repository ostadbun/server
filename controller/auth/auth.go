package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type auth struct {
	Access_token   string `json:"access_token"`
	Expires_at     string `json:"expires_at"`
	Expires_in     string `json:"expires_in"`
	Refresh_token  string `json:"refresh_token"`
	Provider_token string `json:"provider_token"`
}

func (c Controller) Auth(f *fiber.Ctx) error {

	var data auth
	err := f.BodyParser(&data)

	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("🐥")
	}

	return f.SendString(fmt.Sprintf("Hello, World! %s", data.Access_token))
}
