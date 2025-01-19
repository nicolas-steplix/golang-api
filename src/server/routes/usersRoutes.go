package routes

import (
	"golang-api/src/controllers"
	"golang-api/src/services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app *fiber.App) {
	service := services.NewUsersService()
	controller := controllers.NewUsersController(service)

	group := app.Group("/users")
	group.Get("/", controller.GetUsers)
	group.Get("/:id", controller.GetUserById)
}
