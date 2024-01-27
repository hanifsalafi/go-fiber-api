package route

import (
	"github.com/gofiber/fiber/v2"
)

func MainRouteInit(r *fiber.App) {
	r.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "Green",
		})
	})
}
