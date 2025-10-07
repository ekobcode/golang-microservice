package app

import (
	"database/sql"
	"fmt"

	httpDelivery "golang-microservice/internal/delivery/http"
	"golang-microservice/internal/delivery/http/handler"
	"golang-microservice/internal/infrastructure/config"
	"golang-microservice/internal/repository"
	"golang-microservice/internal/usecase"

	"go.uber.org/zap"
)

type App struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
	db     *sql.DB
}

func NewApp(cfg *config.Config, logger *zap.SugaredLogger, db *sql.DB) *App {
	return &App{cfg, logger, db}
}

func (a *App) Run() {
	userRepo := repository.NewUserRepository(a.db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC, a.logger)

	router := httpDelivery.NewRouter(userHandler, a.cfg.APIKey)

	a.logger.Infof("🚀 Server running on port %s", a.cfg.AppPort)
	router.Run(fmt.Sprintf(":%s", a.cfg.AppPort))
}
