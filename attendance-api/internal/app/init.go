package app

import (
	"fmt"

	"attendance-api/internal/bootstrap"
	"attendance-api/internal/config"

	"github.com/gofiber/fiber/v2"
)

func New() (*Application, error) {

	cfg := config.LoadConfig()

	db, err := bootstrap.SetupDatabase(cfg)
	if err != nil {
		return nil, fmt.Errorf("database initialization failed: %w", err)
	}

	server := fiber.New(fiber.Config{
		AppName: "University Attendance Management API",
	})

	bootstrap.SetupMiddlewares(server)
	bootstrap.SetupRoutes(server)

	return &Application{
		Config:   cfg,
		Server:   server,
		Database: db,
	}, nil
}

func (a *Application) Run() error {
	return a.Server.Listen(":" + a.Config.Port)
}

func (a *Application) Shutdown() {
	if a.Database != nil {
		a.Database.Close()
	}
}
