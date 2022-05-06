package serviceroutes

import (
	"github.com/gofiber/fiber/v2"
	"gomod.pekora.dev/tabasco/internal/eventadapter/usecases"
	"net/http"
	"strings"
)

func Event(c *fiber.Ctx) error {
	serviceKey := c.Params("serviceKey")
	uc := usecases.ServiceIntegrationUseCases{}

	if uc.ValidateServiceKey(serviceKey) != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}
	if !strings.Contains(string(c.Body()), "user_id") {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.SendStatus(http.StatusOK)
}
