package repositories

import (
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

type TagRepositoryImplement struct {
}

// Create implements TagRepository
func (*TagRepositoryImplement) Create(tag *models.Tag) error {
	db := database.DB.Debug().Create(&tag)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

// FindAll implements TagRepository
func (*TagRepositoryImplement) FindAll() ([]models.TagResponseWithPost, error) {
	var tags []models.TagResponseWithPost

	db := database.DB.Preload("Posts").Find(&tags)
	if db.Error != nil {
		return tags, db.Error
	}

	return tags, nil
}

func NewTagReposotory() TagRepository {
	return &TagRepositoryImplement{}
}
