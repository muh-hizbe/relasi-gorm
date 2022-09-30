package services

import (
	"errors"
	"log"

	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/repositories"
)

type TagServiceImplement struct {
	tagRepo repositories.TagRepository
}

// GetAllData implements TagService
func (ts *TagServiceImplement) GetAllData() ([]models.TagResponseWithPost, error) {
	tags, errFindAll := ts.tagRepo.FindAll()
	if errFindAll != nil {
		return tags, errFindAll
	}

	return tags, nil
}

// New implements TagService
func (ts *TagServiceImplement) New(tagRequest models.TagRequest) (models.Tag, error) {
	tag := models.Tag{
		Name: tagRequest.Name,
	}

	errCreate := ts.tagRepo.Create(&tag)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return tag, errCreate
	}

	return tag, nil
}

func (ts *TagServiceImplement) Validation(tagRequest models.TagRequest) error {
	var messageError string
	var isError bool

	// MANUAL VALIDATION
	if tagRequest.Name == "" {
		messageError += "name is required"
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewTagService(tagRepo *repositories.TagRepository) TagService {
	return &TagServiceImplement{
		tagRepo: *tagRepo,
	}
}
