package repositories

import "github.com/muh-hizbe/relasi-gorm/models"

type PostRepository interface {
	FindAll() ([]models.PostResponseWithTag, error)
	Create(post *models.Post) (models.Post, error)
}
