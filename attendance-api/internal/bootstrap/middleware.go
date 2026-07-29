package bootstrap

import (
	"strings"

	"attendance-api/configs"
	"attendance-api/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func SetupMiddlewares(app *fiber.App, cfg *configs.Config) {
	// API Key Middleware (Publishable Key for Client Verification)
	app.Use(func(c *fiber.Ctx) error {
		// Example: Allow health check without API key
		if c.Path() == "/health" {
			return c.Next()
		}

		clientKey := c.Get("X-Api-Key")
		if clientKey == "" || clientKey != cfg.PublishableKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid or missing API Key",
			})
		}
		return c.Next()
	})
}

// ProtectedRoute middleware to enforce JWT validation
func ProtectedRoute(cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Missing or invalid token",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateToken(tokenStr, cfg)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: " + err.Error(),
			})
		}

		// Store user info in context for later handlers
		c.Locals("userID", claims.UserID)
		c.Locals("role", claims.Role)
		return c.Next()
	}
}

// RoleMiddleware checks if the user has the required role
func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: Insufficient privileges",
			})
		}
		return c.Next()
	}
}
