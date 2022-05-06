package apiserver

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func StartHTTPServer() {
	MakeServer().Listen(":3000")
}

func MakeServer() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	{
		app.Post("/service/:serviceKey/event", func(c *fiber.Ctx) error {
			return c.SendStatus(http.StatusOK)
		})
	}
	return app
}
