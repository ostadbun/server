package userhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"ostadbun/entity"
	"ostadbun/param/userparam"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func (h Handler) acceptor(c *fiber.Ctx) error {

	ProviderName := c.Params("provider")
	Code := c.Query("code")

	state := c.Query("state")

	defer removeState(h.redis, state)

	userAgentData, client, err := checkIntoRedis(state, h.redis)
	fmt.Println(Code, state, "this:", client, "check it out")
	fmt.Println(string(userAgentData), err, "magmawei")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	switch ProviderName {

	case "google":
		claim, errC := h.authSvc.AcceptGoogleOauth(Code)
		var userData userparam.Google

		if errC != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("656634734576 - %s", errC))
		}

		if err := json.Unmarshal(claim, &userData); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("12312412312222 - %s", err))
		}

		data := entity.User{
			Email: userData.Email,
			Name:  userData.Name,
		}

		code, name, err := h.authSvc.Login(data, userAgentData)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		if client == "cli" {
			return c.Redirect(cliRedirectString(code))
		} else {
			cookeSetter(c, code, name)
		}

	case "github":
		claim, emails, errC := h.authSvc.AcceptGithubOauth(Code)
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

		code, name, err := h.authSvc.Login(data, userAgentData)

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

type DeviceInfo struct {
	// mobile | desktop
	Type string `json:"type"`

	// web | terminal
	Client string `json:"client"`

	// android | ios | windows | mac | linux
	OS string `json:"os"`
}

func checkIntoRedis(key string, rds *redis.Client) ([]byte, string, error) {

	rs, errR := rds.Get(context.Background(), key).Result()

	if errR != nil {
		return nil, "", errR
	}

	var data DeviceInfo

	if err := json.Unmarshal([]byte(rs), &data); err != nil {
		return nil, "", err
	}

	return []byte(rs), data.Client, nil

}

func removeState(redis *redis.Client, state string) {
	if err := redis.Del(context.Background(), state).Err(); err != nil {
		fmt.Println(err)
		//	TODO log here
	}
}

func cookeSetter(c *fiber.Ctx, code string, username string) {

	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("COOKIE_TOKEN"),
		Value:    code,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		//TODO make true on production https
		Secure:   false,
		SameSite: fiber.CookieSameSiteLaxMode,
	})

	c.Cookie(&fiber.Cookie{
		Name:    os.Getenv("COOKIE_NAME"),
		Value:   username,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	})

}

func cliRedirectString(c string) string {

	return os.Getenv("CLI_CLIENT") + "?code=" + c
}
