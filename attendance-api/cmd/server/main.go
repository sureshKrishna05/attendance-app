package main

import (
	"log"

	"attendance-api/configs"
	"attendance-api/internal/bootstrap"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load Configuration
	cfg := configs.LoadConfig()

	// Initialize Fiber App
	app := fiber.New(fiber.Config{
		AppName: "University Attendance API v1.1",
	})

	// Global Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Basic Health Check (Bypasses API Key due to middleware logic)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Setup custom middlewares (API Key, JWT, etc)
	bootstrap.SetupMiddlewares(app, cfg)

	// Setup Routes
	bootstrap.SetupRoutes(app, cfg)

	// Start Server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
