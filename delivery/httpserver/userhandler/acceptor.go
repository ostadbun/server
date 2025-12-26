package userhandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) acceptor(c *fiber.Ctx) error {

	prm := c.Params("id")
	qry := c.Query("code")

	var (
		claim any
		err   error
	)

	if prm == "google" {
		claim, err = h.authSvc.AcceptGoogleOauth(qry)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	} else if prm == "github" {
		claim, err = h.authSvc.AcceptGithubOauth(qry)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	} else {

		return c.Status(fiber.StatusInternalServerError).SendString("not found")

	}

	return c.JSON(claim)
}
