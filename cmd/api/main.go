package main

import (
	"bootstrap/internal/adapter/input/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(":8080")
}
