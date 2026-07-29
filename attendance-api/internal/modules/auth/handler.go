package auth

import (
	"attendance-api/configs"
	"attendance-api/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role"` // "student" or "faculty"
}

func SetupRoutes(router fiber.Router, cfg *configs.Config) {
	authGroup := router.Group("/auth")

	authGroup.Post("/login", func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Mock authentication logic (For now, accept if password is "password")
		if req.Password != "password" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		// Generate JWT Token
		token, err := jwt.GenerateToken(req.ID, req.Role, cfg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}

		return c.JSON(fiber.Map{
			"message": "Login successful",
			"token":   token,
			"user": fiber.Map{
				"id":   req.ID,
				"role": req.Role,
			},
		})
	})
}
