package main

import (
	"log"

	"golang-microservice/internal/app"
	"golang-microservice/internal/infrastructure/config"
	"golang-microservice/internal/infrastructure/db"
	"golang-microservice/internal/infrastructure/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize logger (Zap)
	zapLogger := logger.NewLogger()

	// Connect to database
	dbConn, err := db.NewPostgresDB(cfg)
	if err != nil {
		zapLogger.Fatal("failed to connect database", err)
	}

	// Run app
	app := app.NewApp(cfg, zapLogger, dbConn)
	app.Run()
}
