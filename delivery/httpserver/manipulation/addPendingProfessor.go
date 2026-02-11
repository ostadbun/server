package manipulation

import (
	"ostadbun/entity"
	manipulationParam "ostadbun/param/manipulation"
	"ostadbun/pkg/httpstorage"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) addPendingProfessor(c *fiber.Ctx) error {

	userId, err := httpstorage.Get(c, "user_id").Number()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	var acceptData manipulationParam.PendingProfessor

	er := c.BodyParser(&acceptData)

	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "error on parsing request",
			"details": er,
		})
	}

	unversityData := entity.PendingProfessor{
		Name:               acceptData.Name,
		NameEnglish:        acceptData.NameEnglish,
		DescriptionEnglish: acceptData.DescriptionEnglish,
		Description:        &acceptData.Description,
		ImageUrl:           &acceptData.ImageUrl,
		EducationHistory:   &acceptData.EducationHistory,
		SubmittedBy:        int64(userId),
	}

	return h.manipulSVC.AddPendingProfessor(unversityData, userId)

}
