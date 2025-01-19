package controllers

import (
	"errors"
	"golang-api/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	Service *services.UsersService
}

func NewUsersController(service *services.UsersService) *UsersController {
	return &UsersController{Service: service}
}

func (uc *UsersController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.Service.GetAllUsers()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (uc *UsersController) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": errors.New("invalid user id")})
	}

	user, err := uc.Service.GetUserById(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if user == nil {
		return c.Status(404).JSON(fiber.Map{"error": errors.New("user not found")})
	}

	return c.JSON(user)
}
