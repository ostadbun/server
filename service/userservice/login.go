package userservice

import (
	"context"
	"encoding/json"
	"fmt"
	"ostadbun/entity"
	"time"

	"github.com/google/uuid"
)

func (r User) Login(u entity.User) (string, error) {
	ctx := context.Background()

	userID, isExist := r.repo.ExistingCheck(u.Email)
	fmt.Println(isExist, userID)

	if !isExist {
		_, err := r.repo.RegisterUser(u)
		if err != nil {
			return "", fmt.Errorf("redis set failed -1242312: %v", err)
		}

		newID := uuid.New().String()
		key := fmt.Sprintf("user:%s:%s", u.Email, newID)

		v, err := json.Marshal(u)
		if err != nil {
			return "", fmt.Errorf("marsh set failed -r1423124100023599: %v", err)
		}
		if err := r.redis.JSONSet(ctx, key, v, 72*time.Hour).Err(); err != nil {
			return "", fmt.Errorf("redis set failed -912324: %v", err)
		}

		return newID, nil
	} else {
		keys, err := r.redis.Keys(ctx, fmt.Sprintf("user:%s:*", u.Email)).Result()
		if err != nil {
			return "", fmt.Errorf("redis set failed -r423423423599: %v", err)
		}

		if len(keys) > 0 {

			for _, key := range keys {
				val, _ := r.redis.Get(ctx, key).Result()
				fmt.Println(val)

			}

		} else {
			newID := uuid.New().String()
			key := fmt.Sprintf("user:%s:%s", u.Email, newID)

			if err := r.redis.JSONSet(ctx, key, u, 72*time.Hour).Err(); err != nil {
				return "", fmt.Errorf("redis set failed -112314234: %v", err)
			}
		}
		return "existing", nil
	}
}
