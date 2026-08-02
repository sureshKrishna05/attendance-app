package bootstrap

import (
	"attendance-api/internal/config"
	"attendance-api/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDatabase(cfg *config.Config) (*pgxpool.Pool, error) {
	return database.NewPool(cfg)
}
