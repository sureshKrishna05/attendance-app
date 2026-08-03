package bootstrap

import (
	"context"
	"log"

	"attendance-api/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabase(cfg *config.Config) *pgxpool.Pool {
	poolConfig, err := pgxpool.ParseConfig(cfg.DBURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	// You can adjust max connections and other settings here
	poolConfig.MaxConns = 10

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Test connection
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database")
	return db
}
