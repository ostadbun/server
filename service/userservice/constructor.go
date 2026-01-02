package userservice

//type IRepo = userRepository.AuthRepo
import (
	"ostadbun/repository/userRepository"

	"ostadbun/service/activityService"

	"ostadbun/service/oauthservice"

	"github.com/redis/go-redis/v9"
)

type User struct {
	oauth    oauthservice.OAuthService
	activity activityService.Activity
	repo     *userRepository.DB
	redis    *redis.Client
}

func New(oauth oauthservice.OAuthService, activity activityService.Activity, redis *redis.Client, repo *userRepository.DB) User {
	return User{
		oauth:    oauth,
		activity: activity,
		redis:    redis,
		repo:     repo,
	}
}
