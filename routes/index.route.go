package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muh-hizbe/relasi-gorm/controllers"
	"github.com/muh-hizbe/relasi-gorm/repositories"
	"github.com/muh-hizbe/relasi-gorm/services"
)

func RouteInit(app *fiber.App) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(&userRepo)
	userController := controllers.NewUserController(&userService)
	app.Get("/users", userController.GetAll)
	app.Post("/users", userController.Create)

	lockerRepo := repositories.NewLockerRepository()
	lockerService := services.NewLockerService(&lockerRepo)
	lockerController := controllers.NewLockerController(&lockerService)
	app.Get("/lockers", lockerController.GetAll)
	app.Post("/lockers", lockerController.Create)

	postRepo := repositories.NewPostRepository()
	postTagRepo := repositories.NewPostTagRepository()
	postService := services.NewPostService(&postRepo, &postTagRepo)
	postController := controllers.NewPostController(&postService)
	app.Get("/posts", postController.GetAll)
	app.Post("/posts", postController.Create)

	tagRepo := repositories.NewTagReposotory()
	tagService := services.NewTagService(&tagRepo)
	tagController := controllers.NewTagController(&tagService)
	app.Get("/tags", tagController.GetAll)
	app.Post("/tags", tagController.Create)
}
