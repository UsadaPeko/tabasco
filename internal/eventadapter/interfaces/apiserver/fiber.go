package apiserver

import (
	"github.com/gofiber/fiber/v2"
	"gomod.pekora.dev/tabasco/internal/eventadapter/interfaces/apiserver/serviceroutes"
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
		app.Post("/service/:serviceKey/event", serviceroutes.Event)
	}
	{
		app.Post("/rules", func(c *fiber.Ctx) error {
			return c.SendStatus(http.StatusCreated)
		})
	}
	return app
}
