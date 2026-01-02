package activityService

import (
	"ostadbun/repository/activityRepository"
)

type Activity struct {
	repo *activityRepository.DB
}

func New(repo *activityRepository.DB) Activity {
	return Activity{
		repo: repo,
	}
}
