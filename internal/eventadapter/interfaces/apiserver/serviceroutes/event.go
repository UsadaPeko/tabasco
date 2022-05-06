package serviceroutes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Event(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
