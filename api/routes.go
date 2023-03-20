package api

import (
	"github.com/alegrecode/api_fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", controllers.AuthorController)
}
