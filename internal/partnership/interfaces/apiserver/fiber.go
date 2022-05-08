package apiserver

import (
	"github.com/UsadaPeko/jsn"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	id := uuid.NewString()
	app.Post("/partnership", func(ctx *fiber.Ctx) error {
		responseBody := jsn.Init()
		responseBody.Set("id", id)
		return ctx.Status(http.StatusCreated).SendString(responseBody.String())
	})
	app.Post("/partnership/"+id, func(ctx *fiber.Ctx) error {
		responseBody := jsn.Init()
		responseBody.Set("id", id)
		responseBody.Set("name", "Tabasco")
		return ctx.Status(http.StatusOK).SendString(responseBody.String())
	})
	return app
}
