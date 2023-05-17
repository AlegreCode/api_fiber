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

	book.Author = author
	database.DB.Create(&book)
	return c.JSON(book)
}

func (bc *BookController) GetAllBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Preload("Author").Find(&books)
	return c.JSON(books)
}

func (bc *BookController) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	database.DB.Preload("Author").First(&book, id)
	return c.JSON(book)
}

func (bc *BookController) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	result := database.DB.First(&book, id)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid book ID",
		})
	}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   err.Error(),
		})
	}

	database.DB.Model(&book).Updates(&book)

	return c.JSON(book)
}
