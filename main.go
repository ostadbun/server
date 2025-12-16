package main

import (
	"fmt"
	"log"
	"ostadbun/controller"
	"ostadbun/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	dbConf := database.Config()

	controller.Config(app, dbConf)

	routes := app.Stack()

	fmt.Println("Registered Routes:")
	for _, stack := range routes {
		for _, route := range stack {
			fmt.Printf("  Method: %s, Path: %s\n", route.Method, route.Path)
		}
	}

	log.Fatal(app.Listen(":3000"))
}
