package main

import (
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
	//dbConf := database.New()

	oauth := oauthservice.NewOAuthService()
	userSvc := userservice.New(*oauth)

	server := httpserver.New(userSvc)

	server.Serve()

}
