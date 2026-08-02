package bootstrap

import (
	"attendance-api/internal/shared/response"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return response.Success(
			c,
			fiber.StatusOK,
			"Attendance API is healthy",
			fiber.Map{
				"status":  "UP",
				"service": "attendance-api",
				"version": "v1",
			},
		)
	})
}
