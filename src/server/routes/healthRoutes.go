package routes

import (
	"golang-api/src/controllers"
	"golang-api/src/services"

	"github.com/gofiber/fiber/v2"
)

func HealthRoutes(app *fiber.App) {
	service := services.NewHealthService()
	controller := controllers.NewHealthController(service)

	group := app.Group("/health")
	group.Get("/", controller.Check)
}
