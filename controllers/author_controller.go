package controllers

import "github.com/gofiber/fiber/v2"

func AuthorController(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
