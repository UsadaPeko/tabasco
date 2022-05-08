package partnershiproutes

import (
	"github.com/UsadaPeko/jsn"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

var (
	cache map[string]*jsn.JSON = map[string]*jsn.JSON{}
)

func PostPartnership(ctx *fiber.Ctx) error {
	id := uuid.NewString()

	requestBody, err := jsn.New(string(ctx.Body()))
	if err != nil {
		return err
	}

	serviceName, ok := requestBody.StringVal("name")
	if !ok {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	responseBody := jsn.Init()
	responseBody.Set("id", id)
	responseBody.Set("name", serviceName)

	cache[id] = responseBody
	return ctx.Status(http.StatusCreated).SendString(responseBody.String())
}

func GetPartnership(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	responseBody := jsn.Init()
	responseBody.Set("id", id)
	responseBody.Set("name", "Tabasco")
	return ctx.Status(http.StatusOK).SendString(responseBody.String())
}
