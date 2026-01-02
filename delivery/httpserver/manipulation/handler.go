package manipulation

import (
	manipulation "ostadbun/service/manipulationService"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	rds        *redis.Client
	manipulSVC manipulation.Manipulation
}

func New(rds *redis.Client, manipulSVC manipulation.Manipulation) Handler {
	return Handler{
		rds:        rds,
		manipulSVC: manipulSVC,
	}
}
