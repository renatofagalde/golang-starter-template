package main

import (
	"bootstrap/internal/adapter/input/api/middleware"
	"bootstrap/internal/adapter/input/api/routes"
	postgres "bootstrap/internal/config/database/postgres/gorm"
	tools "bootstrap/tool"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

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
	ginLambda = ginadapter.New(router)
	lambda.Start(ginLambda.ProxyWithContext)

}
