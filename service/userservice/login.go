package userservice

import (
	"context"
	"encoding/json"
	"fmt"
	"ostadbun/entity"
	"time"

	"github.com/google/uuid"
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

	uniqueKey := uuid.New().String()
	key := fmt.Sprintf("osbn-u-s:%s:%s", u.Email_Hashe(), uniqueKey)

	if err := r.redis.Set(ctx, key, addUSerID(useragent, MainUserID), time.Hour*24).Err(); err != nil {
		return "", "user", fmt.Errorf("redis set faild -kdfhniu733: %v", err)
	}

	fmt.Println(string(useragent), "magma")

	return uniqueKey, username, nil
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

func addUSerID(user []byte, userid int) []byte {

	var data UseragentData

	_ = json.Unmarshal(user, &data)

	Newdata := NewUseragentData{
		Id:     userid,
		Type:   data.Type,
		Client: data.Client,
		Os:     data.Os,
	}

	finalData, _ := json.Marshal(Newdata)

	return finalData
}
