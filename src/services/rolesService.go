package services

import (
	"golang-api/src/database"
	"golang-api/src/database/entities"
)

type RolesService struct{}

func NewRolesService() *RolesService {
	return &RolesService{}
}

func (s *RolesService) GetAllRoles() ([]entities.Role, error) {
	var roles []entities.Role

	result := database.DB.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}

func (s *RolesService) GetRoleById(id int) (*entities.Role, error) {
	var role entities.Role

	result := database.DB.First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}
