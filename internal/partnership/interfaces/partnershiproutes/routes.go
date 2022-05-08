package partnershiproutes

import (
	"github.com/UsadaPeko/jsn"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/speps/go-hashids/v2"
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

	responseBody, ok := cache[id]
	if !ok {
		return ctx.SendStatus(http.StatusNotFound)
	}

	return ctx.Status(http.StatusOK).SendString(responseBody.String())
}

func PostPartnershipNewIntegrations(ctx *fiber.Ctx) error {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 15
	id := ctx.Params("id")

	uid := uuid.MustParse(id)

	hids, _ := hashids.NewWithData(hd)
	key, _ := hids.Encode([]int{uid.ClockSequence()})

	responseBody := jsn.Init()
	responseBody.Set("key", key)

	return ctx.Status(http.StatusCreated).SendString(responseBody.String())
}
