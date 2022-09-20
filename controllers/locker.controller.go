package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

func LockerGetAll(c *fiber.Ctx) error {
	var lockers []models.Locker

	database.DB.Preload("User").Find(&lockers)

	return c.JSON(fiber.Map{
		"lockers": lockers,
	})
}

func CreateLocker(c *fiber.Ctx) error {
	locker := new(models.Locker)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(locker); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	// MANUAL VALIDATION
	if locker.Code == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "code is required",
		})
	}
	if locker.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "user_id is required",
		})
	}

	database.DB.Create(&locker)

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"locker":  locker,
	})
}
