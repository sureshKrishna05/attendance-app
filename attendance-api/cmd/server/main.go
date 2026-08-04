package main

import (
	"log"

	"attendance-api/internal/bootstrap"
	"attendance-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	// Connect to Database
	db := bootstrap.ConnectDatabase(cfg)
	defer db.Close()

	// Initialize Fiber App with Cloudflare proxy support
	app := fiber.New(fiber.Config{
		AppName: "University Attendance API v1.1",
		
		// Cloudflare Tunnel acts as a proxy, so we trust local IPs
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"127.0.0.1", "10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"},
		
		// Cloudflare specifically sets this header with the real user's IP
		ProxyHeader: "CF-Connecting-IP",
	})

	// Global Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Basic Health Check (Bypasses API Key due to middleware logic)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Setup custom middlewares
	bootstrap.SetupMiddlewares(app)

	// Setup Routes
	bootstrap.SetupRoutes(app, cfg, db)

	// Start Server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
