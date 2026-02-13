package manipulation

import (
	"ostadbun/entity"
	manipulationParam "ostadbun/param/manipulation"
	notify "ostadbun/pkg/bale/notif"
	"ostadbun/pkg/httpstorage"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) addPendingUniversity(c *fiber.Ctx) error {

	userId, err := httpstorage.Get(c, "user_id").Number()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	var acceptData manipulationParam.PendingUniversity

	er := c.BodyParser(&acceptData)

	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "error on parsing request",
			"details": er,
		})
	}

	data := entity.PendingUniversity{
		Name:               acceptData.Name,
		NameEnglish:        acceptData.NameEnglish,
		DescriptionEnglish: acceptData.DescriptionEnglish,
		Description:        &acceptData.Description,
		ImageUrl:           &acceptData.ImageUrl,
		City:               acceptData.City,
		Category:           acceptData.Category,
		SubmittedBy:        int64(userId),
	}

	go func() {
		if err := notify.NotifyNewUniversity(data); err != nil {
			//TODO log here
		}
	}()

	return h.manipulSVC.AddPendingUniversity(data, userId)

}
