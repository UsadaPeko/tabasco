package httpapi

import "github.com/gofiber/fiber/v2"

func LunchHTTPServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	go app.Listen(":3000")
}
