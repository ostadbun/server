package userservice

import (
	"strconv"
)

func (r User) LevelCalculator(usrID string) (int, error) {

	userID, ok := strconv.Atoi(usrID)

	if ok != nil {
		return -1, ok
	}

	return r.activity.LevelCalculator(userID)
}
