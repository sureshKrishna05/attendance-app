package bootstrap

import (
	"attendance-api/configs"
	"attendance-api/internal/modules/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *configs.Config) {
	// API v1 Group
	api := app.Group("/api/v1")

	// Setup Auth Routes (Public, but requires API Key)
	auth.SetupRoutes(api, cfg)

	// Example Protected Route
	api.Get("/protected", ProtectedRoute(cfg), func(c *fiber.Ctx) error {
		userID := c.Locals("userID")
		role := c.Locals("role")
		return c.JSON(fiber.Map{
			"message": "You have accessed a protected route",
			"user":    userID,
			"role":    role,
		})
	})
}
