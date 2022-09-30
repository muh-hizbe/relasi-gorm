package repositories

import (
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

type LockerRepositoryImplement struct {
}

// Create implements LockerRepository
func (*LockerRepositoryImplement) Create(locker *models.Locker) (models.Locker, error) {
	db := database.DB.Create(&locker)

	if db.Error != nil {

		return *locker, db.Error
	}
	return *locker, nil
}

// FindAll implements LockerRepository
func (*LockerRepositoryImplement) FindAll() ([]models.Locker, error) {
	var lockers []models.Locker
	db := database.DB.Preload("User").Find(&lockers)
	if db.Error != nil {
		return lockers, db.Error
	}
	return lockers, nil
}

func NewLockerRepository() LockerRepository {
	return &LockerRepositoryImplement{}
}
