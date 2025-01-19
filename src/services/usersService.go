package services

import (
	"golang-api/src/database"
	"golang-api/src/database/entities"
)

type UsersService struct{}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) GetAllUsers() ([]entities.User, error) {
	var users []entities.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *UsersService) GetUserById(id int) (*entities.User, error) {
	var user entities.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
