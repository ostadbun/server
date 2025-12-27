package httpserver

import (
	"fmt"
	"ostadbun/delivery/httpserver/userhandler"
	"ostadbun/service/userservice"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	userService userservice.User
	UserHandler userhandler.Handler
}

func New(
	userService userservice.User,
	redis *redis.Client,
) Server {
	return Server{
		UserHandler: userhandler.New(userService, redis),
		userService: userService,
	}
}

func (s Server) Serve() {

	e := fiber.New()

	e.Use(cors.New())

	s.UserHandler.SetRoutes(e)

	routes := e.Stack()

	fmt.Println("Registered Routes:")
	for _, stack := range routes {
		for _, route := range stack {
			fmt.Printf("  Method: %s, Path: %s\n", route.Method, route.Path)
		}
	}

	e.Listen(":3000")

}
