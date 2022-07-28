package main

import (
	"github.com/leoantony72/rplace/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controllers.Ping)
	router.GET("/board", controllers.Get_Board)

	router.Run("localhost:8080")

}
