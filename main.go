package main

import (
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/database"
	"ostadbun/repository/userRepository"

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

	userRepo := userRepository.Make(dbConf.Conn())
	userSvc := userservice.New(*oauth, rds, userRepo)

	server := httpserver.New(userSvc, rds)

	server.Serve()

}
