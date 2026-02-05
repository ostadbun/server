package userservice

import (
	"context"
	"fmt"
	"ostadbun/entity"
)

func (r User) Login(u entity.User, useragent []byte) (code string, name string, err error) {
	ctx := context.Background()

	userID, username, isExist := r.repo.ExistingCheck(u.Email)
	fmt.Println(isExist, userID, u.Email, "sa7hf83i72i73e")

	var MainUserID int
	if !isExist {

		MainUserID, err = r.repo.RegisterUser(u)

		if err != nil {
			return "", "", fmt.Errorf("register faild -1242312: %v", err)
		}
	} else {
		MainUserID = userID
	}

	code, _, err = r.redis.AddUserSession(ctx, u.Email_Hashe(), useragent, MainUserID)

	if err != nil {
		return "", "", fmt.Errorf("redis error 389wh19we13e: %v", err)
	}

	return code, username, nil

}

type UseragentData struct {
	Type   string `json:"type"`
	Client string `json:"client"`
	Os     string `json:"os"`
}

type NewUseragentData struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Client string `json:"client"`
	Os     string `json:"os"`
}
