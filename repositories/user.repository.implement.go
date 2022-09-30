package repositories

import (
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

type UserRepositoryImplement struct {
}

// Create implements UserRepository
func (*UserRepositoryImplement) Create(user *models.User) error {
	db := database.DB.Create(&user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

// FindAll implements UserRepository
func (*UserRepositoryImplement) FindAll() ([]models.User, error) {
	var users []models.User

	db:=database.DB.Preload("Locker").Preload("Posts").Find(&users)
	if db.Error != nil {
		return users, db.Error
	}

	return users, nil
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplement{}
}
