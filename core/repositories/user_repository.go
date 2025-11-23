package repositories

import (
	"Test/core/models"
	"Test/database/config"
)

type UserRepository struct {
}

func (repository UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (repository UserRepository) GetById(id int) (*models.User, error) {
	var user models.User
	err := config.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (repository UserRepository) Create(user *models.User) (*models.User, error) {
	return nil, config.DB.Create(user).Error
}

func (repository UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("Email = ?", email).First(&user).Error
	return &user, err
}
