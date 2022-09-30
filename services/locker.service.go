package services

import "github.com/muh-hizbe/relasi-gorm/models"

type LockerService interface {
	GetAllData() ([]models.Locker, error)
	New(lockerRequest models.LockerRequest) (models.LockerResponse, error)
	Validation(lockerRequest models.LockerRequest) error
}
