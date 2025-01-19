package services

import (
	"golang-api/src/database"
	"golang-api/src/database/entities"
)

type BranchesService struct{}

func NewBranchesService() *BranchesService {
	return &BranchesService{}
}

func (s *BranchesService) GetAllBranches() ([]entities.Branch, error) {
	var branches []entities.Branch

	result := database.DB.Find(&branches)
	if result.Error != nil {
		return nil, result.Error
	}

	return branches, nil
}

func (s *BranchesService) GetBranchById(id int) (*entities.Branch, error) {
	var branch entities.Branch

	result := database.DB.First(&branch, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &branch, nil
}
