package httpserver

import (
	"fmt"
	"ostadbun/delivery/httpserver/academic"
	"ostadbun/delivery/httpserver/manipulation"
	"ostadbun/delivery/httpserver/userhandler"
	"ostadbun/service/academicservice"
	"ostadbun/service/activityService"
	"ostadbun/service/manipulationService"
	"ostadbun/service/userservice"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	userService    userservice.User
	manipulService manipulationService.Manipulation

	userHandler         userhandler.Handler
	manipulationHandler manipulation.Handler
	academicHandler     academic.Handler
}

func New(
	userService userservice.User,
	activity activityService.Activity,
	manipulService manipulationService.Manipulation,
	academicService academicservice.Service,

) Server {
	return Server{
		userService:         userService,
		manipulService:      manipulService,
		userHandler:         userhandler.New(userService, activity),
		manipulationHandler: manipulation.New(manipulService, userService),
		academicHandler:     academic.New(academicService),
	}
}

func (s Server) Serve() {

	e := fiber.New()

	e.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowOrigins: "*",
		//AllowOrigins:     "https://ostadbun.tech,https://api.ostadbun.tech",
		//AllowCredentials: true,
		//AllowMethods: "GET,POST,PUT,DELETE",
	}))

	s.userHandler.SetRoutes(e)
	s.manipulationHandler.SetRoutes(e)
	s.academicHandler.SetRoutes(e)

	routes := e.Stack()

	fmt.Println("Registered Routes:")
	for _, stack := range routes {
		for _, route := range stack {
			fmt.Printf("  Method: %s, Path: %s\n", route.Method, route.Path)
		}
	}

	e.Listen(":3000")

}
