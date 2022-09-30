package services

import "github.com/muh-hizbe/relasi-gorm/models"

type UserService interface {
	GetAllData() ([]models.User, error)
	New(userRequest models.UserRequest) (models.UserResponse, error)
	Validation(userRequest models.UserRequest) error
}