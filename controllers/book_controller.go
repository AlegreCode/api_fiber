package controllers

import (
	"github.com/alegrecode/api_fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BookController(db *gorm.DB, authorHandler *AuthorHandler) *BookHandler {
	return &BookHandler{db: db, authorHandler: authorHandler}
}

type BookHandler struct {
	db            *gorm.DB
	authorHandler *AuthorHandler
}

func (bh *BookHandler) CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create book",
			"error":   err.Error(),
		})
	}

	author, err := bh.authorHandler.GetAuthor(book.AuthorID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid author ID",
		})
	}

	book.Author = *author
	bh.db.Create(&book)
	return c.JSON(book)
}
