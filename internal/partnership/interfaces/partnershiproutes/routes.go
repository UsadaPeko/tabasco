package partnershiproutes

import (
	"github.com/UsadaPeko/jsn"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func PostPartnership(ctx *fiber.Ctx) error {
	responseBody := jsn.Init()
	responseBody.Set("id", "97FD9E6E-56A6-44B8-8411-E8D3EFD96D6C")
	return ctx.Status(http.StatusCreated).SendString(responseBody.String())
}

func GetPartnership(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	responseBody := jsn.Init()
	responseBody.Set("id", id)
	responseBody.Set("name", "Tabasco")
	return ctx.Status(http.StatusOK).SendString(responseBody.String())
}
