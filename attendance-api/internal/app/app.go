package app

import (
	"attendance-api/internal/config"

	"github.com/gofiber/fiber/v2"
)

type Application struct {
	Config *config.Config
	Server *fiber.App
}

func New(cfg *config.Config) *Application {
	return &Application{
		Config: cfg,
		Server: fiber.New(fiber.Config{
			AppName: "University Attendance Management API",
		}),
	}
}

func (a *Application) Fiber() *fiber.App {
	return a.Server
}
