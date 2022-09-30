package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/services"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService *services.PostService) PostController {
	return PostController{
		postService: *postService,
	}
}

func (pc *PostController) GetAll(c *fiber.Ctx) error {
	posts, _ := pc.postService.GetAllData()

	return c.JSON(fiber.Map{
		"posts": posts,
	})
}

func (pc *PostController) Create(c *fiber.Ctx) error {
	postRequest := new(models.PostRequest)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(postRequest); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	errValidation := pc.postService.Validation(*postRequest)
	if errValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": errValidation.Error(),
		})
	}

	postResponse, errNew := pc.postService.New(*postRequest)
	if errNew != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"post":    postResponse,
	})
}
