package redisUser

import (
	"context"
	"encoding/json"
	"fmt"
	"ostadbun/service/userservice"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisUser struct {
	redis *redis.Client
}

func New(client *redis.Client) *RedisUser {
	return &RedisUser{redis: client}
}

func (r *RedisUser) AddUserSession(ctx context.Context, Email_Hashe string, useragent []byte, MainUserID int) (string, string, error) {

	uniqueKey := uuid.New().String()
	key := fmt.Sprintf("osbn-u-s:%s:%s", Email_Hashe, uniqueKey)

	if err := r.redis.Set(ctx, key, addUSerID(useragent, MainUserID), time.Hour*24).Err(); err != nil {
		return "", "user", fmt.Errorf("redis set faild -kdfhniu733: %v", err)
	}

	return uniqueKey, key, nil
}

func addUSerID(user []byte, userid int) []byte {

	var data userservice.UseragentData

	_ = json.Unmarshal(user, &data)

	Newdata := userservice.NewUseragentData{
		Id:     userid,
		Type:   data.Type,
		Client: data.Client,
		Os:     data.Os,
	}

	finalData, _ := json.Marshal(Newdata)

	return finalData
}

type NewUseragentData struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Client string `json:"client"`
	Os     string `json:"os"`
}

func (r *RedisUser) AuthCheck(ctx context.Context, token string) (int, error) {

	key := fmt.Sprintf("osbn-u-s:*:%s", token)
	//va, err := r.Get(context.Background(), key).Result()

	va, err := r.redis.Keys(ctx, key).Result()

	if err != nil {
		return 0, fmt.Errorf("authorization faild %v", err)
	}

	if len(va) != 1 {
		//	TODO log why store more then 1 uui?
		//if len(va) > 1 {}

		return 0, fmt.Errorf("access denied %v", err)
	}

	userG, errU := r.redis.Get(context.Background(), va[0]).Result()
	if errU != nil {
		return 0, fmt.Errorf("authorization faild %v", err)
	}

	var user NewUseragentData
	errJSON := json.Unmarshal([]byte(userG), &user)

	if errJSON != nil {
		// TODO log here

		return 0, fmt.Errorf("authorization faild %v", err)
	}

	ok, err := r.redis.Expire(context.Background(), va[0], time.Hour*24).Result()

	if err != nil {
		return 0, fmt.Errorf("authorization faild %v", err)
	}

	if !ok {
		return 0, fmt.Errorf("authorization faild %v", err)
	}

	return user.Id, err
}

func (r *RedisUser) RemoveState(ctx context.Context, state string) {

	if err := r.redis.Del(context.Background(), state).Err(); err != nil {
		fmt.Println(err)
		//	TODO log here
	}

}

type DeviceInfo struct {
	// mobile | desktop
	Type string `json:"type"`

	// web | terminal
	Client string `json:"client"`

	// android | ios | windows | mac | linux
	OS string `json:"os"`
}

func (r *RedisUser) CheckIntoRedis(ctx context.Context, key string) ([]byte, string, error) {
	//return  userAgentData, client, err

	rs, errR := r.redis.Get(ctx, key).Result()

	if errR != nil {
		return nil, "", errR
	}

	var data DeviceInfo

	if err := json.Unmarshal([]byte(rs), &data); err != nil {
		return nil, "", err
	}

	return []byte(rs), data.Client, nil

}
