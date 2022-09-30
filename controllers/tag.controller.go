package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/services"
)

type TagController struct {
	tagService services.TagService
}

func NewTagController(tagService *services.TagService) TagController {
	return TagController{
		tagService: *tagService,
	}
}

func (tc *TagController) GetAll(c *fiber.Ctx) error {
	tags, _ := tc.tagService.GetAllData()

	return c.JSON(fiber.Map{
		"tags": tags,
	})
}

func (tc *TagController) Create(c *fiber.Ctx) error {
	tagRequest := new(models.TagRequest)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(tagRequest); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	errValidation := tc.tagService.Validation(*tagRequest)
	if errValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": errValidation.Error(),
		})
	}

	tag, errNew := tc.tagService.New(*tagRequest)
	if errNew != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"tag":     tag,
	})
}
