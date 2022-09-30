package services

import (
	"errors"
	"log"

	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/repositories"
)

type UserServiceImplement struct {
	userRepo repositories.UserRepository
}

// GetAllData implements UserService
func (us *UserServiceImplement) GetAllData() ([]models.User, error) {
	users, errFindAll := us.userRepo.FindAll()
	if errFindAll != nil {
		return users, errFindAll
	}

	return users, nil
}

// New implements UserService
func (us *UserServiceImplement) New(userRequest models.UserRequest) (models.UserResponse, error) {
	var userResponse models.UserResponse
	user := models.User{
		Name: userRequest.Name,
	}

	errCreate := us.userRepo.Create(&user)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return userResponse, errCreate
	}

	userResponse.ID = user.ID
	userResponse.Name = user.Name

	return userResponse, nil
}

func (us *UserServiceImplement) Validation(userRequest models.UserRequest) error {
	var messageError string
	var isError bool

	// MANUAL VALIDATION
	if userRequest.Name == "" {
		messageError += "name is required"
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewUserService(userRepo *repositories.UserRepository) UserService {
	return &UserServiceImplement{
		userRepo: *userRepo,
	}
}
