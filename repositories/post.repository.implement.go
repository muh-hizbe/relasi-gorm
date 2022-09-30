package repositories

import (
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

type PostRepositoryImplement struct{}

// Create implements PostRepository
func (*PostRepositoryImplement) Create(post *models.Post) (models.Post, error) {
	db := database.DB.Debug().Create(&post)
	if db.Error != nil {
		return *post, db.Error
	}

	return *post, nil
}

// FindAll implements PostRepository
func (*PostRepositoryImplement) FindAll() ([]models.PostResponseWithTag, error) {
	var posts []models.PostResponseWithTag

	db := database.DB.Preload("User").Preload("Tags").Find(&posts)
	if db.Error != nil {
		return posts, db.Error
	}

	return posts, nil
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImplement{}
}
