package controllers

import (
	"errors"
	"golang-api/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BranchesController struct {
	Service *services.BranchesService
}

func NewBranchesController(service *services.BranchesService) *BranchesController {
	return &BranchesController{Service: service}
}

func (uc *BranchesController) GetBranches(c *fiber.Ctx) error {
	branches, err := uc.Service.GetAllBranches()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(branches)
}

func (uc *BranchesController) GetBranchById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": errors.New("invalid branch id")})
	}

	branch, err := uc.Service.GetBranchById(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if branch == nil {
		return c.Status(404).JSON(fiber.Map{"error": errors.New("branch not found")})
	}

	return c.JSON(branch)
}
