package controllers

import (
	"github.com/alegrecode/api_fiber/database"
	"github.com/alegrecode/api_fiber/models"
	"github.com/gofiber/fiber/v2"
)

type BookController struct {
}

func (bc *BookController) CreateBook(c *fiber.Ctx) error {
	var book models.Book
	var author models.Author

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create book",
			"error":   err.Error(),
		})
	}

	result := database.DB.First(&author, book.AuthorID)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid author ID",
		})
	}

	database.DB.Create(&book)
	return c.JSON(book)
}
