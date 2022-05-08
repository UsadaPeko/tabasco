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

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Partnership root page")
	})

	app.Post("/partnership", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusCreated)
	})
	return app
}
