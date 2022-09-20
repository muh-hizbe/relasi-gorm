package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/models"
)

func TagGetAll(c *fiber.Ctx) error {
	var tags []models.TagResponseWithPost

	database.DB.Preload("Posts").Find(&tags)

	return c.JSON(fiber.Map{
		"tags": tags,
	})
}

func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(tag); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	// MANUAL VALIDATION
	if tag.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	database.DB.Debug().Create(&tag)

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"tag":     tag,
	})
}
