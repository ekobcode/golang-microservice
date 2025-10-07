package http

import (
	"golang-microservice/internal/delivery/http/handler"
	"golang-microservice/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *handler.UserHandler, apiKey string) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.APIKeyAuth(apiKey))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", userHandler.Create)
		v1.GET("/users", userHandler.GetAll)
		v1.GET("/users/:id", userHandler.GetByID)
		v1.PUT("/users/:id", userHandler.Update)
		v1.DELETE("/users/:id", userHandler.Delete)
	}

	return r
}
