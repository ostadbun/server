package userservice

import (
	"context"
	"fmt"
	"ostadbun/entity"
	"time"

	"github.com/google/uuid"
)

func (r User) Login(u entity.User, useragent []byte) (string, error) {
	ctx := context.Background()

	userID, isExist := r.repo.ExistingCheck(u.Email)
	fmt.Println(isExist, userID)

	if !isExist {
		_, err := r.repo.RegisterUser(u)
		if err != nil {
			return "", fmt.Errorf("register faild -1242312: %v", err)
		}
	}

	uniqueKey := uuid.New().String()
	key := fmt.Sprintf("osbn-u-s:%s:%s", u.Email_Hashe(), uniqueKey)

	if err := r.redis.Set(ctx, key, useragent, time.Hour*71).Err(); err != nil {
		return "", fmt.Errorf("redis set faild -kdfhniu733: %v", err)
	}

	fmt.Println(string(useragent), "magma")

	return uniqueKey, nil
}
