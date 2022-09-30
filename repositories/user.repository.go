package repositories

import "github.com/muh-hizbe/relasi-gorm/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(user *models.User) error
}