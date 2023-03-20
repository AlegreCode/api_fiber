package api

import (
	"github.com/alegrecode/api_fiber/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {

	authorController := controllers.AuthorController(db)

	app.Post("/author", authorController.CreateAuthor)
}
