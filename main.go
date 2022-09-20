package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/database"
	"github.com/muh-hizbe/relasi-gorm/database/migrations"
	"github.com/muh-hizbe/relasi-gorm/routes"
)

func main() {
	// CONNECTION TO DATABASE
	database.DatabaseInit()

	// MIGRATION
	migrations.Migration()

	// FIBER INIT
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "okay",
			"age":     24,
		})
	})

	routes.RouteInit(app)

	app.Listen(":8000")
}
