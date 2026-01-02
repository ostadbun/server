package userhandler

import (
	"fmt"
	"ostadbun/service/activity/activityconstants"
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

	isAdminNow, errSwitch := h.authSvc.SwitchPermission(userId, masterID)

	if errSwitch != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errSwitch.Error())
	}

	var triggerErr error
	if isAdminNow {
		triggerErr = h.authSvc.ActivityTrigger(masterID, activityconstants.Trigger_UnMakeAdmin)
	} else {
		triggerErr = h.authSvc.ActivityTrigger(masterID, activityconstants.Trigger_MakeAdmin)
	}

	fmt.Println(triggerErr)
	//if triggerErr != nil {
	//	return c.Status(fiber.StatusInternalServerError).SendString(triggerErr.Error())
	//}

	return c.Render("user/switchPermission", fiber.Map{})
}
