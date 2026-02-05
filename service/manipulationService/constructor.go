package manipulationService

//type IRepo = userRepository.AuthRepo
import (
	"ostadbun/repository/postgres/manipulationRepository"
	"ostadbun/service/activityService"
)

type Manipulation struct {
	manipulationRepo manipulationRepository.DB

	activity activityService.Activity
}

func New(activity activityService.Activity, manipulationRepo manipulationRepository.DB) Manipulation {
	return Manipulation{
		activity:         activity,
		manipulationRepo: manipulationRepo,
	}
}
