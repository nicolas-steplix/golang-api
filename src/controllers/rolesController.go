package controllers

import (
	"errors"
	"golang-api/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RolesController struct {
	Service *services.RolesService
}

func NewRolesController(service *services.RolesService) *RolesController {
	return &RolesController{Service: service}
}

func (uc *RolesController) GetRoles(c *fiber.Ctx) error {
	roles, err := uc.Service.GetAllRoles()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(roles)
}

func (uc *RolesController) GetRoleById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": errors.New("invalid role id")})
	}

	role, err := uc.Service.GetRoleById(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if role == nil {
		return c.Status(404).JSON(fiber.Map{"error": errors.New("role not found")})
	}

	return c.JSON(role)
}
