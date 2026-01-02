package userservice

//type IRepo = userRepository.AuthRepo
import (
	"ostadbun/repository/userRepository"
	"ostadbun/service/activity"

	"ostadbun/service/oauthservice"

	"github.com/redis/go-redis/v9"
)

//type IRepo interface {
//	ExistingCheck(email string) (int, string, bool)
//	RegisterUser(user entity.User) (int, error)
//	AdminByWho(userID string) (int, error)
//	SwitchPermission(userID int, masterID int) (bool, error)
//}

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
