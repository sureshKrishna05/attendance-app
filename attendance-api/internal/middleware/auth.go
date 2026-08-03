package middleware

import (
	"strings"

	"attendance-api/internal/config"
	"attendance-api/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

// ProtectedRoute validates the JWT and API key.
func ProtectedRoute(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Validate API Key
		apiKey := c.Get("X-Api-Key")
		if apiKey != cfg.PublishableKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid API key"})
		}

		// Validate JWT
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateToken(tokenString, cfg.SecretKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Set locals for subsequent handlers
		c.Locals("userID", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// RoleMiddleware restricts access based on user role.
func RoleMiddleware(allowedRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role != allowedRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}
		return c.Next()
	}
}
