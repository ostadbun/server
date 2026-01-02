package main

import (
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/database"
	"ostadbun/repository/activityRepository"
	"ostadbun/repository/userRepository"
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

	rds := redisAdaptor.New()

	oauth := oauthservice.NewOAuthService(rds)

	activRepo := activityRepository.New(dbConf, rds)
	activSvc := activityService.New(activRepo)

	userRepo := userRepository.New(dbConf)
	userSvc := userservice.New(*oauth, activSvc, rds, userRepo)

	maniSVC := manipulationService.New(activSvc)

	server := httpserver.New(userSvc, rds, activSvc, maniSVC)

	server.Serve()

}
