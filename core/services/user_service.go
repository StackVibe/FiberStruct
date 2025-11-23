package services

import (
	"Test/core/models"
	"Test/core/repositories"
	"Test/core/utilities/general"
	"errors"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (service UserService) GetAllUsers() ([]models.User, error) {
	return service.Repository.GetAll()
}

func (service UserService) CreateUser(user *models.User) (*models.User, error) {
	existUser, err := service.Repository.GetByEmail(user.Email)
	if err == nil && existUser.ID != 0 {
		return nil, errors.New("new user already exist")
	}

	user.Password = general.MakeHash(user.Password)

	return service.Repository.Create(user)
}
