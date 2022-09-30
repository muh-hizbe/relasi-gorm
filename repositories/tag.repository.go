package repositories

import "github.com/muh-hizbe/relasi-gorm/models"

type TagRepository interface {
	FindAll() ([]models.TagResponseWithPost, error)
	Create(tag *models.Tag) error
}