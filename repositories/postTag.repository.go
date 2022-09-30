package repositories

import "github.com/muh-hizbe/relasi-gorm/models"

type PostTagRepository interface {
	Create(postTag *models.PostTag) error
}