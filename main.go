package main

import (
	"github.com/gin-gonic/gin"
	"github.com/anuprajwal/go-mod/routes"
)

func main(){
	router := gin.Default()

	routes.RegesteredRoutes(router)

	router.Run(":8080")
}