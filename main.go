package main

import (
	"API/configs"
	"API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()
	routes.Routes(router)

	router.Run("localhost:8080")
}
