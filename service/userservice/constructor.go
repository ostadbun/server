package userservice

//type IRepo = userRepository.AuthRepo
import (
	"context"
	"ostadbun/repository/postgres/userRepository"
	"ostadbun/service/activityService"

	"ostadbun/service/oauthservice"
)

type RedisOauth interface {
	AddUserSession(ctx context.Context, Email_Hashe string, useragent []byte, MainUserID int) (string, string, error)
	AuthCheck(ctx context.Context, token string) (int, error)
	RemoveState(ctx context.Context, state string)
	CheckIntoRedis(ctx context.Context, key string) ([]byte, string, error)
}

type User struct {
	oauth    oauthservice.OAuthService
	activity activityService.Activity
	repo     *userRepository.DB
	redis    RedisOauth
}

func New(oauth oauthservice.OAuthService, activity activityService.Activity, redis RedisOauth, repo *userRepository.DB) User {
	return User{
		oauth:    oauth,
		activity: activity,
		redis:    redis,
		repo:     repo,
	}
}
