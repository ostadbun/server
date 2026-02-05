package middlewares

import (
	"fmt"
	"os"
	"ostadbun/service/userservice"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Auth(u userservice.User) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		tkn := os.Getenv("COOKIE_TOKEN")

		token := c.Cookies(tkn)

		fmt.Println(token)
		if len(token) < 10 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "access denied",
			})
		}
		userID, err := u.AuthCheck(c.Context(), token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		ID := strconv.Itoa(userID)

		c.Locals("user_id", ID)

		fmt.Println("1323", ID)
		return c.Next()

	}

}

type NewUseragentData struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Client string `json:"client"`
	Os     string `json:"os"`
}
