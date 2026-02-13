package middlewares

import (
	"fmt"
	"os"
	"ostadbun/service/userservice"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(u userservice.User) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {

		token, err := GetAuthToken(c)

		fmt.Println(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "access denied",
			})
		}

		userID, err := u.AuthCheck(c.Context(), token)
		fmt.Println(userID, err)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "session not found",
				"error":   err.Error(),
			})
		}

		ID := strconv.Itoa(userID)

		c.Locals("user_id", ID)

		fmt.Println("1323", ID)
		return c.Next()

	}

}

func GetAuthToken(c *fiber.Ctx) (string, error) {
	tkn := os.Getenv("COOKIE_TOKEN")

	cookieToken := c.Cookies(tkn)

	headerToken := c.Get("Authorization")

	if cookieToken != "" {
		return cookieToken, nil
	}

	if headerToken != "" {
		headerToken = strings.Replace(headerToken, "Bearer ", "", 1)
		return headerToken, nil
	}

	return "", fmt.Errorf("user not authenticated")
}

type NewUseragentData struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Client string `json:"client"`
	Os     string `json:"os"`
}
