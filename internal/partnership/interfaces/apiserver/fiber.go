package apiserver

import (
	"github.com/gofiber/fiber/v2"
	"gomod.pekora.dev/tabasco/internal/partnership/interfaces/partnershiproutes"
)

func StartHTTPServer() {
	MakeServer().Listen(":3000")
}

func MakeServer() *fiber.App {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Partnership root page")
	})

	app.Post("/partnership", partnershiproutes.PostPartnership)
	app.Post("/partnership/:id", partnershiproutes.GetPartnership)
	return app
}
