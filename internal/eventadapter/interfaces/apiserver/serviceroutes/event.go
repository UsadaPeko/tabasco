package serviceroutes

import (
	"github.com/UsadaPeko/jsn"
	"github.com/gofiber/fiber/v2"
	"net/http"

	"gomod.pekora.dev/tabasco/internal/eventadapter/usecases"
)

func Event(c *fiber.Ctx) error {
	serviceKey := c.Params("serviceKey")
	uc := usecases.ServiceIntegrationUseCases{}

	if uc.ValidateServiceKey(serviceKey) != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	j, err := jsn.New(string(c.Body()))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	_, ok := j.StringVal("user_id")
	if !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.SendStatus(http.StatusOK)
}
