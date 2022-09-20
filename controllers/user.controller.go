package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

func UserGetAll(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Preload("Locker").Preload("Posts").Find(&users)

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	// MANUAL VALIDATION
	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"user":    user,
	})
}
