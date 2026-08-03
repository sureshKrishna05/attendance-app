package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	Port           string
	DBURL          string
	SecretKey      string
	PublishableKey string
}

var cfg *Config

func LoadConfig() *Config {
	if cfg != nil {
		return cfg
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env found, using environment variables")
	}

	cfg = &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   getEnv("PORT", "8080"),
		DBURL: getEnv(
			"DATABASE_URL",
			"postgres://postgres:postgres@localhost:5432/univ_project?sslmode=disable",
		),
		SecretKey:      getEnv("SECRET_KEY", "my_super_secret_key"),
		PublishableKey: getEnv("PUBLISHABLE_KEY", "my_publishable_key"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
