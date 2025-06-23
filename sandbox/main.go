package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Создаем новое приложение Fiber
	app := fiber.New()

	// Определяем маршрут для GET запроса на корневой путь
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Привет, мир! 👋")
	})

	// Маршрут с параметром
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		return c.SendString("Привет, " + c.Params("name") + "! 🎉")
	})

	// Маршрут для POST запроса
	app.Post("/api", func(c *fiber.Ctx) error {
		// Возвращаем JSON ответ
		return c.JSON(fiber.Map{
			"success": true,
			"message": "API работает!",
		})
	})

	// Запускаем сервер на порту 3000
	app.Listen(":3000")
}
