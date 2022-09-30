package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/services"
)

type LockerController struct {
	lockerService services.LockerService
}

func NewLockerController(lockerService *services.LockerService) LockerController {
	return LockerController{
		lockerService: *lockerService,
	}
}

func (lc *LockerController) GetAll(c *fiber.Ctx) error {
	lockers, _ := lc.lockerService.GetAllData()

	return c.JSON(fiber.Map{
		"lockers": lockers,
	})
}

func (lc *LockerController) Create(c *fiber.Ctx) error {
	lockerRequest := new(models.LockerRequest)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(lockerRequest); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	errValidation := lc.lockerService.Validation(*lockerRequest)
	if errValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": errValidation.Error(),
		})
	}

	lockerResponse, errNew := lc.lockerService.New(*lockerRequest)
	if errNew != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"locker":  lockerResponse,
	})
}
