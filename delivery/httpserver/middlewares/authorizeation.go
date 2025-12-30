package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func Auth(r *redis.Client) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		tkn := os.Getenv("COOKIE_TOKEN")

		token := c.Cookies(tkn)

		fmt.Println(token)
		if len(token) < 10 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "access denied",
			})
		}

		key := fmt.Sprintf("osbn-u-s:*:%s", token)
		//va, err := r.Get(context.Background(), key).Result()

		va, err := r.Keys(context.Background(), key).Result()

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": fmt.Sprintf("authorization faild %v", err),
			})
		}

		if len(va) != 1 {
			//	TODO log why store more then 1 uui?
			//if len(va) > 1 {}

			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": fmt.Sprintf("access denied"),
			})
		}

		userG, errU := r.Get(context.Background(), va[0]).Result()
		if errU != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": fmt.Sprintf("authorization faild %v", errU),
			})
		}

		var user NewUseragentData
		errJSON := json.Unmarshal([]byte(userG), &user)

		if errJSON != nil {
			// TODO log here
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": fmt.Sprintf("authorization faild %v", errJSON),
			})
		}
		ok, err := r.Expire(context.Background(), va[0], time.Hour*24).Result()
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": fmt.Sprintf("access denied"),
			})
		}
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": fmt.Sprintf("access denied"),
			})
		}
		ID := strconv.Itoa(user.Id)

		c.Set("user_id", ID)

		return c.Next()

	}

}

type NewUseragentData struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Client string `json:"client"`
	Os     string `json:"os"`
}
