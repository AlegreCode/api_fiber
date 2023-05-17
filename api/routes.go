package api

import (
	"github.com/alegrecode/api_fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	var authorController = &controllers.AuthorController{}
	var bookController = &controllers.BookController{}

	app.Get("/authors", authorController.GetAllAuthors)

	app.Post("/author", authorController.CreateAuthor)

	app.Get("/author/:id", authorController.GetAuthor)

	app.Put("/author/:id", authorController.UpdateAuthor)

	app.Delete("/author/:id", authorController.DeleteAuthor)

	app.Get("/books", bookController.GetAllBooks)

	app.Post("/book", bookController.CreateBook)

	app.Get("/book/:id", bookController.GetBook)
}
