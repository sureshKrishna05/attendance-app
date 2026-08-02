package main

import (
	"log"

	"attendance-api/internal/bootstrap"
	"attendance-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create Fiber application
	app := fiber.New(fiber.Config{
		AppName: "University Attendance Management API",
	})

	// Global middleware
	app.Use(fiberLogger.New())
	app.Use(cors.New())

	// Health endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "UP",
		})
	})

	// Register middleware
	bootstrap.SetupMiddlewares(app)

	// Register routes
	bootstrap.SetupRoutes(app)

	log.Printf("Server starting on port %s", cfg.Port)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
