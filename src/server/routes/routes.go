package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	UsersRoutes(app)
	RolesRoutes(app)
	BranchesRoutes(app)
	HealthRoutes(app)
}
