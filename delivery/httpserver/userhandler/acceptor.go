package userhandler

import (
	"encoding/json"
	"fmt"
	"os"
	"ostadbun/entity"
	"ostadbun/param/userparam"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) acceptor(c *fiber.Ctx) error {

	//1) on priority first I give state for remove user state from redis
	state := c.Query("state")

	defer h.userSvc.RemoveState(c.Context(), state)

	//2) now check the error query is empty or not
	errorURL := c.Query("error")
	if errorURL != "" {
		return c.Redirect(os.Getenv("WEB_CLIENT") + "/acceptor?error=" + errorURL)
	}

	ProviderName := c.Params("provider")
	Code := c.Query("code")

	userAgentData, client, err := h.userSvc.CheckIntiRedis(c.Context(), state)

	fmt.Println(Code, state, "this:", client, "check it out")
	fmt.Println(string(userAgentData), err, "magmawei")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	switch ProviderName {

	case "google":
		claim, errC := h.userSvc.AcceptGoogleOauth(Code)

		var userData userparam.Google
		if errC != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("656634134234734576 - %s", errC))
		}

		if err := json.Unmarshal(claim, &userData); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("12312412312222 - %s", err))
		}

		data := entity.User{
			Email: userData.Email,
			Name:  userData.Name,
		}

		code, name, err := h.userSvc.Login(data, userAgentData)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		if client == "cli" {
			return c.Redirect(cliRedirectString(code))
		} else {
			cookeSetter(c, code, name)
		}

	case "github":
		claim, emails, errC := h.userSvc.AcceptGithubOauth(Code)
		var userData userparam.Github

		if errC != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("656634734576 - %s", errC))
		}

		if err := json.Unmarshal(claim, &userData); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		var userEmailsData userparam.GithubEmail
		if err := json.Unmarshal(emails, &userEmailsData); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		for _, email := range userEmailsData {
			if email.Primary == true {
				userData.Email = userEmailsData[0].Email
			}
		}

		if len(userData.Email) < 1 {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("76485384573845 - %s", errC))
		}

		data := entity.User{
			Email: userData.Email,
			Name:  userData.Name,
		}

		code, name, err := h.userSvc.Login(data, userAgentData)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("65asfh872r - %s", errC))
		}

		if client == "cli" {
			return c.Redirect(cliRedirectString(code))
		} else {
			cookeSetter(c, code, name)
		}

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Redirect(os.Getenv("WEB_CLIENT"))
}

func cookeSetter(c *fiber.Ctx, code string, username string) {

	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("COOKIE_TOKEN"),
		Value:    code,
		Path:     "/",
		Domain:   ".ostadbun.tech",
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		//TODO make true on production https
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
	})

	c.Cookie(&fiber.Cookie{
		Name:    os.Getenv("COOKIE_NAME"),
		Value:   username,
		Domain:  ".ostadbun.tech",
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	})

}

func cliRedirectString(c string) string {

	return os.Getenv("CLI_CLIENT") + "?code=" + c
}
