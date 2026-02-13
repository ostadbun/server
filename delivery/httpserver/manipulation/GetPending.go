package manipulation

import (
	"ostadbun/entity"

	"github.com/gofiber/fiber/v2"
)

type IProvide struct {
	University []entity.PendingUniversity `json:"university,omitempty"`
	Lesson     []entity.PendingLesson     `json:"lesson,omitempty"`
	Professor  []entity.PendingProfessor  `json:"professor,omitempty"`
	Major      []entity.PendingMajor      `json:"major,omitempty"`
}

func (h Handler) GetPending(c *fiber.Ctx) error {

	uni, errU := h.manipulSVC.GetPendingUniversity()
	les, errL := h.manipulSVC.GetPendingLesson()
	prof, errP := h.manipulSVC.GetPendingProfessor()
	major, errM := h.manipulSVC.GetPendingMajor()

	if errM != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errM,
			"code":  "major",
		})
	}

	if errL != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errL,
			"code":  "lesson",
		})
	}

	if errU != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errU,
			"code":  "university",
		})
	}

	if errP != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errP,
			"code":  "prof",
		})
	}

	if errU != nil && errL != nil && errP != nil && errM != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Something went wrong",
			"errorsOn": map[string]string{
				"University": errU.Error(),
				"Professor":  errP.Error(),
				"Lesson":     errL.Error(),
				"Major":      errM.Error(),
			},
		})
	}

	Data := IProvide{
		Lesson:     les,
		University: uni,
		Professor:  prof,
		Major:      major,
	}

	return c.JSON(Data)

}
