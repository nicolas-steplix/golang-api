package controllers

import (
	"golang-api/src/services"

	"github.com/gofiber/fiber/v2"
)

type HealthController struct {
	Service *services.HealthService
}

func NewHealthController(service *services.HealthService) *HealthController {
	return &HealthController{Service: service}
}

func (uc *HealthController) Check(c *fiber.Ctx) error {
	check, err := uc.Service.Check()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(check)
}
