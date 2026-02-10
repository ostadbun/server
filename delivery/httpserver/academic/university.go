package academic

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) University(c *fiber.Ctx) error {

	qery := "تهران"

	dta, err := h.academicService.UniversitySearch(qery)
	fmt.Println(qery, dta, err)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dta,
	})

}
