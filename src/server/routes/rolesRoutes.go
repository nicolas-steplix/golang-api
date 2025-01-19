package routes

import (
	"golang-api/src/controllers"
	"golang-api/src/services"

	"github.com/gofiber/fiber/v2"
)

func RolesRoutes(app *fiber.App) {
	service := services.NewRolesService()
	controller := controllers.NewRolesController(service)

	group := app.Group("/roles")
	group.Get("/", controller.GetRoles)
	group.Get("/:id", controller.GetRoleById)
}
