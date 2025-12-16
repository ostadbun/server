package entity

import (
	"fmt"
	"ostadbun/interface/entity"

	"github.com/gofiber/fiber/v2"
)

func (c Controller) Entity(f *fiber.Ctx) error {

	var data entity.ISearch
	err := f.BodyParser(&data)

	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("🐥")
	}

	d := c.Svc.Search(data.Q)

	if err != nil {
		return f.SendString(err.Error())
	}

	return f.JSON(d)
}
