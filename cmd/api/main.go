package main

import (
	"bootstrap/internal/adapter/input/api/routes"
	"bootstrap/internal/config/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.InitRoutes(router)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error starting server on port 8080", err)
	}
}
