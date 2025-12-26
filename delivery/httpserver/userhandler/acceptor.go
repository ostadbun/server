package userhandler

import (
	"encoding/json"
	"ostadbun/entity"
	"ostadbun/param/userparam"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) acceptor(c *fiber.Ctx) error {

	ProviderName := c.Params("provider")
	Code := c.Query("code")

	switch ProviderName {

	case "google":
		claim, errC := h.authSvc.AcceptGoogleOauth(Code)

		if errC != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("invalid provider 6346347364")
		}
		var userData userparam.Google

		if err := json.Unmarshal(claim, &userData); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		data := entity.User{
			Email: userData.Email,
			Name:  userData.Name,
		}

		code, err := h.authSvc.Login(data)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		// get login jwt
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": code,
		})

	case "github":
		claim, emails, errC := h.authSvc.AcceptGithubOauth(Code)
		if errC != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("invalid provider 623413234")
		}
		var userData userparam.Github

		var userEmailsData userparam.GithubEmail

		if err := json.Unmarshal(claim, &userData); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		if err := json.Unmarshal(emails, &userEmailsData); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		for _, email := range userEmailsData {
			if email.Primary == true {
				userData.Email = userEmailsData[0].Email
			}
		}

		if len(userData.Email) < 1 {
			return c.Status(fiber.StatusInternalServerError).SendString("user email is nil")
		}

		data := entity.User{
			Email: userData.Email,
			Name:  userData.Name,
		}

		code, err := h.authSvc.Login(data)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		// get login jwt
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": code,
		})
	}

	return c.SendString("end")
}
