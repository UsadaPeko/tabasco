package serviceroutes

import (
	"encoding/json"
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

	jsonObject := map[string]interface{}{}
	err := json.Unmarshal(c.Body(), &jsonObject)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	_, ok := jsonObject["user_id"]
	if !ok {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.SendStatus(http.StatusOK)
}
