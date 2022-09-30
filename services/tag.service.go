package services

import "github.com/muh-hizbe/relasi-gorm/models"

type TagService interface {
	GetAllData() ([]models.TagResponseWithPost, error)
	New(tagRequest models.TagRequest) (models.Tag, error)
	Validation(tagRequest models.TagRequest) error
}
