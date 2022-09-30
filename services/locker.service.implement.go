package services

import (
	"errors"
	"log"

	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/repositories"
)

type LockerServiceImplement struct {
	lockerRepo repositories.LockerRepository
}

// New implements LockerService
func (ls *LockerServiceImplement) New(lockerRequest models.LockerRequest) (models.LockerResponse, error) {
	var locker models.Locker
	locker.Code = lockerRequest.Code
	locker.UserID = lockerRequest.UserID

	var lockerResponse models.LockerResponse

	locker, errCreate := ls.lockerRepo.Create(&locker)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return lockerResponse, errCreate
	}

	lockerResponse.ID = locker.ID
	lockerResponse.Code = locker.Code
	lockerResponse.UserID = locker.UserID

	return lockerResponse, nil
}

// getAllData implements LockerService
func (ls *LockerServiceImplement) GetAllData() ([]models.Locker, error) {
	lockers, errFindAll := ls.lockerRepo.FindAll()
	if errFindAll != nil {
		log.Fatalln(errFindAll.Error())
		return lockers, errFindAll
	}

	return lockers, nil
}

func (ls *LockerServiceImplement) Validation(lockerRequest models.LockerRequest) error {
	var messageError string
	var isError bool
	// MANUAL VALIDATION
	if lockerRequest.Code == "" {
		messageError = "code is required, "
		isError = true
	}
	if lockerRequest.UserID == 0 {
		messageError += "user_id is required, "
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewLockerService(lockerRepo *repositories.LockerRepository) LockerService {
	return &LockerServiceImplement{
		lockerRepo: *lockerRepo,
	}
}
