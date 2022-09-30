package repositories

import "github.com/muh-hizbe/relasi-gorm/models"

type LockerRepository interface {
	FindAll() ([]models.Locker, error)
	Create(locker *models.Locker) (models.Locker, error)
}