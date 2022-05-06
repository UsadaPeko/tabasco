package serviceroutes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Event(c *fiber.Ctx) error {
	if c.Params("serviceKey") == "InvalidServiceKey" {
		return c.SendStatus(http.StatusUnauthorized)
	}
	return c.SendStatus(http.StatusOK)
}
