package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Привет, мир! 👋")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		return c.SendString("Привет, " + c.Params("name") + "! 🎉")
	})
	
	app.Post("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "API работает!",
		})
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
