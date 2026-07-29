package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	Port           string
	DBUrl          string
	PublishableKey string
	SecretKey      string
	JWTSecret      string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	return &Config{
		AppEnv:         getEnv("APP_ENV", "development"),
		Port:           getEnv("PORT", "8080"),
		DBUrl:          getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/univ_project?sslmode=disable"),
		PublishableKey: getEnv("PUBLISHABLE_KEY", "pk_test_12345"),
		SecretKey:      getEnv("SECRET_KEY", "sk_test_12345"),
		JWTSecret:      getEnv("JWT_SECRET", "supersecretjwtkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
