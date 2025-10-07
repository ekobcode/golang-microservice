package config

import (
	"os"
)

type Config struct {
	AppPort string
	APIKey  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		APIKey:     getEnv("API_KEY", "my-secret-key"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password123#"),
		DBName:     getEnv("DB_NAME", "go-microservice"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
