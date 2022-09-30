package services

import "github.com/muh-hizbe/relasi-gorm/models"

type PostService interface {
	GetAllData() ([]models.PostResponseWithTag, error)
	New(postRequest models.PostRequest) (models.PostResponse, error)
	Validation(postRequest models.PostRequest) error
}
