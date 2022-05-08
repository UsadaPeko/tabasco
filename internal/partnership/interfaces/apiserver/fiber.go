package apiserver

import (
	"github.com/UsadaPeko/jsn"
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
		responseBody := jsn.Init()
		responseBody.Set("id", "AF3D1D06-BB2D-470C-A842-360195FD046A")
		return ctx.Status(http.StatusCreated).SendString(responseBody.String())
	})
	return app
}
