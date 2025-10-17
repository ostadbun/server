package main

import (
	envConf "ostadbun/config"
	"ostadbun/database"
	"ostadbun/server"
)

func main() {

	env := envConf.GetConfig()
	db := database.NewPostgresSql(env)
	app := server.NewFiberServer(env, db)

	app.Start()
}
