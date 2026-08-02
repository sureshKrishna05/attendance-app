package app

import (
	"attendance-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Config   *config.Config
	Server   *fiber.App
	Database *pgxpool.Pool
}
