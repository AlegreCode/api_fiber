package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/alegrecode/api_fiber/models"
)

func AuthorController(db *gorm.DB) *AuthorHandler {
	return &AuthorHandler{db}
}

type AuthorHandler struct {
	db *gorm.DB
}

func (ah *AuthorHandler) GetAllAuthors(c *fiber.Ctx) error {
	var authors []models.Author
	ah.db.Find(&authors)
	return c.JSON(authors)
}

func (ah *AuthorHandler) CreateAuthor(c *fiber.Ctx) error {

	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create author",
			"error":   err.Error(),
		})
	}

	ah.db.Create(&author)
	return c.JSON(author)
}
