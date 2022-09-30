package repositories

import (
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

type PostTagRepositoryImplement struct {
}

// Create implements PostTagRepository
func (*PostTagRepositoryImplement) Create(postTag *models.PostTag) error {
	db := database.DB.Create(&postTag)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func NewPostTagRepository() PostTagRepository {
	return &PostTagRepositoryImplement{}
}
