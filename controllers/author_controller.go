package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alegrecode/api_fiber/database"
	"github.com/alegrecode/api_fiber/models"
)

type AuthorController struct {
}

func (ac *AuthorController) GetAllAuthors(c *fiber.Ctx) error {
	var authors []models.Author
	database.DB.Find(&authors)
	return c.JSON(authors)
}

func (ac *AuthorController) CreateAuthor(c *fiber.Ctx) error {

	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create author",
			"error":   err.Error(),
		})
	}

	database.DB.Create(&author)
	return c.JSON(author)
}

func (ac *AuthorController) GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author

	database.DB.First(&author, id)
	// ah.db.First(&author, id)

	if author.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Author not found",
		})
	}

	return c.JSON(author)
}

func (ac *AuthorController) UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author

	database.DB.First(&author, id)
	// ah.db.First(&author, id)

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

	database.DB.Save(&author)
	// ah.db.Save(&author)

	return c.JSON(author)
}

func (ac *AuthorController) DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author

	database.DB.First(&author, id)
	// ah.db.First(&author, id)

	if author.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Author not found",
		})
	}

	database.DB.Delete(&author)
	// ah.db.Delete(&author)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Author deleted",
	})
}
