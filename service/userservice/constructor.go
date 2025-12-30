package userservice

//type IRepo = userRepository.AuthRepo
import (
	"ostadbun/entity"
	"ostadbun/service/oauthservice"

	"github.com/redis/go-redis/v9"
)

type IRepo interface {
	ExistingCheck(email string) (int, string, bool)
	RegisterUser(user entity.User) (int, error)
	AdminByWho(userID string) (int, error)
}
type User struct {
	oauth oauthservice.OAuthService
	repo  IRepo
	redis *redis.Client
}

func New(oauth oauthservice.OAuthService, redis *redis.Client, repo IRepo) User {
	return User{
		oauth: oauth,
		redis: redis,
		repo:  repo,
	}
}
