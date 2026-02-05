package main

import (
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/database"
	"ostadbun/repository/postgres/activityRepository"
	"ostadbun/repository/postgres/manipulationRepository"
	"ostadbun/repository/postgres/userRepository"
	"ostadbun/repository/redis/redisActivity"
	"ostadbun/repository/redis/redisOauth"
	"ostadbun/repository/redis/redisUser"
	"ostadbun/service/activityService"

	"ostadbun/service/manipulationService"

	"github.com/joho/godotenv"

	"ostadbun/delivery/httpserver"
	"ostadbun/service/oauthservice"
	"ostadbun/service/userservice"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}
	dbConf := database.New()

	redisClient := redisAdaptor.New()

	//oauth
	Oauthrds := redisOauth.New(redisClient)
	oauth := oauthservice.NewOAuthService(Oauthrds)

	//activity
	activRds := redisActivity.New(redisClient)
	activRepo := activityRepository.New(dbConf)
	activSvc := activityService.New(activRepo, activRds)

	//user
	userRds := redisUser.New(redisClient)
	userRepo := userRepository.New(dbConf)
	userSvc := userservice.New(*oauth, activSvc, userRds, userRepo)

	//manipulation
	maniRepo := manipulationRepository.New(dbConf)
	maniSVC := manipulationService.New(activSvc, *maniRepo)

	//engine
	server := httpserver.New(userSvc, activSvc, maniSVC)

	server.Serve()

}
