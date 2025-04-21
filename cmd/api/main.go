package main

import (
	"bootstrap/internal/adapter/input/api/middleware"
	"bootstrap/internal/adapter/input/api/routes"
	postgres "bootstrap/internal/config/database/postgres/gorm"
	"bootstrap/internal/config/logger"
	tools "bootstrap/tool"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	var tools tools.ToolLoadEnvironmet

	sqlConfig, err := tools.Do()
	if err != nil {
		log.Fatal(fmt.Sprintf("Erro FATAL %+v", err))
	}

	database, err := postgres.NewPostgresGORMConnection(context.Background(), sqlConfig.DBSource)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco, error=%s", err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.HeadersMiddleware())

	routes.InitRoutes(&router.RouterGroup, database)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error starting server on port 8080", err)
	}
}
