package userhandler

import (
	"fmt"
	Activityconstants "ostadbun/pkg/constants"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) switchPermission(c *fiber.Ctx) error {

	usida := c.Params("userid")

	userId, errN := strconv.Atoi(usida)

	if errN != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("userID should be a number")
	}

	fmt.Println(c.Locals("user_id"))
	mstrid := c.Locals("user_id")

	masterID, ok := strconv.Atoi(mstrid.(string))

	fmt.Println(masterID, ok)
	if ok != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("masterID not found wdaij")
	}

	isAdminNow, errSwitch := h.userSvc.SwitchPermission(userId, masterID)

	if errSwitch != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errSwitch.Error())
	}

	var triggerErr error
	if isAdminNow {
		triggerErr = h.activitySvc.Trigger(c.Context(), masterID, Activityconstants.Trigger_UnMakeAdmin)
	} else {
		triggerErr = h.activitySvc.Trigger(c.Context(), masterID, Activityconstants.Trigger_MakeAdmin)
	}

	if triggerErr != nil {
		//TODO log here
		fmt.Println(triggerErr.Error())
	}

	return c.SendString(msg(isAdminNow))
}

func msg(a bool) string {
	if a {
		return "Converted to admin"
	} else {
		return "Converted to user"
	}
}
