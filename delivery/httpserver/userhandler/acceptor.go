package userhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"ostadbun/entity"
	"ostadbun/param/userparam"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func (h Handler) acceptor(c *fiber.Ctx) error {

	ProviderName := c.Params("provider")
	Code := c.Query("code")

	state := c.Query("state")

	defer removeState(h.redis, state)

	userAgentData, err := checkIntoRedis(state, h.redis)
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

		code, err := h.authSvc.Login(data, userAgentData)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		// get login jwt
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": code,
		})

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

		code, err := h.authSvc.Login(data, userAgentData)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("65asfh872r - %s", errC))
		}
		// get login jwt
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": code,
		})
	}

	return c.SendString("end")
}

func checkIntoRedis(key string, rds *redis.Client) ([]byte, error) {

	rs, err := rds.Get(context.Background(), key).Result()

	if err != nil {
		return nil, err
	}

	return []byte(rs), nil

}

func removeState(redis *redis.Client, state string) {
	if err := redis.Del(context.Background(), state).Err(); err != nil {
		fmt.Println(err)
		//	TODO log here
	}
}
