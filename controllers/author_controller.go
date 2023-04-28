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

func (ah *AuthorHandler) GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	ah.db.First(&author, id)

	if author.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Author not found",
		})
	}

	return c.JSON(author)
}

func (ah *AuthorHandler) UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	ah.db.First(&author, id)

	if author.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Author not found",
		})
	}

	updatedAuthor := new(models.Author)

	if err := c.BodyParser(&updatedAuthor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   err.Error(),
		})
	}

	author.Name = updatedAuthor.Name
	author.Lastname = updatedAuthor.Lastname
	ah.db.Save(&author)

	return c.JSON(author)
}

func (ah *AuthorHandler) DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	ah.db.First(&author, id)

	if author.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Author not found",
		})
	}

	ah.db.Delete(&author)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Author deleted",
	})
}
