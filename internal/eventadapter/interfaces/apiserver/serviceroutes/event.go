package serviceroutes

import (
	"github.com/gofiber/fiber/v2"
	"gomod.pekora.dev/tabasco/internal/eventadapter/usecases"
	"net/http"
)

func Event(c *fiber.Ctx) error {
	serviceKey := c.Params("serviceKey")
	uc := usecases.ServiceIntegrationUseCases{}

	if uc.ValidateServiceKey(serviceKey) != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}
	return c.SendStatus(http.StatusOK)
}
