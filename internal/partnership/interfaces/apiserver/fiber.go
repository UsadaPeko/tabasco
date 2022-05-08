package apiserver

import "github.com/gofiber/fiber/v2"

func StartHTTPServer() {
	MakeServer().Listen(":3000")
}

func MakeServer() *fiber.App {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Partnership root page")
	})

	return app
}
