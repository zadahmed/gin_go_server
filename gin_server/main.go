package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zadahmed/gin_go_server/controller"
)

func main() {
	r := gin.Default() 

	r.GET("/", controllers.Get)

	r.Run(":9090")

}