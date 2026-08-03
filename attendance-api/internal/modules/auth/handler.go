package auth

import (
	"attendance-api/internal/config"
	"attendance-api/internal/modules/users"
	"attendance-api/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role"` // "student" or "faculty"
}

func SetupRoutes(router fiber.Router, cfg *config.Config, userService users.Service) {
	authGroup := router.Group("/auth")

	authGroup.Post("/login", func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		user, err := userService.VerifyCredentials(c.Context(), req.ID, req.Password, req.Role)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		// Generate JWT Token
		token, err := jwt.GenerateToken(user.ID, user.Role, cfg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}

		return c.JSON(fiber.Map{
			"message": "Login successful",
			"token":   token,
			"user": fiber.Map{
				"id":   user.ID,
				"name": user.Name,
				"role": user.Role,
			},
		})
	})
}
