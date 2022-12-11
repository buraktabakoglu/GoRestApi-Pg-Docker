package main

import (
	"github.com/buraktabakoglu/gin-gorm-rest/config"
	"github.com/buraktabakoglu/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	config.Connect()

	routes.UserRoute(router)

	router.Run(":8080")
}
