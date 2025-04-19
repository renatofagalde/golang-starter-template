package main

import (
	"bootstrap/internal/adapter/input/routes"
	"bootstrap/internal/config/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	routes.InitRoutes(router)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error starting server on port 8080", err)
	}
}
