package services

import (
	"errors"
	"log"

	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/repositories"
)

type PostServiceImplement struct {
	postRepo    repositories.PostRepository
	postTagRepo repositories.PostTagRepository
}

// GetAllData implements PostService
func (ps *PostServiceImplement) GetAllData() ([]models.PostResponseWithTag, error) {
	posts, errFindAll := ps.postRepo.FindAll()
	if errFindAll != nil {
		return posts, errFindAll
	}

	return posts, nil
}

// New implements PostService
func (ps *PostServiceImplement) New(postRequest models.PostRequest) (models.PostResponse, error) {
	var post models.Post
	var postResponse models.PostResponse

	post.Title = postRequest.Title
	post.Body = postRequest.Body
	post.UserID = postRequest.UserID

	post, errCreate := ps.postRepo.Create(&post)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return postResponse, errCreate
	}

	if len(postRequest.TagsID) > 0 {
		for _, tagID := range postRequest.TagsID {
			postTag := new(models.PostTag)
			postTag.PostID = post.ID
			postTag.TagID = tagID
			errCreatePostTag := ps.postTagRepo.Create(postTag)
			if errCreatePostTag != nil {
				log.Println(errCreatePostTag.Error())
			}
		}
	}

	postResponse.ID = post.ID
	postResponse.Title = post.Title
	postResponse.Body = post.Body
	postResponse.UserID = post.UserID

	return postResponse, nil
}

func (ps *PostServiceImplement) Validation(postRequest models.PostRequest) error {
	var messageError string
	var isError bool

	// MANUAL VALIDATION
	if postRequest.Title == "" {
		messageError += "title is required, "
		isError = true
	}
	if postRequest.Body == "" {
		messageError += "body is required, "
		isError = true
	}
	if postRequest.UserID == 0 {
		messageError += "user_id is required, "
		isError = true
	}
	if len(postRequest.TagsID) == 0 {
		messageError += "tags_id is required, "
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewPostService(postRepo *repositories.PostRepository, postTagRepo *repositories.PostTagRepository) PostService {
	return &PostServiceImplement{
		postRepo:    *postRepo,
		postTagRepo: *postTagRepo,
	}
}
