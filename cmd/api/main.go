package main

import (
	"bootstrap/internal/adapter/input/api/middleware"
	"bootstrap/internal/adapter/input/api/routes"
	"bootstrap/internal/config/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.HeadersMiddleware())
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error starting server on port 8080", err)
	}
}
