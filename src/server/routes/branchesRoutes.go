package routes

import (
	"golang-api/src/controllers"
	"golang-api/src/services"

	"github.com/gofiber/fiber/v2"
)

func BranchesRoutes(app *fiber.App) {
	service := services.NewBranchesService()
	controller := controllers.NewBranchesController(service)

	group := app.Group("/branches")
	group.Get("/", controller.GetBranches)
	group.Get("/:id", controller.GetBranchById)
}
