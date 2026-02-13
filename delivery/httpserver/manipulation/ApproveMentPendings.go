package manipulation

import (
	"fmt"
	"os"
	"ostadbun/pkg/httpstorage"
	yesWords "ostadbun/pkg/yesWorlds"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) ApprovementLessonPending(c *fiber.Ctx) error {

	status, targetID, reason, err := Validating(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user id not found",
		})
	}

	var DoError error
	if status {
		DoError = h.manipulSVC.ApprvingLesson(c.Context(), int64(targetID), int64(userID))
	} else {
		DoError = h.manipulSVC.RejectLesson(c.Context(), &reason, int64(targetID), int64(userID))
	}

	if DoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": DoError.Error(),
		})
	}
	return c.SendString("wow it done!")
}

func (h Handler) ApprovementUnivPending(c *fiber.Ctx) error {

	status, targetID, reason, err := Validating(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user id not found",
		})
	}

	var DoError error
	if status {
		DoError = h.manipulSVC.ApprvingUniversity(c.Context(), int64(targetID), int64(userID))
	} else {
		DoError = h.manipulSVC.RejectUniversity(c.Context(), &reason, int64(targetID), int64(userID))
	}

	if DoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": DoError.Error(),
		})
	}
	return c.SendString("wow it done!")
}

func (h Handler) ApprovementProfPending(c *fiber.Ctx) error {

	status, targetID, reason, err := Validating(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user id not found",
		})
	}

	var DoError error
	if status {
		DoError = h.manipulSVC.ApprvingProfessor(c.Context(), int64(targetID), int64(userID))
	} else {
		DoError = h.manipulSVC.RejectProfessor(c.Context(), &reason, int64(targetID), int64(userID))
	}

	if DoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": DoError.Error(),
		})
	}
	return c.SendString("wow it done!")
}

func (h Handler) ApprovementMajorPending(c *fiber.Ctx) error {

	status, targetID, reason, err := Validating(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user id not found",
		})
	}

	var DoError error
	if status {
		DoError = h.manipulSVC.ApprvingMajor(c.Context(), int64(targetID), int64(userID))
	} else {
		DoError = h.manipulSVC.RejectMajor(c.Context(), &reason, int64(targetID), int64(userID))
	}

	if DoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": DoError.Error(),
		})
	}
	return c.SendString("wow it done!")
}

func Validating(c *fiber.Ctx) (bool, int, string, error) {
	sts := c.Params("status")

	status := yesWords.IsYes(sts)
	tid := c.Params("targetID")
	targetID, errTid := strconv.Atoi(tid)

	if errTid != nil {
		return false, 0, "", fmt.Errorf("id is invalid, should be a number")
	}

	reason := string(c.BodyRaw())
	if len(reason) < 5 {
		reason = os.Getenv("REASON_REJECTION")
	}

	fmt.Println("reason:", reason)
	return status, targetID, reason, nil
}
