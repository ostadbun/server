package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

type Req struct {
	Id string `json:"id"`
}

func main() {

	e := echo.New()

	e.POST("google", func(c echo.Context) error {

		// token := ""

		var token Req

		e := c.Bind(&token)

		if e != nil {
			fmt.Println(e.Error())
		}

		payload, err := idtoken.Validate(context.Background(), token.Id, "23967475973-ginlrgh2j8b4aqhtt613r3lut6k9e981.apps.googleusercontent.com")

		if err != nil {

			return c.String(200, err.Error())
		}

		// claims := payload.Claims

		// email := claims["email"]
		// name := claims["name"]
		// picture := claims["picture"]

		return c.JSON(200, payload.Claims)

	})

	e.Start(":3000")

}
