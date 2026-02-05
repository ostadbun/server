package httpserver

import (
	"fmt"
	"ostadbun/delivery/httpserver/manipulation"
	"ostadbun/delivery/httpserver/userhandler"
	"ostadbun/service/activityService"
	"ostadbun/service/manipulationService"
	"ostadbun/service/userservice"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	userService    userservice.User
	manipulService manipulationService.Manipulation

	UserHandler         userhandler.Handler
	manipulationHanlder manipulation.Handler
}

func New(
	userService userservice.User,
	activity activityService.Activity,
	manipulService manipulationService.Manipulation,
) Server {
	return Server{
		userService:         userService,
		manipulService:      manipulService,
		UserHandler:         userhandler.New(userService, activity),
		manipulationHanlder: manipulation.New(manipulService, userService),
	}
}

func (s Server) Serve() {

	e := fiber.New()

	e.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8713",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	s.UserHandler.SetRoutes(e)
	s.manipulationHanlder.SetRoutes(e)

	routes := e.Stack()

	fmt.Println("Registered Routes:")
	for _, stack := range routes {
		for _, route := range stack {
			fmt.Printf("  Method: %s, Path: %s\n", route.Method, route.Path)
		}
	}

	e.Listen(":3000")

}
