package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alegrecode/api_fiber/api"
	"github.com/alegrecode/api_fiber/database"
	"github.com/alegrecode/api_fiber/models"
)

func main() {
	app := fiber.New()

	db := database.ConnectDB()

	db.AutoMigrate(&models.Author{}, &models.Book{})

	api.Router(app)

	app.Listen(":3000")
}
