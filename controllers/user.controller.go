package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/models"
	"github.com/muh-hizbe/relasi-gorm/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) UserController {
	return UserController{
		userService: *userService,
	}
}

func (uc *UserController) GetAll(c *fiber.Ctx) error {
	users, _ := uc.userService.GetAllData()

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	userRequest := new(models.UserRequest)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	errValidation := uc.userService.Validation(*userRequest)
	if errValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": errValidation.Error(),
		})
	}

	user, errNew := uc.userService.New(*userRequest)
	if errNew != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"user":    user,
	})
}
